package handlers

import (
	"messagesApp/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DbGetByUuidPartHandler struct {
	repo repositories.MessageRepository
}

func NewDbGetByUuidPartHandler(repo repositories.MessageRepository) GetByUuidPartHandler {
	return &DbGetByUuidPartHandler{repo: repo}
}

func (h *DbGetByUuidPartHandler) Get(c echo.Context) error {
	uuidPart := c.QueryParam("uuidPart")
	if uuidPart == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Missing 'uuidPart' parameter"})
	}

	messages, err := h.repo.GetMessagesByUuidPart(uuidPart)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if len(messages) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "No messages found for the given uuid part"})
	}

	return c.JSON(http.StatusOK, messages)
}
