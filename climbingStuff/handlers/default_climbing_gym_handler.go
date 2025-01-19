package handlers

import (
	"net/http"
	"strconv"

	"climbingStuff/models"
	"climbingStuff/services"

	"github.com/labstack/echo/v4"
)

type DefaultClimbingGymHandler struct {
	service services.ClimbingGymService
}

func NewDefaultClimbingGymHandler(service services.ClimbingGymService) ClimbingGymHandler {
	return &DefaultClimbingGymHandler{service: service}
}

// GetAll godoc
// @Summary Get all climbing gyms
// @Description Get a list of all climbing gyms
// @Tags gyms
// @Accept json
// @Produce json
// @Success 200 {array} models.ClimbingGym
// @Failure 500 {object} echo.HTTPError
// @Router /climbing-gyms [get]
func (h *DefaultClimbingGymHandler) GetAll(c echo.Context) error {
	gyms, err := h.service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, gyms)
}

// GetByID godoc
// @Summary Get a climbing gym by ID
// @Description Fetches a climbing gym record based on the provided ID.
// @Tags gyms
// @Accept json
// @Produce json
// @Param id path int true "Climbing Gym ID"
// @Success 200 {object} models.ClimbingGym
// @Failure 400 {object} echo.HTTPError "Invalid ID format"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /climbing-gyms/{id} [get]
func (h *DefaultClimbingGymHandler) GetById(c echo.Context) error {
	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	gym, err := h.service.GetByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, gym)
}

// GetByCity godoc
// @Summary Get climbing gyms by city
// @Description Fetches a list of climbing gyms located in the specified city.
// @Tags gyms
// @Accept json
// @Produce json
// @Param city path string true "City name"
// @Success 200 {array} models.ClimbingGym
// @Failure 404 {object} echo.HTTPError "No gyms found in this city"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /climbing-gyms/{city} [get]
func (h *DefaultClimbingGymHandler) GetByCity(c echo.Context) error {
	cityParam := c.Param("city")

	gyms, err := h.service.GetByCity(cityParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if len(gyms) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "No gyms found in this city"})
	}

	return c.JSON(http.StatusOK, gyms)
}

// Insert godoc
// @Summary Add a new climbing gym
// @Description Adds a new climbing gym to the database.
// @Tags gyms
// @Accept json
// @Produce json
// @Param gym body models.ClimbingGymCreate true "Gym creation payload"
// @Success 201 {object} models.ClimbingGym
// @Failure 400 {object} echo.HTTPError "Invalid input"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /climbing-gyms [post]
func (h *DefaultClimbingGymHandler) Insert(c echo.Context) error {
	var gymCreate models.ClimbingGymCreate
	if err := c.Bind(&gymCreate); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	add := models.Address{
		Street:     gymCreate.Address.Street,
		City:       gymCreate.Address.City,
		PostalCode: gymCreate.Address.PostalCode,
		Country:    gymCreate.Address.Country,
	}

	gym := models.ClimbingGym{
		Name:    gymCreate.Name,
		Address: add,
		Email:   gymCreate.Email,
	}

	if err := h.service.Insert(&gym); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, gym)
}

// DeleteByID godoc
// @Summary Delete a climbing gym by ID
// @Description Deletes a climbing gym record based on the provided ID.
// @Tags gyms
// @Param id path int true "Climbing Gym ID"
// @Success 204 "No Content"
// @Failure 400 {object} echo.HTTPError "Invalid ID format"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /climbing-gyms/{id} [delete]
func (h *DefaultClimbingGymHandler) DeleteById(c echo.Context) error {
	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	if err := h.service.DeleteByID(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusNoContent, nil)
}
