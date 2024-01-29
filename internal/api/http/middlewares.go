package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (server *Server) GetUserId(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var userId = c.Request().Header.Get("userId")
		if userId == "" {
			return JsonResult(c, http.StatusBadRequest, "userId should set in request header")
		}

		c.Set("userId", userId)

		return next(c)
	}
}
