package controllers

import (
	"errors"
	"gin-boilerplate/helpers"
	"gin-boilerplate/models"
	"gin-boilerplate/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	ginI18n "github.com/gin-contrib/i18n"
)

func GetCategories(ctx *gin.Context) {
	var example []*models.CategoryOutput
	err := repository.Get(&example)
	if err != nil {
		error := errors.New("Cannot get the list of category")

		ctx.JSON(http.StatusBadRequest, helpers.ResponseError{Error: error.Error()})
		return

	}
	ctx.JSON(http.StatusOK, helpers.Response{
		Data:    &example,
		Message: ginI18n.MustGetMessage(ctx, "sussess"),
	})
}

func GetCategoryById(ctx *gin.Context) {
	categoryId, err := strconv.Atoi(ctx.Param("id"))
	example := new(models.CategoryOutput)
	example.Id = categoryId

	err = repository.GetOne(&example)
	if err != nil {
		error := errors.New("Category is not exist")
		ctx.JSON(http.StatusBadRequest, helpers.ResponseError{Error: error.Error()})
		return

	}
	ctx.JSON(http.StatusOK, helpers.Response{
		Data:    &example,
		Message: ginI18n.MustGetMessage(ctx, "sussess"),
	})
}

func CreateCategory(ctx *gin.Context) {
	example := new(models.Category)
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

func UpdateCategory(ctx *gin.Context) {
	categoryId, err := strconv.Atoi(ctx.Param("id"))

	example := new(models.Category)
	example.Id = categoryId

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

func DeleteCategory(ctx *gin.Context) {
	categoryId, err := strconv.Atoi(ctx.Param("id"))
	example := new(models.Category)
	example.Id = categoryId
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
