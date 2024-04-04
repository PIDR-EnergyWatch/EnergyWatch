package api

import (
	"backend/internal"
	"github.com/labstack/echo/v4"
	"net/http"
)

func request(c echo.Context) error {
	if c.QueryParam("measurement") == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Missing measurement query parameter"})
	}
	res := internal.QueryDB(c.QueryParam("measurement"))

	return c.JSON(http.StatusOK, res)
}
