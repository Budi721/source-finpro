package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func TestRouter(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":      true,
		"status_code": http.StatusOK,
		"message":     "Welcome to API Banking Innovation Co-Create",
	})
}
