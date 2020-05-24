package service

import (
	"context"
	"errors"
	"github.com/baxiang/go-note/auth/model"
)

var (

	ErrClientNotExist = errors.New("clientId is not exist")
	ErrClientSecret = errors.New("invalid clientSecret")

)

// Service Define a service interface
type ClientDetailsService interface {

	GetClientDetailByClientId(ctx context.Context, clientId string, clientSecret string)(*model.ClientDetails, error)

}

type InMemoryClientDetailsService struct {
	clientDetailsDict map[string]*model.ClientDetails

}

func NewInMemoryClientDetailService(clientDetailsList []*model.ClientDetails ) *InMemoryClientDetailsService{
	clientDetailsDict := make(map[string]*model.ClientDetails)

	if clientDetailsList != nil {
		for _, value := range clientDetailsList {
			clientDetailsDict[value.ClientId] = value
		}
	}

	return &InMemoryClientDetailsService{
		clientDetailsDict:clientDetailsDict,
	}
}


func (service *InMemoryClientDetailsService)GetClientDetailByClientId(ctx context.Context, clientId string, clientSecret string)(*model.ClientDetails, error) {

	// 根据 clientId 获取 clientDetails
	if clientDetails, ok := service.clientDetailsDict[clientId]; ok{
		// 比较 clientSecret 是否正确
		if clientDetails.ClientSecret == clientSecret{
			return clientDetails, nil
		}else {
			return nil, ErrClientSecret
		}
	}else {
		return nil, ErrClientNotExist
	}



}


