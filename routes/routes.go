package routes

import (
	"fmt"
	"i-shop/controllers"
	"i-shop/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(ctrl *controllers.Controller) *gin.Engine {
	router := gin.Default()

	router.GET("/products", ctrl.Product.GetProductsByFilters)       
	router.GET("/products/:id", ctrl.Product.GetByID)                
	router.POST("/products", ctrl.Product.CreateProduct)             
	router.PUT("/products/:id", ctrl.Product.UpdateProduct)          
	router.DELETE("/products/:id", ctrl.Product.DeleteProduct)       
	router.PUT("/products/restore/:id", ctrl.Product.RestoreProduct)

	router.GET("/categories", ctrl.Category.GetCategories)          
	router.GET("/categories/:id", ctrl.Category.GetCategoryByID) 
	router.POST("/categories", ctrl.Category.CreateCategory)       
	router.PUT("/categories/:id", ctrl.Category.UpdateCategory)   
	router.DELETE("/categories/:id", ctrl.Category.DeleteCategory)
	router.PUT("/categories/restore/:id", ctrl.Category.RestoreCategory)

	router.GET("/brands", ctrl.Brand.GetBrands)                       
	router.GET("/brands/:id", ctrl.Brand.GetBrandByID)     
	router.POST("/brands", ctrl.Brand.CreateBrand)              
	router.PUT("/brands/:id", ctrl.Brand.UpdateBrand)      
	router.DELETE("/brands/:id", ctrl.Brand.DeleteBrand)   
	router.PUT("/brands/restore/:id", ctrl.Brand.RestoreBrand)

	router.POST("/auth/verify", ctrl.Auth.VerifyCode)
	router.POST("/auth/register", ctrl.Auth.CreateUser)                  
	router.POST("/auth/login", ctrl.Auth.LoginUser)                 

	admin := router.Group("/admin")
	admin.Use(middleware.AutoMiddleware("admin"))           
	{
		admin.GET("/dashboard", func(c *gin.Context) {
			controllers.HandleResponse(c, 200, "Welcome to the admin dashboard")
		})

		admin.GET("/products", ctrl.Product.GetProductsByFilters)      
		admin.POST("/products", ctrl.Product.CreateProduct)         
		admin.PUT("/products/:id", ctrl.Product.UpdateProduct)   
		admin.DELETE("/products/:id", ctrl.Product.DeleteProduct)
		admin.PUT("/products/restore/:id", ctrl.Product.RestoreProduct)

		admin.GET("/categories", ctrl.Category.GetCategories)       
		admin.POST("/categories", ctrl.Category.CreateCategory)     
		admin.PUT("/categories/:id", ctrl.Category.UpdateCategory)   
		admin.DELETE("/categories/:id", ctrl.Category.DeleteCategory)
		admin.PUT("/categories/restore/:id", ctrl.Category.RestoreCategory)

		admin.GET("/brands", ctrl.Brand.GetBrands)                    
		admin.POST("/brands", ctrl.Brand.CreateBrand)         
		admin.PUT("/brands/:id", ctrl.Brand.UpdateBrand)   
		admin.DELETE("/brands/:id", ctrl.Brand.DeleteBrand)
		admin.PUT("/brands/restore/:id", ctrl.Brand.RestoreBrand)
	}

	user := router.Group("/user")
	user.Use(middleware.AutoMiddleware("user"))                     
	{
		user.GET("/products", ctrl.Product.GetProductsByFilters)            
		user.GET("/products/:id", ctrl.Product.GetByID)     
		user.POST("/orders", ctrl.Order.CreateOrder)        

		user.GET("/profile", func(c *gin.Context) {
			email, _ := c.Get("email")
			userEmail := fmt.Sprintf("Welcome to your profile %s", email)
			controllers.HandleResponse(c, 200, userEmail)
		})
	}

	return router
}

































// package routes

// import (
// 	"fmt"
// 	"i-shop/controllers"
// 	"i-shop/middleware"

// 	"github.com/gin-gonic/gin"
// )

// func SetupRouter(ctrl *controllers.Controller) *gin.Engine {
// 	router := gin.Default()

// 	router.GET("/products", ctrl.Product.GetProductsByFilters)
// 	router.GET("/products/:id", ctrl.Product.GetByID)
// 	router.POST("/products", ctrl.Product.CreateProduct)
// 	router.PUT("/products/:id", ctrl.Product.UpdateProduct)
// 	router.DELETE("/products/:id", ctrl.Product.DeleteProduct)

// 	router.GET("/categories", ctrl.Category.GetCategories)
// 	router.GET("/categories/:id", ctrl.Category.GetCategoryByID)
// 	router.POST("/categories", ctrl.Category.CreateCategory)
// 	router.PUT("/categories/:id", ctrl.Category.UpdateCategory)
// 	router.DELETE("/categories/:id", ctrl.Category.DeleteCategory)

// 	router.GET("/brands", ctrl.Brand.GetBrands)
// 	router.GET("/brands/:id", ctrl.Brand.GetBrandByID)
// 	router.POST("/brands", ctrl.Brand.CreateBrand)
// 	router.PUT("/brands/:id", ctrl.Brand.UpdateBrand)
// 	router.DELETE("/brands/:id", ctrl.Brand.DeleteBrand)

// 	router.POST("/register", ctrl.Auth.CreateUser)
// 	router.POST("/login", ctrl.Auth.LoginUser)

// 	admin := router.Group("/admin")
// 	admin.Use(middleware.AutoMiddleware("admin"))
// 	admin.GET("/dashboard", func(c *gin.Context) {
// 		controllers.HandleResponse(c, 200,  "Welcome to the admin dashboard")
// 	})

// 	user := router.Group("/user")
// 	user.Use(middleware.AutoMiddleware("user"))
// 	user.GET("/profile", func(c *gin.Context) {
// 		email, _ := c.Get("email")
// 		email = fmt.Sprintf("Welcome to your profile %s", email)
// 		controllers.HandleResponse(c, 200, email)
// 	})

// 	return router

// }





