package routes

import (
	"fmt"
	"i-shop/controllers"
	"i-shop/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(ctrl *controllers.Controller) *gin.Engine {
	router := gin.Default()

	// Maxsulotlar bilan ishlash uchun umumiy endpointlar
	router.GET("/products", ctrl.Product.GetProductsByFilters)         // User va admin uchun maxsulotlar ro'yxati
	router.GET("/products/:id", ctrl.Product.GetByID)                 // Maxsulotni ID bo'yicha ko'rish
	router.POST("/products", ctrl.Product.CreateProduct)              // Admin uchun maxsulot qo'shish
	router.PUT("/products/:id", ctrl.Product.UpdateProduct)           // Admin uchun maxsulot yangilash
	router.DELETE("/products/:id", ctrl.Product.DeleteProduct)        // Admin uchun maxsulot o'chirish
	router.PUT("/products/restore/:id", ctrl.Product.RestoreProduct)

	// Kategoriyalar bilan ishlash uchun umumiy endpointlar
	router.GET("/categories", ctrl.Category.GetCategories)            // User va admin uchun kategoriyalar ro'yxati
	router.GET("/categories/:id", ctrl.Category.GetCategoryByID)      // Kategoriya ID bo'yicha ko'rish
	router.POST("/categories", ctrl.Category.CreateCategory)         // Admin uchun yangi kategoriya qo'shish
	router.PUT("/categories/:id", ctrl.Category.UpdateCategory)      // Admin uchun kategoriya yangilash
	router.DELETE("/categories/:id", ctrl.Category.DeleteCategory)   // Admin uchun kategoriya o'chirish
	router.PUT("/categories/restore/:id", ctrl.Category.RestoreCategory)

	// Brendlar bilan ishlash uchun umumiy endpointlar
	router.GET("/brands", ctrl.Brand.GetBrands)                       // User va admin uchun brendlar ro'yxati
	router.GET("/brands/:id", ctrl.Brand.GetBrandByID)               // Brendni ID bo'yicha ko'rish
	router.POST("/brands", ctrl.Brand.CreateBrand)                   // Admin uchun yangi brend qo'shish
	router.PUT("/brands/:id", ctrl.Brand.UpdateBrand)                // Admin uchun brend yangilash
	router.DELETE("/brands/:id", ctrl.Brand.DeleteBrand)             // Admin uchun brend o'chirish
	router.PUT("/brands/restore/:id", ctrl.Brand.RestoreBrand)

	// Foydalanuvchi ro'yxatdan o'tish va login qilish endpointlari
	router.POST("/auth/verify", ctrl.Auth.VerifyCode)
	router.POST("/auth/register", ctrl.Auth.CreateUser)                   // Foydalanuvchi ro'yxatdan o'tish
	router.POST("/auth/login", ctrl.Auth.LoginUser)                       // Foydalanuvchi login qilish

	// Admin guruhini yaratish va admin roli uchun maxsus endpointlar
	admin := router.Group("/admin")
	admin.Use(middleware.AutoMiddleware("admin"))                    // Admin rolini tekshirish
	{
		admin.GET("/dashboard", func(c *gin.Context) {
			controllers.HandleResponse(c, 200, "Welcome to the admin dashboard")
		})

		// Admin uchun maxsulotlarni boshqarish endpointlari
		admin.GET("/products", ctrl.Product.GetProductsByFilters)           // Barcha maxsulotlarni ko'rsatish
		admin.POST("/products", ctrl.Product.CreateProduct)           // Yangi maxsulot qo'shish
		admin.PUT("/products/:id", ctrl.Product.UpdateProduct)        // Maxsulotni yangilash
		admin.DELETE("/products/:id", ctrl.Product.DeleteProduct)     // Maxsulotni o'chirish
		admin.PUT("/products/restore/:id", ctrl.Product.RestoreProduct)

		// Admin uchun kategoriya va brendlarni boshqarish endpointlari
		admin.GET("/categories", ctrl.Category.GetCategories)         // Barcha kategoriyalarni ko'rsatish
		admin.POST("/categories", ctrl.Category.CreateCategory)       // Yangi kategoriya qo'shish
		admin.PUT("/categories/:id", ctrl.Category.UpdateCategory)    // Kategoriya yangilash
		admin.DELETE("/categories/:id", ctrl.Category.DeleteCategory) // Kategoriya o'chirish
		admin.PUT("/categories/restore/:id", ctrl.Category.RestoreCategory)

		admin.GET("/brands", ctrl.Brand.GetBrands)                    // Barcha brendlarni ko'rsatish
		admin.POST("/brands", ctrl.Brand.CreateBrand)                 // Yangi brend qo'shish
		admin.PUT("/brands/:id", ctrl.Brand.UpdateBrand)              // Brendni yangilash
		admin.DELETE("/brands/:id", ctrl.Brand.DeleteBrand)           // Brendni o'chirish
		admin.PUT("/brands/restore/:id", ctrl.Brand.RestoreBrand)
	}

	// User guruhini yaratish va user roli uchun maxsus endpointlar
	user := router.Group("/user")
	user.Use(middleware.AutoMiddleware("user"))                     // User rolini tekshirish
	{
		// User uchun maxsulotlarni ko'rish va sotib olish endpointlari
		user.GET("/products", ctrl.Product.GetProductsByFilters)            // Maxsulotlarni category va brand bo'yicha filtrlash
		user.GET("/products/:id", ctrl.Product.GetByID)        // Maxsulotni ID bo'yicha ko'rish
		user.POST("/order", ctrl.Order.CreateOrder)                   // Maxsulotni sotib olish

		// User uchun profilni ko'rsatish
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





