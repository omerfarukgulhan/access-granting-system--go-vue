package main

import (
	"access-granting/common/app"
	"access-granting/common/postgresql"
	"access-granting/controller"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	configurationManager := app.NewConfigurationManager()
	db := postgresql.GetConnection(configurationManager.PostgreSqlConfig)
	postgresql.MigrateTables(db)

	router := gin.Default()
	controller.InitializeRouter(db, router)
	router.Static("/profile-images", "./uploads/profile-images")

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
