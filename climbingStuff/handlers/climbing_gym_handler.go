package handlers

import (
	"github.com/labstack/echo/v4"
)

type ClimbingGymHandler interface {
	GetAll(c echo.Context) error
	GetByCity(c echo.Context) error
	GetById(c echo.Context) error
	Insert(c echo.Context) error
	DeleteById(c echo.Context) error
}
