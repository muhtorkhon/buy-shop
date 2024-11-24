package controllers

import (
	"i-shop/models"
	"i-shop/storage"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	Storage *storage.ProductStore
}

func NewProductController(storage *storage.ProductStore) (*ProductController, error) {
	return &ProductController{Storage: storage}, nil
}

// GetProductsByFilters godoc
// @Summary 		Get products by filters
// @Description 	Get products filtered by category, brand, page, and page size
// @Tags 			products
// @Accept  		json
// @Produce  		json
// @Param 			brand_id query int false "Brand ID"
// @Param 			category_id query int false "Category ID"
// @Param 			page query int false "Page number"
// @Param 			page_size query int false "Number of items per page"
// @Success 		200 {object} Response "Successfully retrieved products"
// @Failure 		400 {object} Response "Invalid input parameters"
// @Failure 		404 {object} Response "Products not found"
// @Router 			/products [get]
func (s *ProductController) GetProductsByFilters(c *gin.Context) {
	var filtr models.ProductFilter

	filtr.BrandID, _ = strconv.Atoi(c.DefaultQuery("brand_id", "0"))
	filtr.CategoryID, _ = strconv.Atoi(c.DefaultQuery("category_id", "0"))
	filtr.Page, _ = strconv.Atoi(c.DefaultQuery("page", "0"))
	filtr.PageSize, _ = strconv.Atoi(c.DefaultQuery("page_size", "10"))

	products, count, err := s.Storage.GetCategory(&filtr)
	if err != nil {
		log.Println("Error retrieving products:", err.Error())
		HandleResponse(c, http.StatusNotFound, "Products not found")
		return
	}
	result := struct {
		Products []models.Product `json:"products"`
		Count    int64            `json:"total_count"`
		Page     int              `json:"page"`
		PageSize int              `json:"page_size"`
	}{
		Products: products,
		Count:    count,
		Page:     filtr.Page,
		PageSize: filtr.PageSize,
	}

	HandleResponse(c, http.StatusOK, result)
}

// GetByID godoc
// @Summary      Get product by ID
// @Description  Retrieves a product by its unique ID.
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path  int  true  "Product ID"
// @Success      200  {object}  Response "Successfully retrieved product"
// @Failure      404  {object}  Response  "Product not found"
// @Failure      500  {object}  Response  "Internal server error"
// @Router       /products/{id} [get]
func (s *ProductController) GetByID(c *gin.Context) {
	id := c.Param("id")

	result, err := s.Storage.GetByID(id)
	if err != nil {
		log.Printf("Product with ID %s not found: %v", id, err)
		HandleResponse(c, http.StatusNotFound, "Product not found")
		return
	}
	log.Printf("Successfully retrieved product with ID: %s", id)
	HandleResponse(c, http.StatusOK, result)
}

// @Summary 		Create a new product
// @Description 	Create a new product in the store
// @Tags 			products
// @Accept 			json
// @Produce 		json
// @Param 			product body models.ProductRequest true "Product to be created"
// @Success 		201 {object} Response "Successfully created product"
// @Failure 		400 {object} Response "Invalid request format"
// @Failure 		500 {object} Response "Failed to create product"
// @Router 			/products [post]
func (s *ProductController) CreateProduct(c *gin.Context) {
	var body models.ProductRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		log.Println("Invalid product creation data:", err.Error())
		HandleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	product := models.Product{
		NameUz:        body.NameUz,
		NameRu:        body.NameRu,
		NameEn:        body.NameEn,
		DescriptionUz: body.DescriptionUz,
		DescriptionRu: body.DescriptionRu,
		DescriptionEn: body.DescriptionEn,
		BrandID:       body.BrandID,
		CategoryID:    body.CategoryID,
		Image:         body.Image,
	}

	if err := s.Storage.Create(product); err != nil {
		log.Println("Failed to create product:", err.Error())
		HandleResponse(c, http.StatusInternalServerError, "Failed to create product")
		return
	}

	log.Println("Successfully created new product")
	HandleResponse(c, http.StatusCreated, product)
}

// @Summary 		Update a product by ID
// @Description 	Update an existing product by its ID
// @Tags 			products
// @Accept 			json
// @Produce 		json
// @Param 			id path int true "Product ID"
// @Param 			product body models.ProductRequest true "Product data to update"
// @Success 		200 {object} Response "Successfully updated product"
// @Failure 		400 {object} Response "Invalid product ID or request format"
// @Failure 		500 {object} Response "Failed to update product"
// @Router 			/products/{id} [put]
func (s *ProductController) UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	var product models.ProductRequest
	if err := c.ShouldBindJSON(&product); err != nil {
		log.Println("Invalid product update data:", err.Error())
		HandleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := s.Storage.Update(&product, id); err != nil {
		log.Println("Failed to update product:", err.Error())
		HandleResponse(c, http.StatusInternalServerError, "Failed to update product")
		return
	}

	log.Println("Successfully updated product with ID:", id)
	HandleResponse(c, http.StatusOK, product)
}

// @Summary 		Delete a product by ID
// @Description 	Delete a specific product by its ID
// @Tags 			products
// @Accept 			json
// @Produce 		json
// @Param   		id path int true "Product ID"
// @Success 		200 {object} Response "Product deleted successfully"
// @Failure 		400 {object} Response "Invalid product ID"
// @Failure 		500 {object} Response "Failed to delete product"
// @Router 			/products/{id} [delete]
func (s *ProductController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	if err := s.Storage.SoftDelete(id); err != nil {
		log.Println("Failed to delete product:", err.Error())
		HandleResponse(c, http.StatusInternalServerError, "Failed to delete product")
		return
	}

	log.Println("Successfully deleted product with ID:", id)
	HandleResponse(c, http.StatusOK, "Product marked as deleted")
}

// RestoreProduct     godoc
// @Summary           Restore a soft deleted product
// @Description       Restore a previously soft deleted product by its ID.
// @Tags              products
// @Accept            json
// @Produce           json
// @Param             id path string true "Product ID" // ID of the product to restore
// @Success           200 {object} map[string]interface{} "Successfully restored the product"
// @Failure           500 {object} map[string]interface{} "Failed to restore product"
// @Router            /products/restore/{id} [put]
func (s *ProductController) RestoreProduct(c *gin.Context) {
	id := c.Param("id")

	if err := s.Storage.Restore(id); err != nil {
		log.Printf("RestoreProduct: Failed to restore product with ID %s: %v", id, err)
		HandleResponse(c, http.StatusInternalServerError, "Failed to restore product")
		return
	}

	log.Printf("RestoreProduct: Product with ID %s restored successfully", id)
	HandleResponse(c, http.StatusOK, "Product restored successfully")
}
