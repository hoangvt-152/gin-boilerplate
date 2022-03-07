package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pkg/src/models"
)

// GetNormalData get normal data if added pagination see example_controller
func (base *Controller) GetNormalData(ctx *gin.Context) {
	var categories []models.Category
	base.DB.Find(&categories)
	ctx.JSON(http.StatusOK, gin.H{"data": categories})

}

// GetForeignRelationData Get Foreign Data with Preload
func (base *Controller) GetForeignRelationData(ctx *gin.Context) {
	var articles []models.Article
	base.DB.Preload("Category").Find(&articles)
	ctx.JSON(http.StatusOK, &articles)

}

// SelectedFiledFetch fields fetch from defining new struct
type SelectedFiledFetch struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

func (base *Controller) GetSelectedFieldData(ctx *gin.Context) {
	var selectData []SelectedFiledFetch
	base.DB.Model(&models.Article{}).Find(&selectData)
	ctx.JSON(http.StatusOK, selectData)

}
