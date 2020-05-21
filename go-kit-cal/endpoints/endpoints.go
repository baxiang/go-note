package endpoints


import (
	"context"
	"errors"
	"github.com/baxiang/go-note/go-kit-trace/service"
	"github.com/go-kit/kit/endpoint"
	"github.com/juju/ratelimit"
	"golang.org/x/time/rate"
	"strings"
)

// CalculateEndpoint define endpoint
type CalEndpoints struct {
	CalEndpoint  endpoint.Endpoint
	HealthCheckEndpoint endpoint.Endpoint
}

var (
	ErrInvalidRequestType = errors.New("RequestType has only four type: Add,Subtract,Multiply,Divide")
)

// ArithmeticRequest define request struct
type CalRequest struct {
	RequestType string `json:"request_type"`
	A           int    `json:"a"`
	B           int    `json:"b"`
}

// ArithmeticResponse define response struct
type CalResponse struct {
	Result int   `json:"result"`
	Error  error `json:"error"`
}

// MakeArithmeticEndpoint make endpoint
func MakeCalEndpoint(svc service.CalService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CalRequest)

		var (
			res, a, b int
			calError  error
		)

		a = req.A
		b = req.B

		if strings.EqualFold(req.RequestType, "Add") {
			res = svc.Add(a, b)
		} else if strings.EqualFold(req.RequestType, "Subtract") {
			res = svc.Subtract(a, b)
		} else if strings.EqualFold(req.RequestType, "Multiply") {
			res = svc.Multiply(a, b)
		} else if strings.EqualFold(req.RequestType, "Divide") {
			res, calError = svc.Divide(a, b)
		} else {
			return nil, ErrInvalidRequestType
		}

		return CalResponse{Result: res, Error: calError}, nil
	}
}

// HealthRequest 健康检查请求结构
type HealthRequest struct{}

// HealthResponse 健康检查响应结构
type HealthResponse struct {
	Status bool `json:"status"`
}

// MakeHealthCheckEndpoint 创建健康检查Endpoint
func MakeHealthCheckEndpoint(svc service.CalService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		status := svc.HealthCheck()
		return HealthResponse{status}, nil
	}
}

var ErrLimitExceed = errors.New("Rate limit exceed!")

// NewTokenBucketLimitterWithJuju 使用juju/ratelimit创建限流中间件
func MakeTokenBucketLimiterWithJuju(bkt *ratelimit.Bucket) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if bkt.TakeAvailable(1) == 0 {
				return nil, ErrLimitExceed
			}
			return next(ctx, request)
		}
	}
}

// NewTokenBucketLimitterWithBuildIn 使用x/time/rate创建限流中间件
func MakeTokenBucketLimiter(bkt *rate.Limiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !bkt.Allow() {
				return nil, ErrLimitExceed
			}
			return next(ctx, request)
		}
	}
}
