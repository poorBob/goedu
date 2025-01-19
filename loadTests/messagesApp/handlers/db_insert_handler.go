package handlers

import (
	"net/http"

	"messagesApp/models"
	"messagesApp/repositories"

	"github.com/labstack/echo/v4"
)

type DbInsertHandler struct {
	repo repositories.MessageRepository
}

func NewDbInsertHandler(repository repositories.MessageRepository) InsertHandler {
	return &DbInsertHandler{repo: repository}
}

func (h *DbInsertHandler) Insert(c echo.Context) error {
	var msg models.Message
	if err := c.Bind(&msg); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if _, err := h.repo.InsertMessage(msg); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, msg)
}
