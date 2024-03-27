package middleware

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// jwtCustomClaims are custom claims extending default ones.
// See https://github.com/golang-jwt/jwt for more examples
type JWTCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

func JWTMiddleware() echo.MiddlewareFunc {
	secret := os.Getenv("JWT_SECRET_KEY")

	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JWTCustomClaims)
		},

		ErrorHandler: func(c echo.Context, err error) error {
			return echo.ErrUnauthorized
		},

		SigningKey: []byte(secret),

		TokenLookup: "cookie:token",
	})
}
