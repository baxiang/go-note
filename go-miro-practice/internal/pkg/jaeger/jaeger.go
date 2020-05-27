package jaeger

import (
	"github.com/google/wire"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-lib/metrics/prometheus"
)

func NewConfiguration(v *viper.Viper, logger *zap.Logger) (*config.Configuration, error) {
	var (
		err error
		c   = new(config.Configuration)
	)

	if err = v.UnmarshalKey("jaeger", c); err != nil {
		return nil, errors.Wrap(err, "unmarshal jaeger configuration error")
	}

	logger.Info("load jaeger configuration success")

	return c, nil
}

func NewJaeger(c *config.Configuration) (opentracing.Tracer, error) {

	metricsFactory := prometheus.New()
	tracer, _, err := c.NewTracer(config.Metrics(metricsFactory))

	if err != nil {
		return nil, errors.Wrap(err, "create jaeger tracer error")
	}

	return tracer, nil
}

var ProviderSet = wire.NewSet(NewJaeger, NewConfiguration)