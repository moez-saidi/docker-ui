package main

import (
	"docker-ui/internal/app/routes"
	"log"

	"github.com/gin-gonic/gin"
)

const Port string = ":10000"

func main() {
	router := gin.Default()
	routes.RegisterRoutes(router)
	if err := router.Run(Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
