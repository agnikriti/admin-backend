package main

import (
	"log"
	"time"

	"agnikriti_admin_backend/config"
	"agnikriti_admin_backend/database"
	"agnikriti_admin_backend/internet_services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	log.Println("Loading environment configuration...")
	config.LoadConfig()

	log.Println("Connecting to database...")
	database.ConnectDB()

	log.Println("Initializing Gin server...")
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	err := router.SetTrustedProxies(nil)

	if err != nil {
		log.Fatal("Failed to configure trusted proxies:", err)
	}

	log.Println("Applying CORS configuration...")

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},

		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},

		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},

		ExposeHeaders: []string{
			"Content-Length",
		},

		AllowCredentials: true,

		MaxAge: 12 * time.Hour,
	}))

	log.Println("Registering application routes...")

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"message": "Server is running",
		})
	})

	internet_services.RegisterRoutes(router)

	log.Printf(
		"Server running on http://localhost:%s",
		config.AppConfig.PORT,
	)

	err = router.Run(":" + config.AppConfig.PORT)

	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
