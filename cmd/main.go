package main

import (
	"log"
	"os"
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
			"https://agnikriti.com",
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

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)

	err = router.Run("0.0.0.0:" + port)

	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
