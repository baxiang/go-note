package service

import (
	"context"
	"github.com/baxiang/go-note/auth/model"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"errors"
	"strconv"
	"time"
	uuid "github.com/satori/go.uuid"
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
	// 根据用户信息和客户端信息生成访问令牌
	CreateAccessToken(oauth2Details *model.OAuth2Details)(*model.OAuth2Token,error)

	// 根据访问令牌获取对应的用户信息和客户端信息
	GetOAuth2DetailsByAccessToken(tokenValue string)(*model.OAuth2Details,error)


	// 根据刷新令牌获取访问令牌
	RefreshAccessToken(refreshToken string)(*model.OAuth2Token,error)

	//根据用户信息和客户端信息获取已生成访问令牌
	GetAccessToken(details *model.OAuth2Details)(*model.OAuth2Token,error)

	// 根据访问令牌获取访问令牌结构体
	ReadAccessToken(tokenValue string)(*model.OAuth2Token,error)

}

type DefaultTokenService struct {
	tokenStore TokenStore
	tokenEnhancer TokenEnhancer
}

func NewTokenService(tokenStore TokenStore, tokenEnhancer TokenEnhancer) TokenService {
	return &DefaultTokenService{
		tokenStore:tokenStore,
		tokenEnhancer:tokenEnhancer,
	}
}


func (tokenService *DefaultTokenService) CreateAccessToken(oauth2Details *model.OAuth2Details) (*model.OAuth2Token, error) {

	existToken, err := tokenService.tokenStore.GetAccessToken(oauth2Details)
	var refreshToken *model.OAuth2Token
	if err == nil{
		// 存在未失效访问令牌，直接返回
		if !existToken.IsExpired(){
			tokenService.tokenStore.StoreAccessToken(existToken, oauth2Details)
			return existToken, nil

		}
		// 访问令牌已失效，移除
		tokenService.tokenStore.RemoveAccessToken(existToken.TokenValue)
		if existToken.RefreshToken != nil {
			refreshToken = existToken.RefreshToken
			tokenService.tokenStore.RemoveRefreshToken(refreshToken.TokenType)
		}
	}

	if refreshToken == nil || refreshToken.IsExpired(){
		refreshToken, err = tokenService.createRefreshToken(oauth2Details)
		if err != nil{
			return nil, err
		}
	}

	// 生成新的访问令牌
	accessToken, err := tokenService.createAccessToken(refreshToken, oauth2Details)
	if err == nil{
		// 保存新生成令牌
		tokenService.tokenStore.StoreAccessToken(accessToken, oauth2Details)
		tokenService.tokenStore.StoreRefreshToken(refreshToken, oauth2Details)
	}
	return accessToken, err

}

func (tokenService *DefaultTokenService) createAccessToken(refreshToken *model.OAuth2Token,
	oauth2Details *model.OAuth2Details) (*model.OAuth2Token, error) {

	validitySeconds := oauth2Details.Client.AccessTokenValiditySeconds
	s, _ := time.ParseDuration(strconv.Itoa(validitySeconds) + "s")
	expiredTime := time.Now().Add(s)
	accessToken := &model.OAuth2Token{
		RefreshToken:refreshToken,
		ExpiresTime:&expiredTime,
		TokenValue:uuid.NewV4().String(),
	}

	if tokenService.tokenEnhancer != nil{
		return tokenService.tokenEnhancer.Enhance(accessToken, oauth2Details)
	}
	return accessToken, nil
}

func (tokenService *DefaultTokenService) createRefreshToken(oauth2Details *model.OAuth2Details) (*model.OAuth2Token, error) {
	validitySeconds := oauth2Details.Client.RefreshTokenValiditySeconds
	s, _ := time.ParseDuration(strconv.Itoa(validitySeconds) + "s")
	expiredTime := time.Now().Add(s)
	refreshToken := &model.OAuth2Token{
		ExpiresTime:&expiredTime,
		TokenValue:uuid.NewV4().String(),
	}

	if tokenService.tokenEnhancer != nil{
		return tokenService.tokenEnhancer.Enhance(refreshToken, oauth2Details)
	}
	return refreshToken, nil
}

