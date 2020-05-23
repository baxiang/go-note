package service

import (
	"context"
	"github.com/baxiang/go-note/discover/config"
	"github.com/baxiang/go-note/discover/discover"
	"errors"
)

type Service interface {
	//健康检查
	HealthCheck() bool

	SayHello()string

	DiscoveryService(ctx context.Context,
		serviceName string,
	)([]interface{},error)
}

var ErrNotServiceInstances = errors.New("instances are not existed")


type DiscoveryServiceImpl struct {
	discoveryClient discover.DiscoveryClient
}

func NewDiscoveryServiceImpl(discoveryClient discover.DiscoveryClient)Service{
	return &DiscoveryServiceImpl{discoveryClient: discoveryClient}
}
func (*DiscoveryServiceImpl) SayHello() string {
	return "Hello World!"
}

func (service *DiscoveryServiceImpl) DiscoveryService(_ context.Context, serviceName string) ([]interface{}, error)  {

	instances := service.discoveryClient.DiscoverServices(serviceName, config.Logger)

	if instances == nil || len(instances) == 0 {
		return nil, ErrNotServiceInstances
	}
	return instances, nil
}


// HealthCheck implement Service method
// 用于检查服务的健康状态，这里仅仅返回true
func (*DiscoveryServiceImpl) HealthCheck() bool {
	return true
}
