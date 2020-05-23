package main

import (
	"context"
	"flag"
	"github.com/baxiang/go-note/stringsvc-v1/endpoints"
	"github.com/baxiang/go-note/stringsvc-v1/middleware"
	"github.com/baxiang/go-note/stringsvc-v1/service"
	transport "github.com/baxiang/go-note/stringsvc-v1/transports"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/kit/log"
	"os"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"net/http"
)

func main() {
	var (
		listen = flag.String("listen", ":8080", "HTTP listen address")
		proxy  = flag.String("proxy", "", "Optional comma-separated list of URLs to proxy uppercase requests")
	)
	flag.Parse()

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "listen", *listen, "caller", log.DefaultCaller)

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "str",
		Subsystem: "string_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)

	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "str",
		Subsystem: "string_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "str",
		Subsystem: "string_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{}) // no fields here

	var svc service.StringService
	svc = service.StringServiceImpl{}
	svc = middleware.ProxyingMiddleware(context.Background(),*proxy,logger)(svc)
	svc = middleware.LoggingMiddleware(logger)(svc)
	svc = middleware.InstrumentingMiddleware(requestCount, requestLatency, countResult)(svc)
	uppercaseHandler := kithttp.NewServer(
		endpoints.MakeUppercaseEndpoint(svc),
		transport.DecodeUppercaseRequest,
		transport.EncodeResponse)

	countHandler := kithttp.NewServer(
		endpoints.MakeCountEndpoint(svc),
		transport.DecodeCountRequest,
		transport.EncodeResponse)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	http.Handle("/metrics", promhttp.Handler())
	logger.Log(http.ListenAndServe(*listen, nil))
}
