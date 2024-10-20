package main

import (
	"access-granting/common/app"
	"access-granting/common/postgresql"
	"access-granting/controller"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	configurationManager := app.NewConfigurationManager()
	db := postgresql.GetConnection(configurationManager.PostgreSqlConfig)
	postgresql.MigrateTables(db)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	controller.InitializeRouter(db, router)
	router.Static("/profile-images", "./uploads/profile-images")

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
