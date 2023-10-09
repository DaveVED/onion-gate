// internal/api/handler/secure.go

package handler

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func RenderSecuredPage(c *gin.Context) {
	tmpl, err := template.ParseFiles("public/templates/base.html", "public/templates/partials/secure/secure.html")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	err = tmpl.ExecuteTemplate(c.Writer, "base.html", map[string]interface{}{
		"Title": "Secure",
	})

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
}