func (tokenService *DefaultTokenService) RefreshAccessToken(refreshTokenValue string) (*model.OAuth2Token, error){

	refreshToken, err := tokenService.tokenStore.ReadRefreshToken(refreshTokenValue)

	if err == nil{
		if refreshToken.IsExpired(){
			return nil, ErrExpiredToken
		}
		oauth2Details, err := tokenService.tokenStore.ReadOAuth2DetailsForRefreshToken(refreshTokenValue)
		if err == nil{
			oauth2Token, err := tokenService.tokenStore.GetAccessToken(oauth2Details)
			// 移除原有的访问令牌
			if err == nil{
				tokenService.tokenStore.RemoveAccessToken(oauth2Token.TokenValue)
			}

			// 移除已使用的刷新令牌
			tokenService.tokenStore.RemoveRefreshToken(refreshTokenValue)
			refreshToken, err = tokenService.createRefreshToken(oauth2Details)
			if err == nil{
				accessToken, err := tokenService.createAccessToken(refreshToken, oauth2Details)
				if err == nil{
					tokenService.tokenStore.StoreAccessToken(accessToken, oauth2Details)
					tokenService.tokenStore.StoreRefreshToken(refreshToken, oauth2Details)
				}
				return accessToken, err
			}
		}
	}
	return nil, err

}

func (tokenService *DefaultTokenService) GetAccessToken(details *model.OAuth2Details) (*model.OAuth2Token, error)  {
	return tokenService.tokenStore.GetAccessToken(details)
}

func (tokenService *DefaultTokenService) ReadAccessToken(tokenValue string) (*model.OAuth2Token, error){
	return tokenService.tokenStore.ReadAccessToken(tokenValue)
}

func (tokenService *DefaultTokenService) GetOAuth2DetailsByAccessToken(tokenValue string) (*model.OAuth2Details, error) {

	accessToken, err := tokenService.tokenStore.ReadAccessToken(tokenValue)
	if err == nil{
		if accessToken.IsExpired(){
			return nil, ErrExpiredToken
		}
		return tokenService.tokenStore.ReadOAuth2Details(tokenValue)
	}
	return nil, err
}


type TokenStore interface {
	// 存储访问令牌
	StoreAccessToken(oauth2Token *model.OAuth2Token, oauth2Details *model.OAuth2Details)
	// 根据令牌值获取访问令牌结构体
	ReadAccessToken(tokenValue string) (*model.OAuth2Token, error)
	// 根据令牌值获取令牌对应的客户端和用户信息
	ReadOAuth2Details(tokenValue string)(*model.OAuth2Details, error)
	// 根据客户端信息和用户信息获取访问令牌
	GetAccessToken(oauth2Details *model.OAuth2Details)(*model.OAuth2Token, error)
	// 移除存储的访问令牌
	RemoveAccessToken(tokenValue string)
	// 存储刷新令牌
	StoreRefreshToken(oauth2Token *model.OAuth2Token, oauth2Details *model.OAuth2Details)
	// 移除存储的刷新令牌
	RemoveRefreshToken(oauth2Token string)
	// 根据令牌值获取刷新令牌
	ReadRefreshToken(tokenValue string)(*model.OAuth2Token, error)
	// 根据令牌值获取刷新令牌对应的客户端和用户信息
	ReadOAuth2DetailsForRefreshToken(tokenValue string)(*model.OAuth2Details, error)
}

type JwtTokenStore struct {
	jwtTokenEnhancer *JwtTokenEnhancer
}
func NewJwtTokenStore(jwtTokenEnhancer *JwtTokenEnhancer) TokenStore {
	return &JwtTokenStore{
		jwtTokenEnhancer:jwtTokenEnhancer,
	}

}

func (tokenStore *JwtTokenStore)StoreAccessToken(oauth2Token *model.OAuth2Token, oauth2Details *model.OAuth2Details){

}
func (tokenStore *JwtTokenStore)ReadAccessToken(tokenValue string) (*model.OAuth2Token, error){
	oauth2Token, _, err := tokenStore.jwtTokenEnhancer.Extract(tokenValue)
	return oauth2Token, err

}

