package controllers

import (
	"errors"
	"exchangeapp/global"
	"exchangeapp/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateArticle(ctx *gin.Context) {
	var article models.Article
	// Bind JSON to article
	if err := ctx.ShouldBindJSON(&article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create article database
	if err := global.Db.AutoMigrate(&article); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// add article to database
	if err := global.Db.Create(&article).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, article)
}

func GetArticles(ctx *gin.Context) {
	var articles []models.Article
	// Get all articles from database
	if err := global.Db.Find(&articles).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return articles
	ctx.JSON(http.StatusOK, articles)
}

func GetArticlesByID(ctx *gin.Context) {
	// Get article ID from URL
	id := ctx.Param("id")
	// Get article from database
	var article models.Article
	if err := global.Db.Where("id = ?", id).First(&article).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	// Return article
	ctx.JSON(http.StatusOK, article)
}
