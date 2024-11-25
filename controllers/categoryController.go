package controllers

import (
	"i-shop/models"
	"i-shop/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

type CategoryController struct {
	Storage *storage.CategoryStore
}

func NewCategoryController(storage *storage.CategoryStore) (*CategoryController, error) {
	return &CategoryController{Storage: storage}, nil
}


// @Summary 		Create a new category
// @Description 	Create a new category in the store
// @Tags 			categories
// @Accept 			json
// @Produce 		json
// @Security ApiKeyAuth
// @Param 			category body models.CategoryRequest true "Category to be created"
// @Success 		201 {object} models.CategoryRequest
// @Failure 		400 {object} Response "Invalid request format"
// @Failure 		500 {object} Response "Internal Server Error"
// @Router 			/categories [post]
func (s *CategoryController) CreateCategory(c *gin.Context) {
	var body models.CategoryRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Errorf("Failed to bind JSON for creating category: %v", err)
		HandleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	category := models.Category{
		NameUz: body.NameUz,
		NameRu: body.NameRu,
		NameEn: body.NameEn,
	}

	if err := s.Storage.Create(c.Request.Context(), &body); err != nil {
		log.Errorf("Failed to create category: %v", err)
		HandleResponse(c, http.StatusInternalServerError, "Failed to create category")
		return
	}
	HandleResponse(c, http.StatusCreated, category)
}

// @Summary 		Get all categories
// @Description 	Get all categories from the store
// @Tags 			categories
// @Accept 			json
// @Produce 		json
// @Success 		200 {array}  models.Category
// @Failure 		500 {object} Response "Internal Server Error"
// @Router 			/categories [get]
func (s *CategoryController) GetCategories(c *gin.Context) {
	categories, err := s.Storage.GetAll()
	if err != nil {
		log.Error("Failed to retrieve categories", err)
		HandleResponse(c, http.StatusInternalServerError, "Failed to retrieve categories")
		return
	}
	HandleResponse(c, http.StatusOK, categories)
}

// @Summary 		Get a category by ID
// @Description 	Get a specific category by its ID
// @Tags 			categories
// @Accept 			json
// @Produce 		json
// @Param 			id path int true "Category ID"
// @Success 		200 {object} models.Category
// @Failure 		400 {object} Response "Invalid category ID"
// @Failure 		404 {object} Response "Category not found"
// @Router 			/categories/{id} [get]
func (s *CategoryController) GetCategoryByID(c *gin.Context) {
	id := c.Param("id")

	category, err := s.Storage.GetByID(id)
	if err != nil {
		log.Warnf("Category with ID %d not found", id)
		HandleResponse(c, http.StatusNotFound, "Category not found")
		return
	}
	HandleResponse(c, http.StatusOK, category)
}

// @Summary 		Update a category by ID
// @Description 	Update an existing category by its ID
// @Tags 			categories
// @Accept 			json
// @Produce 		json
// @Param 			id path int true "Category ID"
// @Param 			category body models.CategoryRequest true "Category to be updated"
// @Success 		200 {object} models.CategoryRequest
// @Failure 		400 {object} Response "Invalid category ID or request format"
// @Failure 		500 {object} Response "Internal Server Error"
// @Router 			/categories/{id} [put]
func (s *CategoryController) UpdateCategory(c *gin.Context) {
	id := c.Param("id")

	var body models.CategoryRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Errorf("Failed to bind JSON for updating category ID %s: %v", id, err)
		HandleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := s.Storage.Update(&body, id); err != nil {
		log.Errorf("Failed to update category with ID %s: %v", id, err)
		HandleResponse(c, http.StatusInternalServerError, "Failed to update category")
		return
	}
	HandleResponse(c, http.StatusOK, body)
}

// @Summary 		Delete a category by ID
// @Description 	Delete a specific category by its ID
// @Tags 			categories
// @Accept 			json
// @Produce 		json
// @Param 			id path int true "Category ID"
// @Success 		200 {object} Response "Category deleted successfully"
// @Failure 		400 {object} Response "Invalid category ID"
// @Failure 		500 {object} Response "Internal Server Error"
// @Router 			/categories/{id} [delete]
func (s *CategoryController) DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	if err := s.Storage.SoftDelete(id); err != nil {
		log.Errorf("Failed to delete category with ID %s: %v", id, err)
		HandleResponse(c, http.StatusInternalServerError, "Failed to delete category")
		return
	}
	HandleResponse(c, http.StatusOK, "Category deleted successfully")
}

// RestoreProduct     godoc
// @Summary           Restore a soft deleted category
// @Description       Restore a previously soft deleted category by its ID.
// @Tags              categories
// @Accept            json
// @Produce           json
// @Param             id path string true "category ID" // ID of the category to restore
// @Success           200 {object} Response "Successfully restored the category"
// @Failure           500 {object} Response "Failed to restore category"
// @Router            /categories/restore/{id} [put]
func (s *CategoryController) RestoreCategory(c *gin.Context) {
	if err := s.Storage.Restore(c.Param("id")); err != nil {
		log.Printf("RestoreCategory: Failed to restore category: %v", err)
		HandleResponse(c, http.StatusInternalServerError, "Failed to restore category")
		return
	}

	HandleResponse(c, http.StatusOK, "Category restored successfully")
}
