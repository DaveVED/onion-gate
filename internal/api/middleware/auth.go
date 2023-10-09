// internal/api/middleware/auth.go

package middleware

import (
	"github.com/DaveVED/onion-gate/internal/storage/session"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetAuthStatusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID, err := c.Cookie("session_id")
		if err != nil || !session.IsValidSession(sessionID) {
			c.Set("IsLoggedIn", false)
		} else {
			c.Set("IsLoggedIn", true)
		}
		c.Next()
	}
}

func RestrictAccessMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isLoggedIn, _ := c.Get("IsLoggedIn")
		if loggedIn, ok := isLoggedIn.(bool); !ok || !loggedIn {
			c.Redirect(http.StatusSeeOther, "/")
			c.Abort()
			return
		}
		c.Next()
	}
}
