package main

import (
	"docker-ui/internal/app/routes"
	"docker-ui/internal/db"
	"log"

	"github.com/gin-gonic/gin"
)

const Port string = ":10000"

func main() {
	database := db.Init()
	defer func() {
		sqlDB, err := database.DB()
		if err != nil {
			log.Fatalf("Failed to close database connection: %v", err)
		}
		sqlDB.Close()
	}()
	router := gin.Default()
	routes.RegisterRoutes(router)

	router.Run(Port) // nolint:errcheck
}
