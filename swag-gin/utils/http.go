package utils

import "github.com/gin-gonic/gin"

type HTTPError struct {
	Code int `json:"code"`
	Message string `json:"message"`
} 

func NewError(ctx *gin.Context,status int ,err error){
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}