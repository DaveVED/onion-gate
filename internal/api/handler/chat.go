// internal/api/handler/chat.go

package handler

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func RenderChatPage(c *gin.Context) {
	tmpl, err := template.ParseFiles("public/templates/base.html", "public/templates/partials/chat/chat.html")
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
