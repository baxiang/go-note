package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/baxiang/go-note/go-kit-trace/discover"
	"github.com/baxiang/go-note/go-kit-trace/endpoints"
	"github.com/baxiang/go-note/go-kit-trace/middleware"
	"github.com/baxiang/go-note/go-kit-trace/service"
	"github.com/baxiang/go-note/go-kit-trace/transports"
	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	kitzipkin "github.com/go-kit/kit/tracing/zipkin"
	"github.com/openzipkin/zipkin-go"
	stdprometheus "github.com/prometheus/client_golang/prometheus"

	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
	"golang.org/x/time/rate"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	var (
		hostIP ="192.168.43.53"
		consulHost  = flag.String("consul.host", hostIP, "consul ip address")
		consulPort  = flag.String("consul.port", "8500", "consul port")
		serviceHost = flag.String("service.host", hostIP, "service ip address")
		servicePort = flag.String("service.port", "9000", "service port")
		zipkinURL   = flag.String("zipkin.url", fmt.Sprintf("http://%s:9411/api/v2/spans",hostIP),
			"Zipkin server url")
	)

	flag.Parse()

	ctx := context.Background()
	errChan := make(chan error)

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	fieldKeys := []string{"method"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "go_kit",
		Subsystem: "calculate_service",
		Name:      "cal_request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)

	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "go_kit",
		Subsystem: "calculate_service",
		Name:      "cal_request_latency",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)

	var zipkinTracer *zipkin.Tracer
	{
		var (
			err           error
			hostPort      = *serviceHost + ":" + *servicePort
			serviceName   = "calculate-service"
			useNoopTracer = (*zipkinURL == "")
			reporter      = zipkinhttp.NewReporter(*zipkinURL)
		)
		defer reporter.Close()
		zEP, _ := zipkin.NewEndpoint(serviceName, hostPort)
		zipkinTracer, err = zipkin.NewTracer(
			reporter, zipkin.WithLocalEndpoint(zEP), zipkin.WithNoopTracer(useNoopTracer),
		)
		if err != nil {
			logger.Log("err", err)
			os.Exit(1)
		}
		if !useNoopTracer {
			logger.Log("tracer", "Zipkin", "type", "Native", "URL", *zipkinURL)
		}
	}

	//add ratelimit,refill every second,set capacity 3
	ratebucket := rate.NewLimiter(rate.Every(time.Second*1), 100)

	var svc service.CalService
	svc = service.CalServiceImpl{}

	// add logging middleware to service
	svc = middleware.LoggingMiddleware(logger)(svc)
	svc = middleware.Metrics(requestCount, requestLatency)(svc)



	endpoint := endpoints.MakeCalEndpoint(svc)
	endpoint = endpoints.MakeTokenBucketLimiter(ratebucket)(endpoint)
	endpoint = kitzipkin.TraceEndpoint(zipkinTracer, "calculate-endpoint")(endpoint)

	//创建健康检查的Endpoint
	healthEndpoint := endpoints.MakeHealthCheckEndpoint(svc)
	healthEndpoint = endpoints.MakeTokenBucketLimiter(ratebucket)(healthEndpoint)
	healthEndpoint = kitzipkin.TraceEndpoint(zipkinTracer, "health-endpoint")(healthEndpoint)

	//把算术运算Endpoint和健康检查Endpoint封装至ArithmeticEndpoints
	endpts := endpoints.CalEndpoints{
		CalEndpoint:  endpoint,
		HealthCheckEndpoint: healthEndpoint,
	}

	//创建http.Handler

	r := transports.MakeHttpHandler(ctx, endpts, zipkinTracer, logger)


	//创建注册对象
	registar := discover.Register(*consulHost, *consulPort, *serviceHost, *servicePort, logger)

	go func() {
		fmt.Println("Http Server start at port:" + *servicePort)
		//启动前执行注册
		registar.Register()
		handler := r
		errChan <- http.ListenAndServe(":"+*servicePort, handler)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	error := <-errChan
	//服务退出取消注册
	registar.Deregister()
	fmt.Println(error)
}
