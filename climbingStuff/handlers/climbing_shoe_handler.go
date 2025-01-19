package handlers

import "github.com/labstack/echo/v4"

type ClimbingShoeHandler interface {
	GetAll(c echo.Context) error
	Insert(c echo.Context) error
}
