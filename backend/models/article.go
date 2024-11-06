package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title   string `binding:"required" json:"title"`
	Content string `binding:"required" json:"content"`
	Preview string `binding:"required" json:"preview"`
}
