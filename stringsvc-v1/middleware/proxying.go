package middleware

import (
	"context"
	"fmt"
	"github.com/baxiang/go-note/stringsvc-v1/endpoints"
	"github.com/baxiang/go-note/stringsvc-v1/service"
	transport "github.com/baxiang/go-note/stringsvc-v1/transports"
	"github.com/go-kit/kit/endpoint"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/ratelimit"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/url"
	"strings"
	"time"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/sony/gobreaker"
	"golang.org/x/time/rate"

)



func ProxyingMiddleware(ctx context.Context,instances string,logger log.Logger)service.ServiceMiddleware{
	if instances == "" {
		logger.Log("proxy_to", "none")
		return func(next service.StringService) service.StringService { return next }
	}

	// Set some parameters for our client.
	var (
		qps         = 100                    // beyond which we will return an error
		maxAttempts = 3                      // per request, before giving up
		maxTime     = 250 * time.Millisecond // wallclock time, before giving up
	)
	var (
		instanceList = split(instances)
		endpointer   sd.FixedEndpointer
	)
	logger.Log("proxy_to", fmt.Sprint(instanceList))
	for _, instance := range instanceList {
		var e endpoint.Endpoint
		e = makeUppercaseProxy(ctx, instance)
		e = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(e)
		e = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), qps))(e)
		endpointer = append(endpointer, e)
	}

	// Now, build a single, retrying, load-balancing endpoint out of all of
	// those individual endpoints.

	balancer := lb.NewRoundRobin(endpointer)
	retry := lb.Retry(maxAttempts, maxTime, balancer)

	// And finally, return the ServiceMiddleware, implemented by proxymw.
	return func(next service.StringService) service.StringService {
		return proxymw{ctx, next, retry}
	}
}

type proxymw struct {
	ctx context.Context
	next service.StringService
	uppercase endpoint.Endpoint
}

func(mw proxymw)Count(s string)int{
	return mw.next.Count(s)
}

func (mw proxymw) Uppercase(s string) (string, error) {
	response, err := mw.uppercase(mw.ctx, endpoints.UppercaseRequest{S: s})
	if err != nil {
		return "", err
	}

	resp := response.(endpoints.CommonResponse)
	//if resp.Err != "" {
	//	return resp.V, errors.New(resp.Err)
	//}
	return resp.Message, nil

	//return mw.next.Uppercase(s)
}

func makeUppercaseProxy(ctx context.Context,instance string)endpoint.Endpoint{
	if !strings.HasPrefix(instance,"http"){
		instance = "http://"+instance
	}
	u,err :=url.Parse(instance)
	if err!=nil{

	}
	if u.Path ==""{
		u.Path = "/uppercase"
	}
	return httptransport.NewClient(
		"POST",
		u,
		transport.EncodeRequest,
		transport.DecodeUppercaseResponse,
		).Endpoint()

}
func split(s string) []string {
	a := strings.Split(s, ",")
	for i := range a {
		a[i] = strings.TrimSpace(a[i])
	}
	return a
}