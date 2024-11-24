package main

import (
	"i-shop/config"
	"i-shop/controllers"
	"i-shop/pkg/db"
	"i-shop/pkg/db/rediss"

	"i-shop/routes"
	"i-shop/storage"
	"log"

	_ "i-shop/docs"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func main() {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",            
		DB:       0,              
	})

	redisDb := rediss.NewRedis(redisClient)
	if err := redisDb.Ping(); err != nil {
		log.Fatalf("Redis ulanish xatosi: %v", err)
	}


	cfg := config.LoadConfig()
	gin.SetMode(gin.ReleaseMode)

	conn, err := db.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect db: %v", err)
	}

	authSt := storage.NewUserStorage(conn)
	authController := controllers.AuthController{
		Storage: authSt,
		Redis: redisDb,
	}

	productSt := storage.NewProductStorage(conn)
	prController := controllers.ProductController{
		Storage: productSt,
	}

	categorySt := storage.NewCategoryStorage(conn)
	catController := controllers.CategoryController{
		Storage: categorySt,
	}

	brandSt := storage.NewBrandStorage(conn)
	brController := controllers.BrandController{
		Storage: brandSt,
	}

	router := routes.SetupRouter(&controllers.Controller{
		Brand: &brController,
		Category: &catController,
		Product: &prController,
		Auth: &authController,
	})

	router.GET("/ai-shop/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run on port: 8080 -> %v", err)
	}
}
