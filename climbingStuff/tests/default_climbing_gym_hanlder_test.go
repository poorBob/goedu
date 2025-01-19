package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"climbingStuff/handlers"
	"climbingStuff/models"
	"climbingStuff/tests/mocks"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetByID(t *testing.T) {
	e := echo.New()

	mockService := &mocks.MockClimbingGymService{
		Gyms: []models.ClimbingGym{
			{ID: 1, Name: "Obiekto", Email: "i@obiekto.pl", Address: models.Address{Street: "Woronicza 17", City: "Warszawa", PostalCode: "55-123", Country: "Poland"}},
			{ID: 2, Name: "Bronx Bouldering", Email: "bronx@bouldering.pl", Address: models.Address{Street: "Krakowska 21", City: "Kraków", PostalCode: "20-331", Country: "Poland"}},
		},
	}

	handler := handlers.NewDefaultClimbingGymHandler(mockService)

	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/climbing-gyms/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(1))

		err := handler.GetById(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Obiekto")
	})

	t.Run("Not Found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/climbing-gyms/3", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(3))

		err := handler.GetById(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Contains(t, rec.Body.String(), "gym not found")
	})
}

func TestGetByCity(t *testing.T) {
	e := echo.New()

	mockService := &mocks.MockClimbingGymService{
		Gyms: []models.ClimbingGym{
			{ID: 1, Name: "Obiekto", Email: "i@obiekto.pl", Address: models.Address{Street: "Woronicza 17", City: "Warszawa", PostalCode: "55-123", Country: "Poland"}},
			{ID: 2, Name: "Bronx Bouldering", Email: "bronx@bouldering.pl", Address: models.Address{Street: "Krakowska 21", City: "Kraków", PostalCode: "20-331", Country: "Poland"}},
		},
	}

	handler := handlers.NewDefaultClimbingGymHandler(mockService)

	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/climbing-gyms/Warszawa", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("city")
		c.SetParamValues("Warszawa")

		err := handler.GetByCity(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Obiekto")
	})

	t.Run("Not found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/climbing-gyms/Wrocław", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("city")
		c.SetParamValues("Wrocław")

		err := handler.GetByCity(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, rec.Code)

		// Extract the error message from the JSON response
		var responseMap map[string]string
		jsonErr := json.Unmarshal(rec.Body.Bytes(), &responseMap)
		assert.NoError(t, jsonErr)

		// Check the `error` field
		errorMessage, exists := responseMap["error"]
		assert.True(t, exists)
		assert.Equal(t, "No gyms found in this city", errorMessage)
	})
}
