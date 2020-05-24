package service

import (
	"context"
	"github.com/baxiang/go-note/auth/model"
	"net/http"
	"errors"
)

var (

	ErrNotSupportGrantType = errors.New("grant type is not supported")
	ErrNotSupportOperation = errors.New("no support operation")
	ErrInvalidUsernameAndPasswordRequest      = errors.New("invalid username, password")
	ErrInvalidTokenRequest = errors.New("invalid token")
	ErrExpiredToken        = errors.New("token is expired")


)
type TokenGranter interface {
	Grant(ctx context.Context,grantType string,client *model.ClientDetails,req *http.Request)(*model.OAuth2Token,error)
}

type ComposeTokenGranter struct {
	TokenGrantDict map[string] TokenGranter
}

func NewComposeTokenGranter(tokenGrantDict map[string]TokenGranter)TokenGranter{
	return &ComposeTokenGranter{
		TokenGrantDict: tokenGrantDict,
	}
}

func(tokenGranter *ComposeTokenGranter)Grant(ctx context.Context,
	grantType string,client *model.ClientDetails,req *http.Request)(*model.OAuth2Token,error){
	granter := tokenGranter.TokenGrantDict[grantType]
	if granter==nil{
		return nil,ErrNotSupportGrantType
	}
	return granter.Grant(ctx, grantType, client, req)
}


type UsernamePasswordTokenGranter struct {
	supportGrantType string
	userDetailsService UserDetailService
	tokenService TokenService
}

func(tokenGranter *UsernamePasswordTokenGranter)Grant(ctx context.Context,
	grantType string,
	client *model.ClientDetails,
	req *http.Request)(*model.OAuth2Token,error){
	if grantType !=tokenGranter.supportGrantType{
		return nil,ErrNotSupportGrantType
	}
	userName := req.FormValue("username")
	password := req.FormValue("password")

	if len(userName)==0||len(password)==0{
		return nil, ErrInvalidUsernameAndPasswordRequest
	}

	userDetails, err := tokenGranter.userDetailsService.GetUserDetailByUserName(ctx, userName, password)
	if err!=nil{
		return nil, ErrInvalidUsernameAndPasswordRequest
	}

	return tokenGranter.tokenService.CreateAccessToken(&model.OAuth2Details{
		Client: client,
		User:   userDetails,
	})
}

//刷新令牌

type RefreshTokenGranter struct {
	supportGrantType string
	tokenService TokenService
}

func(tokenGranter *RefreshTokenGranter)Grant(ctx context.Context,
	grantType string,
	client *model.ClientDetails,
	req *http.Request)(*model.OAuth2Token,error){
	if grantType !=tokenGranter.supportGrantType{
		return nil,ErrNotSupportGrantType
	}
	refreshTokenValue := req.URL.Query().Get("refresh_token")
	if len(refreshTokenValue)==0{
		return  nil,ErrInvalidTokenRequest
	}
	return tokenGranter.tokenService.RefreshAccessToken(refreshTokenValue)
}





type TokenService interface {
	// 根据访问令牌获取对应的用户信息和客户端信息
	GetOAuth2DetailsByAccessToken(tokenValue string)(*model.OAuth2Details,error)

	// 根据用户信息和客户端信息生成访问令牌
	CreateAccessToken(oauth2Details *model.OAuth2Details)(*model.OAuth2Token,error)

	// 根据刷新令牌获取访问令牌
	RefreshAccessToken(refreshToken string)(*model.OAuth2Token,error)

	//根据用户信息和客户端信息获取已生成访问令牌
	GetAccessToken(details *model.OAuth2Details)(*model.OAuth2Token,error)

	// 根据访问令牌获取访问令牌结构体
	ReadAccessToken(tokenValue string)(*model.OAuth2Token,error)

}