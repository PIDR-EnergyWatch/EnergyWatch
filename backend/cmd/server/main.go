package main

import (
	"backend/api/api"
	"backend/api/auth"
	"backend/internal"
	"net/http"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	e := echo.New()

	// CORS Allowed Origins
	origins := make([]string, 0)
	allowed := os.Getenv("ALLOWED_ORIGINS")
	for _, origin := range strings.Split(allowed, ",") {
		origins = append(origins, strings.TrimSpace(origin))
	}

	// Middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     origins,
		AllowCredentials: true,
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Skipper: nil,
		// Root directory from where the static content is served.
		Root: "build",
		// Index file for serving a directory.
		// Optional. Default value "index.html".
		Index: "login.html",
		// Enable HTML5 mode by forwarding all not-found requests to root so that
		// SPA (single-page application) can handle the routing.
		HTML5:      true,
		Browse:     false,
		IgnoreBase: false,
		Filesystem: nil,
	}))

	e.Validator = &CustomValidator{validator: validator.New()}

	// connect to the database
	internal.ConnectDB()
	defer internal.CloseDB()

	// register auth v1 routes from auth package
	authGroup := e.Group("/api/auth")
	auth.Register(authGroup)

	// register api v1 routes from api package
	apiGroup := e.Group("/api")
	api.Register(apiGroup)

	e.Logger.Fatal(e.Start(":8000"))
}
