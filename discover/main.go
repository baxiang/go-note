package main

import (
	"flag"
	"context"
	"fmt"
	"github.com/baxiang/go-note/discover/config"
	"github.com/baxiang/go-note/discover/discover"
	"github.com/baxiang/go-note/discover/endpoint"
	"github.com/baxiang/go-note/discover/service"
	"github.com/baxiang/go-note/discover/transport"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		// 服务地址和服务名
		servicePort = flag.String("service.port", "8080", "service port")
		serviceHost = flag.String("service.host", "192.168.1.112", "service host")
		serviceName = flag.String("service.name", "SayHello", "service name")
		// consul 地址
		consulHost = flag.String("consul.host", "127.0.0.1", "consul host")
		consulPort = flag.String("consul.port", "8500", "consul port")

	)
	flag.Parse()




	ctx := context.Background()
	errChan := make(chan error)

	// 声明服务发现客户端

	var discoveryClient discover.DiscoveryClient
	discoveryClient,_= discover.NewKitDiscoverClient(*consulHost, *consulPort)

	svc := service.NewDiscoveryServiceImpl(discoveryClient)

	helloEndpoint := endpoint.MakeSayHelloEndpoint(svc)
	discoveryEndpoint := endpoint.MakeDiscoveryEndpoint(svc)
	checkEndpoint := endpoint.MakeHealthCheckEndpoint(svc)

	endpoints := endpoint.DiscoveryEndpoints{
		SayHelloEndpoint:    helloEndpoint,
		DiscoveryEndpoint:   discoveryEndpoint,
		HealthCheckEndpoint: checkEndpoint,
	}

	r := transport.MakeHttpHandler(ctx, endpoints, config.KitLogger)

	instanceId := fmt.Sprintf("%s-%s",*serviceName,uuid.NewV4().String())
    config.Logger.Println(instanceId)
	go func() {
		if !discoveryClient.Register(*serviceName,
			instanceId,
			*serviceHost,
			*servicePort,
			"/health",
			nil,
			config.Logger,
			){
			config.Logger.Printf("string-service for service %s failed.", *serviceName)
			// 注册失败，服务启动失败
			os.Exit(-1)
		}
		handler := r
		errChan <- http.ListenAndServe(":"  + *servicePort, handler)
	}()

	go func() {
		// 监控系统信号，等待 ctrl + c 系统信号通知服务关闭
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
	error := <-errChan
	//服务退出取消注册
	discoveryClient.DeRegister(instanceId, config.Logger)
	config.Logger.Println(error)
}