package api

import (
	"backend/api/middleware"

	"github.com/labstack/echo/v4"
)

// Register registers all /api/ routes and applies middleware (if any)
func Register(e *echo.Group) {
	e.GET("/status", status)

	// Protected routes using JWT claims
	protectedGroup := e.Group("", middleware.JWTMiddleware())
	protectedGroup.GET("/loggedIn", loggedIn)
	protectedGroup.GET("/request", request)
}
