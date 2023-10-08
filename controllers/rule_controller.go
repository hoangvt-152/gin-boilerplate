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

func GetProviderRules(ctx *gin.Context) {
	var example []*models.ProviderRule
	err := repository.Get(&example)
	if err != nil {
		error := errors.New("Cannot get the list of ProviderRule")

		ctx.JSON(http.StatusBadRequest, helpers.ResponseError{Error: error.Error()})
		return

	}
	ctx.JSON(http.StatusOK, helpers.Response{
		Data:    &example,
		Message: ginI18n.MustGetMessage(ctx, "sussess"),
	})
}

func GetProviderRuleById(ctx *gin.Context) {
	ProviderRuleId, err := strconv.Atoi(ctx.Param("id"))
	example := new(models.ProviderRule)
	example.Id = ProviderRuleId

	err = repository.GetOne(&example)
	if err != nil {
		error := errors.New("ProviderRule is not exist")
		ctx.JSON(http.StatusBadRequest, helpers.ResponseError{Error: error.Error()})
		return

	}
	ctx.JSON(http.StatusOK, helpers.Response{
		Data:    &example,
		Message: ginI18n.MustGetMessage(ctx, "sussess"),
	})
}

func CreateProviderRule(ctx *gin.Context) {
	example := new(models.ProviderRule)
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

func UpdateProviderRule(ctx *gin.Context) {
	ProviderRuleId, err := strconv.Atoi(ctx.Param("id"))

	example := new(models.ProviderRule)
	example.Id = ProviderRuleId

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

func DeleteProviderRule(ctx *gin.Context) {
	ProviderRuleId, err := strconv.Atoi(ctx.Param("id"))
	example := new(models.ProviderRule)
	example.Id = ProviderRuleId
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
