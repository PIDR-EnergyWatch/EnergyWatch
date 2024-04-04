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
	internal.QueryDB(c.QueryParam("measurement"))

	return c.JSON(http.StatusOK, map[string]string{"message": "J'ai les donnée, j'ai les datas"})
}
