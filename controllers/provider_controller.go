package controllers

import (
	"errors"
	"gin-boilerplate/helpers"
	"gin-boilerplate/models"
	"gin-boilerplate/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	ginI18n "github.com/gin-contrib/i18n"
)

func GetProviders(ctx *gin.Context) {
	var example []*models.Provider
	err := repository.Get(&example)
	if err != nil {
		error := errors.New("Cannot get the list of Provider")

		ctx.JSON(http.StatusBadRequest, helpers.ResponseError{Error: error.Error()})
		return

	}
	ctx.JSON(http.StatusOK, helpers.Response{
		Data:    &example,
		Message: ginI18n.MustGetMessage(ctx, "sussess"),
	})
}

func GetProviderById(ctx *gin.Context) {
	ProviderId, err := strconv.Atoi(ctx.Param("id"))
	example := new(models.Provider)
	example.Id = ProviderId

	err = repository.GetOne(&example)
	if err != nil {
		error := errors.New("Provider is not exist")
		ctx.JSON(http.StatusBadRequest, helpers.ResponseError{Error: error.Error()})
		return

	}
	ctx.JSON(http.StatusOK, helpers.Response{
		Data:    &example,
		Message: ginI18n.MustGetMessage(ctx, "sussess"),
	})
}

func CreateProvider(ctx *gin.Context) {
	example := new(models.Provider)
	err := ctx.ShouldBindJSON(&example)
	if err != nil {
		error := errors.New("Format error")
		ctx.JSON(http.StatusBadRequest, helpers.ResponseError{Error: error.Error()})
		return
	}

	err = repository.Save(&example)
	if err != nil {
		error := errors.New("Internal Error")
		ctx.JSON(http.StatusBadRequest, helpers.ResponseError{Error: error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, helpers.Response{
		Data:    &example,
		Message: ginI18n.MustGetMessage(ctx, "sussess"),
	}) //&example)
}

func UpdateProvider(ctx *gin.Context) {
	ProviderId, err := strconv.Atoi(ctx.Param("id"))

	example := new(models.Provider)
	example.Id = ProviderId

	err = ctx.ShouldBindJSON(&example)
	if err != nil {
		error := errors.New("Format error")
		ctx.JSON(http.StatusBadRequest, helpers.ResponseError{Error: error.Error()})
		return
	}

	err = repository.Update(&example)
	if err != nil {
		error := errors.New("Internal Error")
		ctx.JSON(http.StatusBadRequest, helpers.ResponseError{Error: error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, helpers.Response{
		Data:    &example,
		Message: ginI18n.MustGetMessage(ctx, "sussess"),
	}) //&example)
}

func DeleteProvider(ctx *gin.Context) {
	ProviderId, err := strconv.Atoi(ctx.Param("id"))
	example := new(models.Provider)
	example.Id = ProviderId
	err = repository.Delete(&example)
	if err != nil {
		error := errors.New("Internal Error")
		ctx.JSON(http.StatusBadRequest, helpers.ResponseError{Error: error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, helpers.Response{
		Data:    &example,
		Message: ginI18n.MustGetMessage(ctx, "sussess"),
	}) //&example)
}
