package handlers

import (
	"net/http"

	"climbingStuff/models"
	"climbingStuff/services"

	"github.com/labstack/echo/v4"
)

type DefaultClimbingShoeHandler struct {
	service services.ClimbingShoeService
}

func NewDefaultClimbingShoeHandler(service services.ClimbingShoeService) ClimbingShoeHandler {
	return &DefaultClimbingShoeHandler{service: service}
}

// GetAll godoc
// @Summary Get all climbing shoes
// @Description Retrieves a list of all climbing shoes in the database.
// @Tags shoes
// @Accept json
// @Produce json
// @Success 200 {array} models.ClimbingShoe
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /climbing-shoes [get]
func (h *DefaultClimbingShoeHandler) GetAll(c echo.Context) error {
	shoes, err := h.service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, shoes)
}

// Insert godoc
// @Summary Add a new climbing shoe
// @Description Adds a new climbing shoe to the database.
// @Tags shoes
// @Accept json
// @Produce json
// @Param shoe body models.ClimbingShoeCreate true "Shoe creation payload"
// @Success 201 {object} models.ClimbingShoe
// @Failure 400 {object} echo.HTTPError "Invalid input"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /climbing-shoes [post]
func (h *DefaultClimbingShoeHandler) Insert(c echo.Context) error {
	var shoeCreate models.ClimbingShoeCreate
	if err := c.Bind(&shoeCreate); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	shoe := models.ClimbingShoe{
		Brand: shoeCreate.Brand,
		Model: shoeCreate.Model,
		Size:  shoeCreate.Size,
	}

	if _, err := h.service.Add(shoe); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, shoe)
}
