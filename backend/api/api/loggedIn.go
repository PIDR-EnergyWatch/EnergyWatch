package api

import (
	"backend/api/middleware"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
)

func loggedIn(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JWTCustomClaims)
	name := claims.Name
	return c.JSON(http.StatusOK, map[string]string{"message": "Logged in as " + name + " successfully"})
}
