package controllers

import (
	"gin-boilerplate/models"
	"gin-boilerplate/helpers"
	"gin-boilerplate/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	ginI18n "github.com/gin-contrib/i18n"

)

func GetProviders(ctx *gin.Context) {
	var example []*models.Provider
	repository.Get(&example)
	ctx.JSON(http.StatusOK, helpers.Response{
		Data:&example,
		Message: ginI18n.MustGetMessage(ctx,"sussess"),
	    })
}

func CreateProviders(ctx *gin.Context) {
	example := new(models.Provider)
	repository.Save(&example)
	ctx.JSON(http.StatusOK, &example)
  }
