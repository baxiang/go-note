package endpoints

import (
	"context"
	"github.com/baxiang/go-note/stringsvc-v1/service"
	"github.com/go-kit/kit/endpoint"
)

type UppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct {
	V string `json:"v"`
}

type CountRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}
type commonResponse struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

func MakeUppercaseEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return commonResponse{Data: nil, Message: err.Error(), Code: 1000}, nil
		}
		return commonResponse{Data: uppercaseResponse{v}, Message: "success", Code: 0}, nil
	}
}

func MakeCountEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CountRequest)
		v := svc.Count(req.S)
		return commonResponse{Data: countResponse{v}, Message: "", Code: 0}, nil
	}
}
