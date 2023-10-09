// internal/api/handler/util.go

package handler

import "github.com/gin-gonic/gin"

func prepareTemplateData(c *gin.Context, specificData map[string]interface{}) map[string]interface{} {
	isLoggedIn, _ := c.Get("IsLoggedIn")
	commonData := map[string]interface{}{
		"IsLoggedIn": isLoggedIn,
	}

	// Merge commonData and specificData
	for key, value := range specificData {
		commonData[key] = value
	}

	return commonData
}