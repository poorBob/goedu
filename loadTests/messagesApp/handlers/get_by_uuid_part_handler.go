package handlers

import "github.com/labstack/echo/v4"

type GetByUuidPartHandler interface {
	Get(c echo.Context) error
}
