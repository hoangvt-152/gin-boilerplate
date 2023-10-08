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

func GetProviderProductions(ctx *gin.Context) {
	var example []*models.ProviderProduction
	err := repository.Get(&example)
	if err != nil {
		error := errors.New("Cannot get the list of ProviderProduction")

		ctx.JSON(http.StatusBadRequest, helpers.ResponseError{Error: error.Error()})
		return

	}
	ctx.JSON(http.StatusOK, helpers.Response{
		Data:    &example,
		Message: ginI18n.MustGetMessage(ctx, "sussess"),
	})
}

func GetProviderProductionById(ctx *gin.Context) {
	ProviderProductionId, err := strconv.Atoi(ctx.Param("id"))
	example := new(models.ProviderProduction)
	example.Id = ProviderProductionId

	err = repository.GetOne(&example)
	if err != nil {
		error := errors.New("ProviderProduction is not exist")
		ctx.JSON(http.StatusBadRequest, helpers.ResponseError{Error: error.Error()})
		return

	}
	ctx.JSON(http.StatusOK, helpers.Response{
		Data:    &example,
		Message: ginI18n.MustGetMessage(ctx, "sussess"),
	})
}

func CreateProviderProduction(ctx *gin.Context) {
	example := new(models.ProviderProduction)
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

func UpdateProviderProduction(ctx *gin.Context) {
	ProviderProductionId, err := strconv.Atoi(ctx.Param("id"))

	example := new(models.ProviderProduction)
	example.Id = ProviderProductionId

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

func DeleteProviderProduction(ctx *gin.Context) {
	ProviderProductionId, err := strconv.Atoi(ctx.Param("id"))
	example := new(models.ProviderProduction)
	example.Id = ProviderProductionId
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
