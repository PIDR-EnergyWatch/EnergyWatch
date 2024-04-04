package auth

import (
	"backend/api/middleware"
	"github.com/labstack/echo/v4"
)

// Register registers all /api/auth/ routes and applies middleware (if any)
func Register(e *echo.Group) {
	e.POST("/login", login)

	// Protected routes using JWT claims
	protectedGroup := e.Group("", middleware.JWTMiddleware())
	protectedGroup.POST("/logout", logout)
}
