package storage

import (
	"i-shop/models"
	"log"
	"time"
	"fmt"

	"gorm.io/gorm"
)

type BrandStore struct {
	db *gorm.DB
}

func NewBrandStorage(db *gorm.DB) *BrandStore {
	return &BrandStore{db: db}
}

func (s *BrandStore) Create(brand *models.Brand) error {
	if err := s.db.Create(&brand).Error; err != nil {
		log.Printf("Create: Error creating brand: %v", err)
		return err
	}
	return nil
}

func (s *BrandStore) GetAll() ([]models.Brand, error) {
	var brands []models.Brand

	err := s.db.Where("deleted_at IS NULL").Find(&brands).Error
	if err != nil {
		log.Printf("GetAll: Error getting brand names: %v", err)
		return nil, err
	}
	return brands, nil
}

func (s *BrandStore) GetByID(id string) (*models.Brand, error) {
	var brand models.Brand

	err := s.db.Where("deleted_at IS NULL").First(&brand, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("GetByID: Brand with ID %s not found", id)
		} else {
			log.Printf("GetByID: Error getting brand by ID %s: %v", id, err)
		}
		return nil, err
	}
	return &brand, nil
}

func (s *BrandStore) Update(newbrand *models.BrandRequest, id string) error {
	var brand models.Brand

	if err := s.db.First(&brand, id).Error; err != nil {
		log.Printf("Update: Brand with ID %s not found: %v", id, err)
		return err
	}

	updateFields := map[string]interface{}{
		"name_uz": newbrand.NameUz,
		"name_ru": newbrand.NameRu,
		"name_en": newbrand.NameEn,
	}

	if err := s.db.Model(&brand).Updates(&updateFields).Error; err != nil {
		log.Printf("Update: Failed to update brand with ID %s: %v", id, err)
		return err
	}
	return nil
}

func (s *BrandStore) SoftDelete(id string) error {
	var brand models.Brand
	
	if err := s.db.Unscoped().Where("id = ?", id).First(&brand).Error; err != nil {
		log.Printf("SoftDelete: Failed to find brand with ID %s: %v", id, err)
		return err
	}
	
	if brand.DeletedAt != nil {
		log.Printf("SoftDelete: Brand with ID %s is already soft deleted", id)
		return fmt.Errorf("brand with ID %s is already soft deleted", id)
	}

	now := time.Now()
	brand.DeletedAt = &now
	
	if err := s.db.Save(&brand).Error; err != nil {
		log.Printf("SoftDelete: Failed to soft delete brand with ID %s: %v", id, err)
		return err
	}

	log.Printf("SoftDelete: Brand with ID %s soft deleted", id)
	return nil
}

func (s *BrandStore) Restore(id string) error {

	var brand models.Brand
	if err := s.db.Where("id = ?", id).First(&brand).Error; err != nil {
		log.Printf("Restore: Failed to find soft deleted brand with ID %s: %v", id, err)
		return err
	}

	brand.DeletedAt = nil
	
	if err := s.db.Save(&brand).Error; err != nil {
		log.Printf("Restore: Failed to restore brand with ID %s: %v", id, err)
		return err
	}
	return nil
}



