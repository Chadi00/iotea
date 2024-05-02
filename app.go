package main

import (
	"iotea/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	routes.GenerateRoutes(server)

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	server.Run(":" + PORT)

}
