package middleware


import (
	"github.com/baxiang/go-note/go-kit-trace/service"
	"github.com/go-kit/kit/log"
	"time"
)

// loggingMiddleware Make a new type
// that contains Service interface and logger instance
type loggingMiddleware struct {
	service.CalService
	logger log.Logger
}

// LoggingMiddleware make logging middleware
func LoggingMiddleware(logger log.Logger) service.ServiceMiddleware {
	return func(next service.CalService) service.CalService {
		return loggingMiddleware{next, logger}
	}
}

func (mw loggingMiddleware) Add(a, b int) (ret int) {

	defer func(beign time.Time) {
		mw.logger.Log(
			"function", "Add",
			"a", a,
			"b", b,
			"result", ret,
			"took", time.Since(beign),
		)
	}(time.Now())

	ret = mw.CalService.Add(a, b)
	return ret
}

func (mw loggingMiddleware) Subtract(a, b int) (ret int) {

	defer func(beign time.Time) {
		mw.logger.Log(
			"function", "Subtract",
			"a", a,
			"b", b,
			"result", ret,
			"took", time.Since(beign),
		)
	}(time.Now())

	ret = mw.CalService.Subtract(a, b)
	return ret
}

func (mw loggingMiddleware) Multiply(a, b int) (ret int) {

	defer func(beign time.Time) {
		mw.logger.Log(
			"function", "Multiply",
			"a", a,
			"b", b,
			"result", ret,
			"took", time.Since(beign),
		)
	}(time.Now())

	ret = mw.CalService.Multiply(a, b)
	return ret
}

func (mw loggingMiddleware) Divide(a, b int) (ret int, err error) {

	defer func(beign time.Time) {
		mw.logger.Log(
			"function", "Divide",
			"a", a,
			"b", b,
			"result", ret,
			"took", time.Since(beign),
		)
	}(time.Now())

	ret, err = mw.CalService.Divide(a, b)
	return
}

func (mw loggingMiddleware) HealthCheck() (result bool) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "HealthChcek",
			"result", result,
			"took", time.Since(begin),
		)
	}(time.Now())
	result = mw.CalService.HealthCheck()
	return
}