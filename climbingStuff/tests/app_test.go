package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"climbingStuff/app"
	"climbingStuff/tests/mocks"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestNewApp(t *testing.T) {
	// Mock dependencies
	mockGymService := &mocks.MockClimbingGymService{}
	mockShoeService := &mocks.MockClimbingShoeService{}
	mockGymHandler := &mocks.MockClimbingGymHandler{}
	mockShoeHandler := &mocks.MockClimbingShoeHandler{}

	mockGymHandler.GetAllFunc = func(c echo.Context) error {
		return c.JSON(http.StatusOK, []string{"Mock Gym 1", "Mock Gym 2"})
	}

	app := app.NewApp(nil, mockGymService, mockShoeService, mockGymHandler, mockShoeHandler)

	// Simulate an HTTP GET request
	req := httptest.NewRequest(http.MethodGet, "/climbing-gyms", nil)
	rec := httptest.NewRecorder()
	app.Echo.ServeHTTP(rec, req)

	// Validate the response
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `["Mock Gym 1", "Mock Gym 2"]`, rec.Body.String())
}
