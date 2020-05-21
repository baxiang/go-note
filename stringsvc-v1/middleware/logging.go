package middleware

import (
	"github.com/baxiang/go-note/stringsvc-v1/service"
	"github.com/go-kit/kit/log"
	"time"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next service.StringService
}
func (mw LoggingMiddleware)Uppercase(s string)(output string,err error){
	defer func(begin time.Time) {
		_ =mw.Logger.Log(
			"method", "uppercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
			)
	}(time.Now())
	output, err = mw.Next.Uppercase(s)
	return
}
func (mw LoggingMiddleware) Count(s string) (n int) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method","count",
			"input",s,
			"output",n,
			"took",time.Since(begin),
		)
	}(time.Now())
	n = mw.Next.Count(s)
	return
}