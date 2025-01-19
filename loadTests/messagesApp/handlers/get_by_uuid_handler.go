package handlers

import "github.com/labstack/echo/v4"

type GetByUuidHandler interface {
	Get(c echo.Context) error
}
