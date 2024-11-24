package controllers

import (
	"i-shop/models"
	"i-shop/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

type BrandController struct {
	Storage *storage.BrandStore
}

func NewBrandController(storage *storage.BrandStore) (*BrandController, error) {
	log.Info("Initializing BrandController")
	return &BrandController{Storage: storage}, nil
}

// @Summary 		Create a new brand
// @Description 	Create a new brand in the store
// @Tags 			brands
// @Accept 			json
// @Produce 		json
// @Param 			data body models.BrandRequest true "Brand to be created"
// @Success 		201 {object} models.BrandRequest "Brand created successfully"
// @Failure 		400 {object} Response "Invalid request format"
// @Failure 		500 {object} Response "Internal Server Error"
// @Router 			/brands [post]
func (s *BrandController) CreateBrand(c *gin.Context) {
	var body models.BrandRequest

	log.Info("Creating a new brand")
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Errorf("Failed to parse request body: %v", err)
		HandleResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	brand := models.Brand{
		NameUz:    body.NameUz,
		NameRu:    body.NameRu,
		NameEn:    body.NameEn,
	}

	if err := s.Storage.Create(&brand); err != nil {
		log.Errorf("Failed to create brand: %v", err)
		HandleResponse(c, http.StatusInternalServerError, "Failed to create brand")
		return
	}
	HandleResponse(c, http.StatusCreated, brand)
}

// @Summary 		Get all brands
// @Description 	Get all brands from the store
// @Tags 			brands
// @Accept 			json
// @Produce 		json
// @Success 		200 {object} models.Brand "List of all brands"
// @Failure 		500 {object} Response "Internal Server Error"
// @Router 			/brands [get]
func (s *BrandController) GetBrands(c *gin.Context) {
	log.Info("Getting all brands")
	brands, err := s.Storage.GetAll()
	if err != nil {
		log.Errorf("Failed to retrieve brands: %v", err)
		HandleResponse(c, http.StatusInternalServerError, "Failed to retrieve brands")
		return
	}
	HandleResponse(c, http.StatusOK, brands)
}

// @Summary 		Get a brand by ID
// @Description 	Get a specific brand by its ID
// @Tags 			brands
// @Accept 			json
// @Produce 		json
// @Param 			id path int true "Brand ID"
// @Success 		200 {object} models.Brand "Brand details"
// @Failure 		400 {object} Response "Invalid brand ID"
// @Failure 		404 {object} Response "Brand not found"
// @Router 			/brands/{id} [get]
func (s *BrandController) GetBrandByID(c *gin.Context) {
	id := c.Param("id")
	
	log.Infof("Getting brand with ID: %s", id)
	brand, err := s.Storage.GetByID(id)
	if err != nil {
		log.Errorf("Brand with ID %s not found: %v", id, err)
		HandleResponse(c, http.StatusNotFound, "Brand not found")
		return
	}
	HandleResponse(c, http.StatusOK, brand)
}

// @Summary 		Update a brand by ID
// @Description 	Update an existing brand by its ID
// @Tags 			brands
// @Accept 			json
// @Produce 		json
// @Param 			id path int true "Brand ID"
// @Param 			brand body models.BrandRequest true "Brand to be updated"
// @Success 		200 {object} models.BrandRequest "Brand updated successfully"
// @Failure 		400 {object} Response "Invalid brand ID or request format"
// @Failure 		500 {object} Response "Internal Server Error"
// @Router 			/brands/{id} [put]
func (s *BrandController) UpdateBrand(c *gin.Context) {
	id := c.Param("id")

	var brand models.BrandRequest
	if err := c.ShouldBindJSON(&brand); err != nil {
		log.Errorf("Failed to bind JSON for brand ID %s: %v", id, err)
		HandleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := s.Storage.Update(&brand, id); err != nil {
		log.Errorf("Failed to update brand with ID %s: %v", id, err)
		HandleResponse(c, http.StatusInternalServerError, "Failed to update brand")
		return
	}
	HandleResponse(c, http.StatusOK, brand)
}

// @Summary 		Delete a brand by ID
// @Description 	Delete a specific brand by its ID
// @Tags 			brands
// @Accept 			json
// @Produce 		json
// @Param 			id path int true "Brand ID"
// @Success 		200 {object} Response "Brand deleted successfully"
// @Failure 		400 {object} Response "Invalid brand ID"
// @Failure 		500 {object} Response "Internal Server Error"
// @Router 			/brands/{id} [delete]
func (s *BrandController) DeleteBrand(c *gin.Context) {
	id := c.Param("id")

	if err := s.Storage.SoftDelete(id); err != nil {
		log.Errorf("Failed to delete brand with ID %s: %v", id, err)
		HandleResponse(c, http.StatusInternalServerError, "Failed to delete brand")
		return
	}
	HandleResponse(c, http.StatusOK, "Brand deleted successfully")
}

// RestoreProduct     godoc
// @Summary           Restore a soft deleted brand
// @Description       Restore a previously soft deleted brand by its ID.
// @Tags              brands
// @Accept            json
// @Produce           json
// @Param             id path string true "brand ID" // ID of the brand to restore
// @Success           200 {object} map[string]interface{} "Successfully restored the brand"
// @Failure           500 {object} map[string]interface{} "Failed to restore brand"
// @Router            /brands/restore/{id} [put]
func (s *BrandController) RestoreBrand(c *gin.Context) {
	id := c.Param("id")

	if err := s.Storage.Restore(id); err != nil {
		log.Printf("RestoreBrand: Failed to restore brand with ID %s: %v", id, err)
		HandleResponse(c, http.StatusInternalServerError, "Failed to restore brand")
		return
	}
	HandleResponse(c, http.StatusOK, "Brand restored successfully")
}
