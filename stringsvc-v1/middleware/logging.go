package middleware

import (
	"github.com/baxiang/go-note/stringsvc-v1/service"
	"github.com/go-kit/kit/log"
	"time"
)

func LoggingMiddleware(logger log.Logger)service.ServiceMiddleware{
	return func(next service.StringService) service.StringService {
		 return Logmw{
			 Logger:        logger,
			 StringService: next,
		 }
	}
}

type Logmw struct {
	Logger log.Logger
	service.StringService
}
func (mw Logmw)Uppercase(s string)(output string,err error){
	defer func(begin time.Time) {
		_ =mw.Logger.Log(
			"method", "uppercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
			)
	}(time.Now())
	output, err = mw.StringService.Uppercase(s)
	return
}
func (mw Logmw) Count(s string) (n int) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method","count",
			"input",s,
			"output",n,
			"took",time.Since(begin),
		)
	}(time.Now())
	n = mw.StringService.Count(s)
	return
}