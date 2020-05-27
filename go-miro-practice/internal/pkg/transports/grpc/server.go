package grpc

import (
	consulApi "github.com/hashicorp/consul/api"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
)

type ServerOptions struct {
	Port int
}

func NewServerOptions(v *viper.Viper) (*ServerOptions, error) {
	var (
		err error
		o   = new(ServerOptions)
	)
	if err = v.UnmarshalKey("grpc", o); err != nil {
		return nil, err
	}

	return o, nil
}

type Server struct {
	o         *ServerOptions
	app       string
	host      string
	port      int
	logger    *zap.Logger
	server    *grpc.Server
	consulCli *consulApi.Client
}

type InitServers func(s *grpc.Server)

func NewGrpcServer(o *ServerOptions, logger *zap.Logger, init InitServers, consulCli *consulApi.Client, tracer opentracing.Tracer) (*Server, error) {
	// initialize grpc server
	var gs *grpc.Server
	logger = logger.With(zap.String("type", "grpc"))
	{
		grpc_prometheus.EnableHandlingTimeHistogram()
		gs = grpc.NewServer(
			grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
				grpc_ctxtags.StreamServerInterceptor(),
				grpc_prometheus.StreamServerInterceptor,
				grpc_zap.StreamServerInterceptor(logger),
				grpc_recovery.StreamServerInterceptor(),
				otgrpc.OpenTracingStreamServerInterceptor(tracer),
			)),
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
				grpc_ctxtags.UnaryServerInterceptor(),
				grpc_prometheus.UnaryServerInterceptor,
				grpc_zap.UnaryServerInterceptor(logger),
				grpc_recovery.UnaryServerInterceptor(),
				otgrpc.OpenTracingServerInterceptor(tracer),
			)),
		)
		init(gs)
	}

	return &Server{
		o:         o,
		logger:    logger.With(zap.String("type", "grpc.Server")),
		server:    gs,
		consulCli: consulCli,
	}, nil
}

