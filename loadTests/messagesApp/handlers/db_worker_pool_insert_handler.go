package handlers

import (
	"messagesApp/models"
	"messagesApp/workers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DbWOrkerPoolInsertHandler struct {
	workerPool workers.InsertWorkerPool
}

func NewDbWOrkerPoolInsertHandler(workerPool workers.InsertWorkerPool) InsertHandler {
	return &DbWOrkerPoolInsertHandler{workerPool: workerPool}
}

func (h *DbWOrkerPoolInsertHandler) Insert(c echo.Context) error {
	var msg models.Message
	if err := c.Bind(&msg); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	h.workerPool.AddJob(msg)

	// Immediately return a 'success' response
	return c.JSON(http.StatusCreated, map[string]string{"status": "queued"})
}
