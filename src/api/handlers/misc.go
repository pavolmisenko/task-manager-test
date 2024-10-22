package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ToggleDarkMode(c echo.Context) error {
	// check if the cookie is already set if it is set then delete it
	if cookie, err := c.Cookie("dark-mode"); err == nil {
		cookie.Value = ""
		cookie.MaxAge = -1
		c.SetCookie(cookie)
		return c.JSON(http.StatusOK, "deleted")
	}

	// if the cookie is not set then set it
	cookie := new(http.Cookie)
	cookie.Name = "dark-mode"
	cookie.Value = "true"
	cookie.Expires = cookie.Expires.AddDate(1, 0, 0)
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, "set")
}
