package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetExample(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "example"})
}
