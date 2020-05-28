package main

import (
	"github.com/baxiang/go-miro-practice/internal/app/reviews/repositories"
	"github.com/baxiang/go-miro-practice/internal/app/reviews/services"
	"github.com/baxiang/go-miro-practice/internal/pkg/app"
	"github.com/baxiang/go-miro-practice/internal/pkg/config"
	"github.com/baxiang/go-miro-practice/internal/pkg/consul"
	"github.com/baxiang/go-miro-practice/internal/pkg/database"
	"github.com/baxiang/go-miro-practice/internal/pkg/jaeger"
	"github.com/baxiang/go-miro-practice/internal/pkg/log"
	"github.com/baxiang/go-miro-practice/internal/pkg/transports/grpc"
	"github.com/baxiang/go-miro-practice/internal/pkg/transports/http"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	services.ProviderSet,
	consul.ProviderSet,
	jaeger.ProviderSet,
	http.ProviderSet,
	grpc.ProviderSet,
	//reviews.ProviderSet,
	repositories.ProviderSet,
	//controllers.ProviderSet,
	//grpcservers.ProviderSet,
)

func CreateApp(cf string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}