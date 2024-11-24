package storage

import (
	"fmt"
	"i-shop/models"
	"log"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductStore struct {
	db *gorm.DB
}

func NewProductStorage(db *gorm.DB) *ProductStore {
	return &ProductStore{db: db}
}

func (s *ProductStore) Create(product models.Product) error {
	result := s.db.Create(&product)
	if result.Error != nil {
		log.Printf("Create product error: %v", result.Error)
		return result.Error
	}
	return nil
}

func (s *ProductStore) GetCategory(filtr *models.ProductFilter) ([]models.Product, int64, error) {
	var products []models.Product
	var count int64

	db := s.db.Model(&models.Product{})

	if filtr.CategoryID != 0 {
		db = db.Where("category_id = ?", filtr.CategoryID).Preload("Category")
	}

	if filtr.BrandID != 0 {
		db = db.Where("brand_id = ?", filtr.BrandID).Preload("Brand")
	}

	db = db.Preload("Image")

	if err := db.Count(&count).Error; err != nil {
		log.Printf("Error counting products by filters (CategoryID: %d, BrandID: %d): %v", filtr.CategoryID, filtr.BrandID, err)
		return nil, 0, err
	}

	if filtr.Page < 1 {
		filtr.Page = 1
	}
	if filtr.PageSize < 1 {
		filtr.PageSize = 10
	}
	if filtr.PageSize > 100 {
		filtr.PageSize = 100
	}

	limit := filtr.PageSize
	offset := (filtr.Page - 1) * filtr.PageSize

	err := db.Limit(limit).
		Offset(offset).
		Find(&products).Error

	if err != nil {
		log.Printf("Error retrieving products by filters (CategoryID: %d, BrandID: %d): %v", filtr.CategoryID, filtr.BrandID, err)
		return nil, 0, err
	}

	return products, count, nil
}

func (s *ProductStore) GetByID(id string) (*models.Product, error) {
	var product models.Product
	err := s.db.Where("deleted_at IS NULL").Debug().First(&product, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("GetByID: Brand with ID %s not found", id)
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}

func (s *ProductStore) Update(newProduct *models.ProductRequest, id string) error {
	var product models.Product

	updateFields := map[string]interface{}{
		"name_uz":        newProduct.NameUz,
		"name_ru":        newProduct.NameRu,
		"name_en":        newProduct.NameEn,
		"price":          newProduct.Price,
		"description_uz": newProduct.DescriptionUz,
		"description_ru": newProduct.DescriptionRu,
		"description_en": newProduct.DescriptionEn,
		"category_id":    newProduct.CategoryID,
		"brand_id":       newProduct.BrandID,
	}

	rows := s.db.Clauses(clause.Returning{}).Model(&product).Where("id = ?", id).Updates(&updateFields)
	if rows.RowsAffected == 0 {
		return fmt.Errorf("not found record")
	} else if rows.Error != nil {
		return rows.Error
	}

	return nil
}

func (s *ProductStore) SoftDelete(id string) error {
	var product models.Product

	if err := s.db.Unscoped().Where("id = ?", id).First(&product).Error; err != nil {
		log.Printf("SoftDelete: Failed to find product with ID %s: %v", id, err)
		return err
	}

	if product.DeletedAt != nil {
		log.Printf("SoftDelete: Product with ID %s is already soft deleted", id)
		return fmt.Errorf("product with ID %s is already soft deleted", id)
	}

	now := time.Now()
	product.DeletedAt = &now

	if err := s.db.Save(&product).Error; err != nil {
		log.Printf("SoftDelete: Failed to soft delete product with ID %s: %v", id, err)
		return err
	}
	return nil
}

func (s *ProductStore) Restore(id string) error {

	var product models.Product
	if err := s.db.Unscoped().Where("id = ?", id).First(&product).Error; err != nil {
		log.Printf("Restore: Failed to find soft deleted product with ID %s: %v", id, err)
		return err
	}

	product.DeletedAt = nil

	if err := s.db.Save(&product).Error; err != nil {
		log.Printf("Restore: Failed to restore product with ID %s: %v", id, err)
		return err
	}
	return nil
}
