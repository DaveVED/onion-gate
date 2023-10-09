// internal/api/middleware/auth.go

package middleware

import (
	"github.com/DaveVED/onion-gate/internal/storage/session"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID, err := c.Cookie("session_id")
		if err != nil || !session.IsValidSession(sessionID) {
			c.Redirect(http.StatusSeeOther, "/")
			c.Abort()
			return
		}
		c.Next()
	}
}
