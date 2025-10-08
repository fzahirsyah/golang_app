package router

import (
	"golang-app/app/api"
	"os"

	"github.com/gin-gonic/gin"
)

func registerUserRoutes(r *gin.Engine) {
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	// Register login route with credentials from .env
	r.POST("/login", func(c *gin.Context) {
		api.LoginHandler(c, username, password)
	})
}
