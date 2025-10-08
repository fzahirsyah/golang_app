package api

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(c *gin.Context, correctUsername, correctPassword string) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if req.Username == correctUsername && req.Password == correctPassword {
		session := sessions.Default(c)
		session.Set("user", req.Username)
		session.Save()
		c.JSON(http.StatusOK, gin.H{"message": "Login successful!"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}

// SessionUserHandler returns current session user info
func SessionUserHandler(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not logged in"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// LogoutHandler destroys the session
func LogoutHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}
