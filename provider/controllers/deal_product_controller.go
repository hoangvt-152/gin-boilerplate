package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"gin-boilerplate/config"
	"gin-boilerplate/helpers"
	"gin-boilerplate/infra/logger"
	"gin-boilerplate/models"
	"gin-boilerplate/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	ginI18n "github.com/gin-contrib/i18n"
)

func GetProductDeals(ctx *gin.Context) {
	var example []*models.ProductDeal
	err := repository.Get(&example)
	if err != nil {
		error := errors.New("Cannot get the list of ProductDeal")

		ctx.JSON(http.StatusBadRequest, helpers.ResponseError{Error: error.Error()})
		return

	}
	ctx.JSON(http.StatusOK, helpers.Response{
		Data:    &example,
		Message: ginI18n.MustGetMessage(ctx, "sussess"),
	})
}

func GetProductDealById(ctx *gin.Context) {
	ProductDealId, err := strconv.Atoi(ctx.Param("id"))
	example := new(models.ProductDeal)
	example.Id = ProductDealId

	err = repository.GetOne(&example)
	if err != nil {
		error := errors.New("ProductDeal is not exist")
		ctx.JSON(http.StatusBadRequest, helpers.ResponseError{Error: error.Error()})
		return

	}
	ctx.JSON(http.StatusOK, helpers.Response{
		Data:    &example,
		Message: ginI18n.MustGetMessage(ctx, "sussess"),
	})
}

func CreateProductDeal(ctx *gin.Context) {
	example := new(models.ProductDeal)
	err := ctx.ShouldBindJSON(&example)
	if err != nil {
		error := errors.New("Format error")
		ctx.JSON(http.StatusBadRequest, helpers.ResponseError{Error: error.Error()})
		return
	}

	err = repository.Save(&example)
	qrCodeData := map[string]interface{}{
		"id":              example.Id,
		"provider_id":     example.ProductId,
		"state":           example.State,
		"discount number": example.DiscountNumber,
		"discount type":   example.DiscountType,
		"expired":         example.ExpiredTime,
	}
	fmt.Println(config.QrCodeImageFolder)
	qrCodeImagePathFile := fmt.Sprintf("%s/qr_code_%d.png", config.QrCodeImageFolder, example.Id)
	jsonQrData, jerr := json.Marshal(qrCodeData)
	helpers.QrCodeImageGenerator(string(jsonQrData), qrCodeImagePathFile)
	if jerr != nil {
		logger.Debugf("Error happends during generate json")
	}

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

func UpdateProductDeal(ctx *gin.Context) {
	ProductDealId, err := strconv.Atoi(ctx.Param("id"))

	example := new(models.ProductDeal)
	example.Id = ProductDealId

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

func DeleteProductDeal(ctx *gin.Context) {
	ProductDealId, err := strconv.Atoi(ctx.Param("id"))
	example := new(models.ProductDeal)
	example.Id = ProductDealId
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