// 根据令牌值获取令牌对应的客户端和用户信息
func  (tokenStore *JwtTokenStore) ReadOAuth2Details(tokenValue string)(*model.OAuth2Details, error){
	_, oauth2Details, err := tokenStore.jwtTokenEnhancer.Extract(tokenValue)
	return oauth2Details, err

}

// 根据客户端信息和用户信息获取访问令牌
func (tokenStore *JwtTokenStore) GetAccessToken(oauth2Details *model.OAuth2Details)(*model.OAuth2Token, error){
	return nil, ErrNotSupportOperation
}
// 移除存储的访问令牌
func (tokenStore *JwtTokenStore) RemoveAccessToken(tokenValue string){

}
// 存储刷新令牌
func (tokenStore *JwtTokenStore) StoreRefreshToken(oauth2Token *model.OAuth2Token, oauth2Details *model.OAuth2Details){

}
// 移除存储的刷新令牌
func (tokenStore *JwtTokenStore)RemoveRefreshToken(oauth2Token string){

}
// 根据令牌值获取刷新令牌
func (tokenStore *JwtTokenStore) ReadRefreshToken(tokenValue string)(*model.OAuth2Token, error){
	oauth2Token, _, err := tokenStore.jwtTokenEnhancer.Extract(tokenValue)
	return oauth2Token, err
}
// 根据令牌值获取刷新令牌对应的客户端和用户信息
func (tokenStore *JwtTokenStore)ReadOAuth2DetailsForRefreshToken(tokenValue string)(*model.OAuth2Details, error){
	_, oauth2Details, err := tokenStore.jwtTokenEnhancer.Extract(tokenValue)
	return oauth2Details, err
}

type TokenEnhancer interface {
	Enhance(token *model.OAuth2Token,details *model.OAuth2Details)(*model.OAuth2Token,error)
	Extract(tokenValue string) (*model.OAuth2Token, *model.OAuth2Details, error)
}

type JwtTokenEnhancer struct {
	secretKey []byte
}

func NewJwtTokenEnhancer(secretKey string)TokenEnhancer{
	return &JwtTokenEnhancer{
		secretKey: []byte(secretKey),
	}
}

func (enhancer *JwtTokenEnhancer) Enhance(oauth2Token *model.OAuth2Token, oauth2Details *model.OAuth2Details) (*model.OAuth2Token, error) {
	return enhancer.sign(oauth2Token, oauth2Details)
}


func (enhancer *JwtTokenEnhancer) Extract(tokenValue string) (*model.OAuth2Token, *model.OAuth2Details, error)  {

	token, err := jwt.ParseWithClaims(tokenValue, &OAuth2TokenCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return enhancer.secretKey, nil
	})
	if err ==nil{
		claims := token.Claims.(*OAuth2TokenCustomClaims)
		expiresTime := time.Unix(claims.ExpiresAt,0)
		return &model.OAuth2Token{
				RefreshToken:&claims.RefreshToken,
				TokenValue:tokenValue,
				ExpiresTime: &expiresTime,
			}, &model.OAuth2Details{
				User:&claims.UserDetails,
				Client:&claims.ClientDetails,
			}, nil
	}
	return nil, nil, err
}

type OAuth2TokenCustomClaims struct {
	UserDetails model.UserDetails
	ClientDetails model.ClientDetails
	RefreshToken model.OAuth2Token
	jwt.StandardClaims
}

func(enhancer *JwtTokenEnhancer)sign(oauth2Token *model.OAuth2Token, oauth2Details *model.OAuth2Details)  (*model.OAuth2Token, error) {
	expireTime := oauth2Token.ExpiresTime
	clientDetails := *oauth2Details.Client
	userDetails := *oauth2Details.User
	clientDetails.ClientSecret = ""
	userDetails.Password = ""
	claims := OAuth2TokenCustomClaims{
		UserDetails:    userDetails,
		ClientDetails:  clientDetails,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer :"System",
		},
	}
	if oauth2Token.RefreshToken != nil{
		claims.RefreshToken = *oauth2Token.RefreshToken
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenValue, err := token.SignedString(enhancer.secretKey)
	if err==nil{
		oauth2Token.TokenValue = tokenValue
		oauth2Token.TokenType = "jwt"
		return oauth2Token,nil
	}
	return nil,err
}