// internal/api/middleware/database.go

package middleware

import (
	"github.com/DaveVED/onion-gate/internal/storage"
	"github.com/gin-gonic/gin"
	"log"
)

func DatabaseMiddleware() gin.HandlerFunc {
	db, err := storage.DBInit()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
