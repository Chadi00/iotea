package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func discover(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"action": "discover"})
}

func register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"action": "register"})
}

func command(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"action": "command"})
}

func receive(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"action": "receive"})
}
