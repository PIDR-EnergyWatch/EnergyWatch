package auth

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func logout(c echo.Context) error {
	cookie := &http.Cookie{
		Name:   "token",
		Path:   "/",
		MaxAge: -1,
	}

	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, map[string]string{"message": "Logged out successfully"})
}
