package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/baxiang/go-note/stringsvc-v1/endpoints"
	"io/ioutil"
	"net/http"
)

func DecodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoints.UppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
func DecodeUppercaseResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var response endpoints.CommonResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}
func DecodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoints.CountRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func EncodeRequest(_ context.Context, r *http.Request,request interface{})error{
	var buf bytes.Buffer
	if err:=json.NewEncoder(&buf).Encode(request);err!=nil{
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}
