package models

import (
	"time"
)

type Image struct {
	ID        int      `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID int      `gorm:"not null" json:"product_id"`
	URL       string   `gorm:"type:text;not null" json:"url"`
	Alt       string   `gorm:"type:varchar(255)" json:"alt"`
	Width     int      `json:"width"`
	Height    int      `json:"height"`
	Product   *Product `gorm:"foreignKey:ProductID" json:"product"`
}

type Product struct {
	ID            int        `gorm:"primaryKey;autoIncrement" json:"id"`
	NameUz        string     `gorm:"type:varchar(255);not null" json:"name_uz"`
	NameRu        string     `gorm:"type:varchar(255);not null" json:"name_ru"`
	NameEn        string     `gorm:"type:varchar(255);not null" json:"name_en"`
	Price         float64    `gorm:"not null" json:"price" binding:"required"`
	DescriptionUz string     `gorm:"type:text" json:"description_uz"`
	DescriptionRu string     `gorm:"type:text" json:"description_ru"`
	DescriptionEn string     `gorm:"type:text" json:"description_en"`
	Stock         int        `gorm:"default:0" json:"stock"`
	CategoryID    int        `gorm:"not null" json:"category_id" binding:"required"`
	BrandID       int        `gorm:"not null" json:"brand_id" binding:"required"`
	Image         []Image    `gorm:"foreignKey:ProductID" json:"images"`
	DeletedAt     *time.Time `gorm:"index" json:"-"`
	CreatedAt     *time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Brand         *Brand     `gorm:"foreignKey:BrandID" json:"brand"`
	Category      *Category  `gorm:"foreignKey:CategoryID" json:"category"`
}

type ProductFilter struct {
	BrandID    int `form:"brand_id"`
	CategoryID int `form:"category_id"`
	Page       int `form:"page"`
	PageSize   int `form:"page_size"`
}

type ProductRequest struct {
	NameUz        string  `json:"name_uz"`
	NameRu        string  `json:"name_ru"`
	NameEn        string  `json:"name_en"`
	Price         float64 `json:"price" binding:"required"`
	DescriptionUz string  `json:"description_uz"`
	DescriptionRu string  `json:"description_ru"`
	DescriptionEn string  `json:"description_en"`
	CategoryID    int     `json:"category_id" binding:"required"`
	BrandID       int     `json:"brand_id" binding:"required"`
	Image         []Image `json:"images"`
}


