package storage

import (
	"context"
	"fmt"
	"i-shop/models"
	"log"
	"time"

	"gorm.io/gorm"
)

type CategoryStore struct {
	db *gorm.DB
}

func NewCategoryStorage(db *gorm.DB) *CategoryStore {
	return &CategoryStore{db: db}
}

func (s *CategoryStore) Create(ctx context.Context, category *models.CategoryRequest) error {
	if err := s.db.WithContext(ctx).Create(&category).Model(&models.Category{}).Error; err != nil {
		log.Printf("Create: Failed to create category: %v", err)
		return err
	}
	return nil
}

func (s *CategoryStore) GetAll() ([]*models.Category, error) {
	var categories []*models.Category
	err := s.db.Find(&categories).Error
	if err != nil {
		log.Printf("GetAll: Error getting category names: %v", err)
		return nil, err
	}
	return categories, nil
}

func (s *CategoryStore) GetByID(id string) (*models.Category, error) {
	var category models.Category
	err := s.db.First(&category, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("GetByID: Category with ID %s not found", id)
		} else {
			log.Printf("GetByID: Error getting category by ID %s: %v", id, err)
		}
		return nil, err
	}
	return &category, nil
}

func (s *CategoryStore) Update(newcategory *models.CategoryRequest, id string) error {
	var category models.Category
	if err := s.db.First(&category, id).Error; err != nil {
		log.Printf("Update: Category with ID %s not found: %v", id, err)
		return err
	}

	updateFields := map[string]interface{}{
		"name_uz": newcategory.NameUz,
		"name_ru": newcategory.NameRu,
		"name_en": newcategory.NameEn,
	}

	if err := s.db.Model(&category).Updates(&updateFields).Error; err != nil {
		log.Printf("Update: Failed to update category with ID %s: %v", id, err)
		return err
	}

	log.Printf("Update: Successfully updated category with ID %s", id)
	return nil
}

func (s *CategoryStore) SoftDelete(id string) error {
	var category models.Category

	if err := s.db.Unscoped().Where("id = ?", id).First(&category).Error; err != nil {
		log.Printf("SoftDelete: Failed to find category with ID %s: %v", id, err)
		return err
	}

	if category.DeletedAt != nil {
		log.Printf("SoftDelete: Category with ID %s is already soft deleted", id)
		return fmt.Errorf("category with ID %s is already soft deleted", id)
	}

	now := time.Now()
	category.DeletedAt = &now

	if err := s.db.Save(&category).Error; err != nil {
		log.Printf("SoftDelete: Failed to soft delete category with ID %s: %v", id, err)
		return err
	}
	return nil
}

func (s *CategoryStore) Restore(id string) error {

	var category models.Category
	if err := s.db.Unscoped().Where("id = ?", id).First(&category).Error; err != nil {
		log.Printf("Restore: Failed to find soft deleted category with ID %s: %v", id, err)
		return err
	}

	category.DeletedAt = nil

	if err := s.db.Save(&category).Error; err != nil {
		log.Printf("Restore: Failed to restore category with ID %s: %v", id, err)
		return err
	}
	return nil
}
