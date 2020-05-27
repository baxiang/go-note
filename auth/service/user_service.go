package service

import (
	"context"
	"errors"
	"github.com/baxiang/go-note/auth/model"
)

var (
	ErrUserNotExist = errors.New("username is not exist")
	ErrPassword = errors.New("invalid password")
)

type UserDetailService interface {
	GetUserDetailByUserName(ctx context.Context,username,password string)(*model.UserDetails,error)
}

type InMemoryUserDetailService struct {
	userDetailsDict map[string]*model.UserDetails
}
func (service *InMemoryUserDetailService)GetUserDetailByUserName(ctx context.Context,username,password string)(*model.UserDetails,error){
	if  userDetails,ok := service.userDetailsDict[username];ok{
		if userDetails.Password == password{
			return userDetails,nil
		}else {
			return  nil,ErrPassword
		}
	}else {
		return nil,ErrUserNotExist
	}
}

func NewInMemoryUserDetailsService(userDetailsList []*model.UserDetails)*InMemoryUserDetailService{
	userDetailsDict :=make(map[string]*model.UserDetails)
	if len(userDetailsDict)>0{
		for _,value :=range userDetailsList{
			userDetailsDict[value.Username] = value
		}
	}
	return &InMemoryUserDetailService{
		userDetailsDict: userDetailsDict,
	}
}