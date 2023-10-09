// internal/api/handler/signout.go

package handler

import (
	"github.com/DaveVED/onion-gate/internal/storage/session"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleSignout(c *gin.Context) {
	// Fetch the session ID from the cookie
	sessionId, err := c.Cookie("session_id")
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to get session cookie")
		return
	}

	// Invalidate the session (assuming you have a function to delete a session by its ID)
	err = session.DeleteSession(sessionId)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to delete session")
		return
	}

	// Clear the session cookie
	clearCookie := &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Path:     "/",
		Domain:   "localhost",
		MaxAge:   -1, // MaxAge negative means delete cookie now.
		HttpOnly: true,
		Secure:   false,
	}

	http.SetCookie(c.Writer, clearCookie)

	// Redirect to the login page
	c.String(http.StatusOK, "<script>window.location.href='/login';</script>")
}
