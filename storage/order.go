package storage

import (
	"i-shop/models"

	"gorm.io/gorm"
)


type OrderStore struct {
	db *gorm.DB
}

func NewOrderStorage(db *gorm.DB) *OrderStore {
	return &OrderStore{db: db}
}

func (o *OrderStore) GetUserIDByEmail(email string) (int, error) {
	var user models.Users
	if err := o.db.Where("email = ?", email).First(&user).Error; err != nil {
		return 0, err 
	}
	return user.ID, nil
}

func (o *OrderStore) CheckProductAvailability(productID int, quantity int) bool {
	var product models.Product
	
	if err := o.db.First(&product, productID).Error; err != nil {
		return false
	}

	return product.Stock >= quantity
}

func (o *OrderStore) CalculateTotalAmount(productID int, quantity int) (float64, error) {
	var product models.Product

	if err := o.db.First(&product, productID).Error; err != nil {
		return 0, err
	}

	return product.Price * float64(quantity), nil
}

func (o *OrderStore) SaveOrder(order models.Order) error {
	err := o.db.Create(&order).Error
	if err != nil {
		return err
	}
	return nil
}


