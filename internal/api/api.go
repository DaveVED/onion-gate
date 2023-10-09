// internal/api/api.go

package api

import (
	"github.com/DaveVED/onion-gate/internal/api/handler"
	"github.com/DaveVED/onion-gate/internal/api/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

func Start() {
	log.Println("Onion Gate is starting up...")
	router := gin.Default()

	// Global Middleware
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.DatabaseMiddleware())

	// load css
	router.Static("/public/templates", "./public/templates")

	// Login Routes
	router.GET("/login", handler.RenderLoginForm)
	router.POST("/login", handler.HandleLogin)

	// Sign up Routes
	router.GET("/signup", handler.RenderSignUpForm)
	router.POST("/signup", handler.HandleSignup)

	router.GET("/secured", middleware.AuthMiddleware(), handler.RenderSecuredPage)

	router.Run()
}
