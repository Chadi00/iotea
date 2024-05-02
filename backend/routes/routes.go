package routes

import "github.com/gin-gonic/gin"


func GenerateRoutes(server *gin.Engine){
	server.GET("/discover",discover)
	server.POST("/register", register)
	server.POST("/request", command)
	server.POST("/receive", receive)
}

