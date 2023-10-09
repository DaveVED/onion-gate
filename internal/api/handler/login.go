// internal/api/handler/login.go

package handler

import (
    "database/sql"
	"github.com/DaveVED/onion-gate/internal/storage/session"
    "github.com/DaveVED/onion-gate/internal/storage/creds"
	"github.com/DaveVED/onion-gate/internal/storage/users"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func RenderLoginForm(c *gin.Context) {
	isLoggedIn, ok := c.Get("IsLoggedIn")
	if ok && isLoggedIn.(bool) {
		c.Redirect(http.StatusSeeOther, "/chat")
		return
	}

	tmpl, err := template.ParseFiles("public/templates/base.html", "public/templates/partials/login/login.html")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	data := prepareTemplateData(c, map[string]interface{}{
		"Title": "Login",
	})

	err = tmpl.ExecuteTemplate(c.Writer, "base.html", data)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

func HandleLogin(c *gin.Context) {
    db, _ := c.MustGet("db").(*sql.DB)

	username := c.PostForm("username")
	password := c.PostForm("password")
	authenticated := validateCreds(db, username, password)

	if authenticated {
		userSession, err := session.CreateSessionForUser(username)
		if err != nil {
			c.String(http.StatusUnauthorized, "Invalid credentials")
		}

		cookie := &http.Cookie{
			Name:     "session_id",
			Value:    userSession,
			Path:     "/",
			Domain:   "localhost",
			MaxAge:   3600,
			HttpOnly: true,
			Secure:   false,
		}

		http.SetCookie(c.Writer, cookie)

		c.String(http.StatusOK, "<script>window.location.href='/chat';</script>")
	} else {
		c.String(http.StatusUnauthorized, "Invalid credentials")
	}
}

func validateCreds(db *sql.DB, username string, password string) bool {
    passwordHash, err := users.FetchUserPassword(db, username)
    if err != nil {
        return false
    }

    if creds.CheckPasswordHash(password, passwordHash) {
        return true
    }

	return false
}
