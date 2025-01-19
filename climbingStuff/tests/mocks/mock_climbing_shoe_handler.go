package mocks

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type MockClimbingShoeHandler struct {
	GetAllFunc func(c echo.Context) error
	InsertFunc func(c echo.Context) error
}

func (m *MockClimbingShoeHandler) GetAll(c echo.Context) error {
	if m.GetAllFunc != nil {
		return m.GetAllFunc(c)
	}
	return c.JSON(http.StatusOK, []string{"Mock Shoe 1", "Mock Shoe 2"})
}

func (m *MockClimbingShoeHandler) Insert(c echo.Context) error {
	if m.InsertFunc != nil {
		return m.InsertFunc(c)
	}
	return c.JSON(http.StatusCreated, map[string]string{"status": "Shoe created"})
}
