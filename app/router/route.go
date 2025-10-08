package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// SetupRouter returns a *gin.Engine with all routes registered
func SetupRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	registerUserRoutes(r)

	return r
}
