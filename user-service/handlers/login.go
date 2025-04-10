package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func StartLogin(c *gin.Context) {
	kratosURL := os.Getenv("KRATOS_PUBLIC_URL")
	c.Redirect(http.StatusFound, kratosURL+"/self-service/login/browser")
}
