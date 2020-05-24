package model

type ClientDetails struct {
	//client 的标识
	ClientId string

	//client 的密钥
	ClientSecret string

	//访问令牌有效时间，秒
	AccessTokenValiditySeconds int

	RefreshTokenValiditySeconds int

	RegisteredRedirectUri string

	// 可以使用的授权类型
	AuthorizedGrantTypes []string
}