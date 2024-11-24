package models

import (
	"github.com/golang-jwt/jwt/v5"
)

type Users struct {
	ID          int     `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName   string  `gorm:"type:varchar(255);not null" json:"first_name"`
	Email       string  `gorm:"type:varchar(255);unique;not null" json:"email"`
	PhoneNumber string  `gorm:"type:varchar(255);unique;not null" json:"phone_number"`
	Password    string  `gorm:"type:varchar(255);not null" json:"password"`
	Role        string  `gorm:"type:varchar(255);not null" json:"role"`
	IsActive    bool    `gorm:"default:false" json:"is_active"`
	Orders      []Order `gorm:"foreignKey:UserID" json:"orders"`
}

type UsersResponse struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	Role        string `json:"role"`
	IsActive    bool   `json:"is_active"`
}

// type UserRegister struct {
// 	FirstName   string `json:"first_name" binding:"required"`                    
// 	Email       string `json:"email" binding:"required" validate:"email"`        
// 	PhoneNumber string `json:"phone_number" binding:"required" validate:"phone"` 
// 	Password    string `json:"password" validate:"password"`  
// 	Role        string `json:"role" binding:"required,oneof=admin user"`        
// }

type UserRegister struct {
	FirstName   string `json:"first_name" binding:"required"`                         
	Email       string `json:"email" binding:"required,email"`                      
	PhoneNumber string `json:"phone_number" binding:"required" validate:"phone"`                    
	Password    string `json:"password" binding:"required,min=8"`                   
	Role        string `json:"role" binding:"required,oneof=admin user"`            
}



type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

type Request struct {
	PhoneNumber string `json:"phone_number"`
	Code        string `json:"code"`
}
