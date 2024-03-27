package auth

import "github.com/labstack/echo/v4"

// Register registers all /api/auth/ routes and applies middleware (if any)
func Register(e *echo.Group) {
	e.POST("/login", login)
}
