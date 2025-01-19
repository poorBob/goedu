package handlers

import (
	"messagesApp/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DbGetByUuidHandler struct {
	repo repositories.MessageRepository
}

func NewDbGetByUuidHandler(repo repositories.MessageRepository) GetByUuidHandler {
	return &DbGetByUuidHandler{repo: repo}
}

func (h *DbGetByUuidHandler) Get(c echo.Context) error {
	uuid := c.QueryParam("uuid")
	if uuid == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Missing 'uuid' parameter"})
	}

	message, err := h.repo.GetMessageByUuid(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if message.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "No message found for the given uuid"})
	}

	return c.JSON(http.StatusOK, message)
}
