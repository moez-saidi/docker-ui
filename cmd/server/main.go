package main

import (
	"docker-ui/internal/app/routes"

	"github.com/gin-gonic/gin"
)

const Port string = ":10000"

func main() {
	router := gin.Default()
	routes.RegisterRoutes(router)
	router.Run(Port)
}
