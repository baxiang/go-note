package router

import (
	"context"
	"github.com/baxiang/go-note/micro-cal/transports"
	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func MakeHttpHandler(ctx context.Context, endpoint endpoint.Endpoint,) http.Handler {
	r := mux.NewRouter()
	r.Methods("POST").Path("/cal/{type}").Handler(kithttp.NewServer(
		endpoint,
		transports.DecodeCalRequest,
		transports.EncodeCalResponse,
	))
	return r
}