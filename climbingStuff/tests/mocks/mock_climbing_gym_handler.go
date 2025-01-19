package mocks

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type MockClimbingGymHandler struct {
	GetAllFunc     func(c echo.Context) error
	GetByCityFunc  func(c echo.Context) error
	GetByIdFunc    func(c echo.Context) error
	InsertFunc     func(c echo.Context) error
	DeleteByIdFunc func(c echo.Context) error
}

func (m *MockClimbingGymHandler) GetAll(c echo.Context) error {
	if m.GetAllFunc != nil {
		return m.GetAllFunc(c)
	}
	return c.JSON(http.StatusOK, []string{"Mock Gym 1", "Mock Gym 2"})
}

func (m *MockClimbingGymHandler) GetByCity(c echo.Context) error {
	if m.GetByCityFunc != nil {
		return m.GetByCityFunc(c)
	}
	return c.JSON(http.StatusOK, []string{"Mock Gym in City"})
}

func (m *MockClimbingGymHandler) GetById(c echo.Context) error {
	if m.GetByIdFunc != nil {
		return m.GetByIdFunc(c)
	}
	return c.JSON(http.StatusOK, map[string]string{"id": "1", "name": "Mock Gym"})
}

func (m *MockClimbingGymHandler) Insert(c echo.Context) error {
	if m.InsertFunc != nil {
		return m.InsertFunc(c)
	}
	return c.JSON(http.StatusCreated, map[string]string{"status": "Gym created"})
}

func (m *MockClimbingGymHandler) DeleteById(c echo.Context) error {
	if m.DeleteByIdFunc != nil {
		return m.DeleteByIdFunc(c)
	}
	return c.JSON(http.StatusOK, map[string]string{"status": "Gym deleted"})
}
