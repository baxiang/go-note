package transport

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/baxiang/go-note/discover/endpoint"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	ErrorBadRequest = errors.New("invalid request parameter")
)

func MakeHttpHandler(ctx context.Context,endpoints endpoint.DiscoveryEndpoints,logger log.Logger)http.Handler{
	r := mux.NewRouter()

	options := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	r.Methods("GET").Path("/hello").Handler(kithttp.NewServer(
		endpoints.SayHelloEndpoint,
		decodeSayHelloRequest,
		encodeJsonResponse,
		options...
		))
	r.Methods("GET").Path("/discovery").Handler(kithttp.NewServer(
		endpoints.DiscoveryEndpoint,
		decodeDiscoveryRequest,
		encodeJsonResponse,
		options...
	))
	r.Methods("GET").Path("/health").Handler(kithttp.NewServer(
		endpoints.HealthCheckEndpoint,
		decodeHeathCheckRequest,
		encodeJsonResponse,
		options...
	))

	return r
}

func decodeSayHelloRequest(_ context.Context,r *http.Request)(interface{},error){
    return endpoint.SayHelloRequest{},nil
}
func decodeDiscoveryRequest(_ context.Context,r *http.Request)(interface{},error){
	serviceName := r.URL.Query().Get("name")
    if serviceName == ""{
    	return nil,ErrorBadRequest
	}
	return endpoint.DiscoveryRequest{ServiceName: serviceName},nil
}

func decodeHeathCheckRequest(_ context.Context,r *http.Request)(interface{},error){
	return endpoint.HealthRequest{},nil
}


func encodeJsonResponse(_ context.Context, w http.ResponseWriter,response interface{})error{
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
    return  json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context,err error, w http.ResponseWriter){
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	switch err {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}