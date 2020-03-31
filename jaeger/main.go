package main

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"log"
	"time"
)

func main() {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			BufferFlushInterval: 1 * time.Second,
			LogSpans:            true,
		},
		ServiceName: "hello-world",
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		log.Fatal(err)
	}
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()

	parent := opentracing.GlobalTracer().StartSpan("hello")
	defer parent.Finish()
	child := opentracing.GlobalTracer().StartSpan(
		"world", opentracing.ChildOf(parent.Context()))
	defer child.Finish()

}
