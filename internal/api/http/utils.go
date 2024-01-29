package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func JsonResult(c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, struct {
		Message string
	}{Message: message})
}

func InternalServerError(c echo.Context, err error) error {
	return JsonResult(c, http.StatusInternalServerError, err.Error())
}

func Ok(c echo.Context, message string) error {
	return JsonResult(c, http.StatusOK, message)
}
