package routes

import (
	"github.com/gin-gonic/gin"
	"user-service/handlers"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/register", handlers.StartRegistration)
	r.GET("/login", handlers.StartLogin)
	r.GET("/me", handlers.GetProfile)
}
