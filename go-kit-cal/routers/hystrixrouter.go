package routers

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/hashicorp/consul/api"
	"github.com/openzipkin/zipkin-go"
	zipkinhttpsvr "github.com/openzipkin/zipkin-go/middleware/http"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"strings"
	"sync"
	"errors"
	"github.com/go-kit/kit/log"
)

type HystrixRouter struct {
	svcMap *sync.Map
	logger log.Logger
	fallbackMsg string
	consulClient *api.Client
	tracer *zipkin.Tracer
}

func Routes(client *api.Client,
	zipkinTracer *zipkin.Tracer,
	fallbackMsg string,
	logger log.Logger) http.Handler{
	return HystrixRouter{
		svcMap:       &sync.Map{},
		logger:       logger,
		fallbackMsg:  fallbackMsg,
		consulClient: client,
		tracer:       zipkinTracer,
	}
}

func (router HystrixRouter)ServeHTTP(w http.ResponseWriter, r *http.Request){
     path :=r.URL.Path
     if path ==""{
		 return
	 }
	 pathList := strings.Split(path,"/")
	 var serviceName =""
	 if len(pathList)>1{
	 	serviceName = pathList[1]
	 }
	 if _,ok:=router.svcMap.Load(serviceName);!ok{
	 	  hystrix.ConfigureCommand(serviceName,hystrix.CommandConfig{Timeout: 1000})
	 	  router.svcMap.Store(serviceName, struct {}{})
	 }
	 err := hystrix.Do(serviceName, func() error {
		  result ,_,err := router.consulClient.Catalog().Service(serviceName,"",nil)
		 if err != nil {
			 router.logger.Log("ReverseProxy failed", "query service instace error", err.Error())
			  return err
		 }
		 if len(result) == 0 {
			 router.logger.Log("ReverseProxy failed", "no such service instance", serviceName)
			 return errors.New("no such service instance")
		 }
		 director := func(req *http.Request) {
			 //重新组织请求路径，去掉服务名称部分
			 destPath := strings.Join(pathList[2:], "/")

			 //随机选择一个服务实例
			 tgt := result[rand.Int()%len(result)]
			 router.logger.Log("service id", tgt.ServiceID)

			 //设置代理服务地址信息
			 req.URL.Scheme = "http"
			 req.URL.Host = fmt.Sprintf("%s:%d", tgt.ServiceAddress, tgt.ServicePort)
			 req.URL.Path = "/" + destPath
		 }

		 var proxyError error = nil
		 // 为反向代理增加追踪逻辑，使用如下RoundTrip代替默认Transport
		 roundTrip, _ := zipkinhttpsvr.NewTransport(router.tracer, zipkinhttpsvr.TransportTrace(true))

		 //反向代理失败时错误处理
		 errorHandler := func(ew http.ResponseWriter, er *http.Request, err error) {
			 proxyError = err
		 }

		 proxy := &httputil.ReverseProxy{
			 Director:     director,
			 Transport:    roundTrip,
			 ErrorHandler: errorHandler,
		 }
		 proxy.ServeHTTP(w, r)
		 return proxyError

	 }, func(err error) error {
		 //run执行失败，返回fallback信息
		 router.logger.Log("fallback error description", err.Error())

		 return errors.New(router.fallbackMsg)

	 })
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}
}

