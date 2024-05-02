package routes

import (
	"iotea/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func discover(ctx *gin.Context) {
	devices, err := utils.ScanNetwork()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Devices": devices})
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
