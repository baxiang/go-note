package controller

import (
	"github.com/baxiang/go-note/swag-gin/model"
	"github.com/baxiang/go-note/swag-gin/utils"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"
	"net/http"
	"strconv"
)

// ShowAccount example
// @Summary Show an account
// @Description get string by ID
// @Tags ShowAccount
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} model.Account
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /accounts/{id} [get]
func (c *Controller)ShowAccount(ctx *gin.Context){
	id := ctx.Param("id")
	aid,err := strconv.Atoi(id)
	if err!=nil{
        utils.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	account, err := model.AccountOne(aid)
	if err != nil {
		utils.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, account)
}

// AddAccount godoc
// @Summary Add an account
// @Description add by json account
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param account body model.AddAccount true "Add account"
// @Success 200 {object} model.Account
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /accounts [post]
func (c *Controller) AddAccount(ctx *gin.Context) {
	var addAccount model.AddAccount
	if err := ctx.ShouldBindJSON(&addAccount); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	if err := addAccount.Validation(); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	account := model.Account{
		Name: addAccount.Name,
	}
	lastID, err := account.Insert()
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	account.ID = lastID
	ctx.JSON(http.StatusOK, account)
}