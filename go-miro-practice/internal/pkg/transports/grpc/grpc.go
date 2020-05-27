package grpc

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewServerOptions, NewGrpcServer,NewClientOptions,NewGrpcClient)