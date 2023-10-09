// internal/api/handler/signup.go

package handler

import (
	"database/sql"
	"github.com/DaveVED/onion-gate/internal/models"
	"github.com/DaveVED/onion-gate/internal/storage/creds"
	"github.com/DaveVED/onion-gate/internal/storage/users"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

func RenderSignUpForm(c *gin.Context) {
	tmpl, err := template.ParseFiles("public/templates/base.html", "public/templates/partials/signup/signup.html")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	err = tmpl.ExecuteTemplate(c.Writer, "base.html", map[string]interface{}{
		"Title": "Signup",
	})

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

func HandleSignup(c *gin.Context) {
	db, _ := c.MustGet("db").(*sql.DB)

	username := c.PostForm("username")
	password := c.PostForm("password")

	hashedPassword, err := creds.HashPassword(password)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to hash password")
		return
	}

	user := models.User{
		Username:     username,
		PasswordHash: hashedPassword,
		DateCreated:  time.Now(),
		IsActive:     true,
	}

	err = users.InsertUser(db, &user)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to insert user")
		return
	}

	c.Redirect(http.StatusSeeOther, "/login")
}
