package models

import (
	"time"
)

type Category struct {
	ID        int        `gorm:"primaryKey;autoIncrement" json:"id"`
	NameUz    string     `gorm:"type:varchar(255);not null" json:"name_uz"`
	NameRu    string     `gorm:"type:varchar(255);not null" json:"name_ru"`
	NameEn    string     `gorm:"type:varchar(255);not null" json:"name_en"`
	DeletedAt *time.Time `gorm:"index" json:"-"`
	CreatedAt *time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type CategoryRequest struct {
	NameUz string `gorm:"type:varchar(255);not null" json:"name_uz"`
	NameRu string `gorm:"type:varchar(255);not null" json:"name_ru"`
	NameEn string `gorm:"type:varchar(255);not null" json:"name_en"`
}
