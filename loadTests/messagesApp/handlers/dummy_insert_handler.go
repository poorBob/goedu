package handlers

import (
	"messagesApp/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DummyInsertHandler struct {
}

func NewDummyInsertHandler() InsertHandler {
	return &DummyInsertHandler{}
}

func (h *DummyInsertHandler) Insert(c echo.Context) error {
	var msg models.Message
	if err := c.Bind(&msg); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, msg)
}
