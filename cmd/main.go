package main

import (
	"gateway/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	inventoryHandler := handler.NewInventoryHandler()

	router.GET("/ping", inventoryHandler.Pong)
	router.Run("0.0.0.0:8080")
}
