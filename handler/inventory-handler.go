package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type InventoryHandler struct {
	inventoryServiceURL string
}

func NewInventoryHandler() *InventoryHandler {
	return &InventoryHandler{
		inventoryServiceURL: "http://inventory-service:8081",
	}
}

func (H *InventoryHandler) Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Pong",
	})
}
