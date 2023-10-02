package controllers

import (
	"gin-boilerplate/models"
	"gin-boilerplate/helpers"
	"gin-boilerplate/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	ginI18n "github.com/gin-contrib/i18n"

)

func GetData(ctx *gin.Context) {
	var example []*models.Example
	repository.Get(&example)
	ctx.JSON(http.StatusOK, helpers.Response{
		Data:&example,
		Message: ginI18n.MustGetMessage(ctx,"sussess"),
	    })
}
