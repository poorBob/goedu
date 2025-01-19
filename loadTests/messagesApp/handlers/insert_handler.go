package handlers

import "github.com/labstack/echo/v4"

type InsertHandler interface {
	Insert(c echo.Context) error
}
