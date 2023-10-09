// internal/api/handler/home.go

package handler

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func RenderHome(c *gin.Context) {
	tmpl, err := template.ParseFiles("public/templates/base.html", "public/templates/partials/home/home.html")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	data := prepareTemplateData(c, map[string]interface{}{
		"Title": "Home",
	})

	err = tmpl.ExecuteTemplate(c.Writer, "base.html", data)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
}