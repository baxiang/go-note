package middleware

import (
	"fmt"
	"github.com/baxiang/go-note/stringsvc-v1/service"
	"github.com/go-kit/kit/metrics"
	"time"
)

func InstrumentingMiddleware(
	requestCount metrics.Counter,
	requestLatency metrics.Histogram,
	countResult metrics.Histogram,
	)service.ServiceMiddleware{
	   return func(next service.StringService) service.StringService {
	   	  return  Instrumentingmw{
			  RequestCount:   requestCount,
			  RequestLatency: requestLatency,
			  CountResult:    countResult,
			  StringService:  next,
		  }
	   }
}


type Instrumentingmw struct {
	RequestCount metrics.Counter
	RequestLatency metrics.Histogram
	CountResult metrics.Histogram
	service.StringService
}

func (mw Instrumentingmw)Uppercase(s string)(output string, err error){
	defer func(begin time.Time) {
		lvs :=[]string{"method","uppercase","error",fmt.Sprint(err!=nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	output, err = mw.StringService.Uppercase(s)
	return
}
func (mw Instrumentingmw) Count(s string) (n int) {
	defer func(begin time.Time) {
		lvs := []string{"method", "count", "error", "false"}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
		mw.CountResult.Observe(float64(n))
	}(time.Now())

	n = mw.StringService.Count(s)
	return
}