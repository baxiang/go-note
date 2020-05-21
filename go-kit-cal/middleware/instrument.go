package middleware

import (
	"github.com/baxiang/go-note/go-kit-trace/service"
	"github.com/go-kit/kit/metrics"
	"time"
)



// metricMiddleware 定义监控中间件，嵌入Service
// 新增监控指标项：requestCount和requestLatency
type metricMiddleware struct {
	service.CalService
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
}

// Metrics 封装监控方法
func Metrics(requestCount metrics.Counter, requestLatency metrics.Histogram) service.ServiceMiddleware {
	return func(next service.CalService) service.CalService {
		return metricMiddleware{
			next,
			requestCount,
			requestLatency}
	}
}

func (mw metricMiddleware) Add(a, b int) (ret int) {

	defer func(beign time.Time) {
		lvs := []string{"method", "Add"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(beign).Seconds())
	}(time.Now())

	ret = mw.CalService.Add(a, b)
	return ret
}

func (mw metricMiddleware) Subtract(a, b int) (ret int) {

	defer func(beign time.Time) {
		lvs := []string{"method", "Subtract"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(beign).Seconds())
	}(time.Now())

	ret = mw.CalService.Subtract(a, b)
	return ret
}

func (mw metricMiddleware) Multiply(a, b int) (ret int) {

	defer func(beign time.Time) {
		lvs := []string{"method", "Multiply"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(beign).Seconds())
	}(time.Now())

	ret = mw.CalService.Multiply(a, b)
	return ret
}

func (mw metricMiddleware) Divide(a, b int) (ret int, err error) {

	defer func(beign time.Time) {
		lvs := []string{"method", "Divide"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(beign).Seconds())
	}(time.Now())

	ret, err = mw.CalService.Divide(a, b)
	return
}

func (mw metricMiddleware) HealthCheck() (result bool) {

	defer func(begin time.Time) {
		lvs := []string{"method", "HealthCheck"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	result = mw.CalService.HealthCheck()
	return
}