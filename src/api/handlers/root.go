package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Root(c echo.Context) error {
	return c.Redirect(http.StatusPermanentRedirect, "/tasks")
}
