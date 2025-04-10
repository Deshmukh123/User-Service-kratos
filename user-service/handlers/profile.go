package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "User profile placeholder (you can check Kratos session here)",
	})
}
