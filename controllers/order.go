package controllers

import (
	"i-shop/models"
	"i-shop/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)


type OrderController struct {
	Storage *storage.OrderStore
}

func NewOrderController(storage *storage.OrderStore) *OrderController {
	log.Info("Initializing OrderController")
	return &OrderController{Storage: storage}

}

// CreateOrder 	 godoc
// @Summary      Create a new order
// @Description  Allows the authenticated user to create a new order by providing product ID, quantity, and other details. 
//               It validates the input, checks product availability, and calculates the total amount before saving the order.
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        order body models.Order true "Order data"
// @Success      201 {object} models.Order "Successfully created the order"
// @Failure      400 {object} map[string]interface{} "Invalid input or product not available"
// @Failure      401 {object} map[string]interface{} "Unauthorized"
// @Failure      500 {object} map[string]interface{} "Internal server error"
// @Router       /orders [post]
func (o *OrderController) CreateOrder(c *gin.Context) {
	var order models.Order

	email, exists := c.Get("email")
	if !exists {
		HandleResponse(c, http.StatusUnauthorized, "Unauthorized")
		log.Warn("Error: Failed to fetch user email from context")
		return
	}

	userID, err := o.Storage.GetUserIDByEmail(email.(string))
	if err != nil {
		HandleResponse(c, http.StatusInternalServerError, "Failed to fetch user ID")
		log.Errorf("Error: Unable to fetch user ID for email %s: %v", email.(string), err)
		return
	}
	
	if err := c.ShouldBindJSON(&order); err != nil {
		HandleResponse(c, http.StatusBadRequest, "Invalid input data")
		log.Warnf("Error: Invalid order data received: %v", err)
		return
	}

	order.UserID = userID

	if !o.Storage.CheckProductAvailability(order.ProductID, order.Quantity) {
		HandleResponse(c, http.StatusBadRequest, "Product not available")
		log.Warnf("Error: Product ID %d is not available in sufficient quantity", order.ProductID)
		return
	}

	summa, err := o.Storage.CalculateTotalAmount(order.ProductID, order.Quantity)
	if err != nil {
		HandleResponse(c, http.StatusBadRequest, "Failed to calculate total amount")
    	log.Errorf("Error calculating total amount for Product ID %d: %v", order.ProductID, err)
    	return
	}

	order.TotalAmount = summa
	order.OrderStatus = "new"

	if err := o.Storage.SaveOrder(order); err != nil {
		HandleResponse(c, http.StatusInternalServerError, "Failed to create order")
		log.Errorf("Error: Failed to save order for user ID %d: %v", userID, err)
		return
	}

	HandleResponse(c, http.StatusCreated, order)
	log.Printf("Success: Order created for user ID %d with order details: %+v", userID, order)

}


