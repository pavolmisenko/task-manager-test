package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func FileSearch(context echo.Context) error {
	return context.Render(http.StatusOK, "file_search", nil)
}