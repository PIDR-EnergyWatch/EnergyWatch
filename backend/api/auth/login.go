package auth

import (
	"backend/api/middleware"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type loginData struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func login(c echo.Context) error {
	var creds loginData
	if err := c.Bind(&creds); err != nil {
		return err
	}
	if err := c.Validate(creds); err != nil {
		return err
	}

	// TODO: Logic to check username/password from database/service
	if creds.Username != os.Getenv("DASHBOARD_USER") || creds.Password != os.Getenv("DASHBOARD_PASSWORD") {
		return echo.ErrUnauthorized
	}

	// Set custom claims and sign token for client to use for authentication
	claims := &middleware.JWTCustomClaims{
		Name:  "Svelte Echo User",
		Admin: true,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET_KEY")
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return err
	}

	cookie := &http.Cookie{
		Name:     "token",
		Value:    t,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 72),
	}

	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]string{"message": "Logged in successfully"})
}
