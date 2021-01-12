package main

import (
	"context"
	"github.com/baxiang/go-note/tag-grpc-service/global"
	"github.com/baxiang/go-note/tag-grpc-service/internal/middleware"
	"github.com/baxiang/go-note/tag-grpc-service/pkg/tracer"
	pb "github.com/baxiang/go-note/tag-grpc-service/proto"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
)

func setupTracer()error{
	jaegerTracer, _, err := tracer.NewJaegerTracer("article-service",
		"127.0.0.1:6831")
	if err!=nil{
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}

func main() {
	err := setupTracer()
	if err!= nil{
		log.Fatalf("tracer error :%v",err)
	}
	ctx := metadata.AppendToOutgoingContext(context.Background(), "hello", "world")
	conn, err := GetClientConn(ctx, "127.0.0.1:8004", []grpc.DialOption{
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
			middleware.UnaryContextTimeout(),
			middleware.ClientTracing(),
		)),
	})
	if err!= nil{
		log.Fatalf("getClientConn error :%v",err)
	}
	defer conn.Close()
	client := pb.NewTagServiceClient(conn)
	resp, err := client.GetTagList(ctx, &pb.GetTagListReq{
		Name: "tag",
	})
	if err != nil {
		log.Fatalf("tagServiceClient.GetTagList err: %v", err)
	}
	log.Printf("resp: %v\n", resp)
}
func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, target, opts...)
}
