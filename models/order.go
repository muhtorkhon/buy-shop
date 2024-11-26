package models

import "time"

type Order struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID      int       `json:"user_id" gorm:"not null"`
	ProductID   int       `json:"product_id" gorm:"not null"`
	Quantity    int       `json:"quantity" gorm:"not null"`
	TotalAmount float64   `json:"total_amount" gorm:"not null"`
	OrderStatus string    `json:"order_status" gorm:"default:'pending'"`
	User        *Users    `json:"user" gorm:"foreignKey:UserID"`
	Product     *Product  `json:"product" gorm:"foreignKey:ProductID"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type OrderRequest struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
