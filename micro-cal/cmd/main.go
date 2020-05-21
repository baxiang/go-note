package main

import (
	"fmt"
	"github.com/baxiang/go-note/micro-cal/router"
	"github.com/baxiang/go-note/micro-cal/service"
	"net/http"
	"os"
	"context"
	"os/signal"
	"syscall"
	"github.com/baxiang/go-note/micro-cal/endpoints"
)

func main() {
	errChan := make(chan error)
	svc := service.Calculator{}
	endpoint := endpoints.MakeCalEndpoint(svc)

	handler := router.MakeHttpHandler(context.Background(), endpoint)
	go func() {
		errChan <- http.ListenAndServe(":9000", handler)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	fmt.Println(<-errChan)

}
