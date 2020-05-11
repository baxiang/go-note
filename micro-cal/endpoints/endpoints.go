package endpoints

import (
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/baxiang/go-note/micro-cal/service"
	"context"
)

type CalRequest struct {
	CalType  string `json:"type"`
	A int `json:"a"`
	B int `json:"b"`
}

type CalResponse struct {
	Data interface{} `json:"data"`
	Code int  `json:"code"`
	Msg string `json:"msg"`
}

func MakeCalEndpoint(s service.CalService)endpoint.Endpoint{
     return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		  req :=  request.(CalRequest)
		  a :=req.A
		  b :=req.B
		  var result int
		  if req.CalType == "add"{
		  	 result = s.Add(a,b)
		  }else if req.CalType=="sub"{
			  result = s.Sub(a,b)
		  }else if req.CalType=="mul"{
			  result = s.Mul(a,b)
		  }else if req.CalType=="div" {
			  result,err = s.Div(a,b)
		  }else {
		  	 return nil,errors.New("请求错误")
		  }
         return CalResponse{
         	Data:result,
         	Code :0,
         	Msg: "success",
		 },nil

	 }
}