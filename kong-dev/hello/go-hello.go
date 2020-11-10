package hello

import (
	"github.com/Kong/go-pdk"
)

type Config struct {
	Message string
}

func New() interface{} {
	return &Config{Message: "kong"}
}

func (conf Config) Access(kong *pdk.PDK) {
	message := conf.Message
	kong.Response.SetHeader("x-kong-request", message)
}
