package main

import (
	"log"
	"net/http"
	"time"

	"messagesApp/handlers"
	"messagesApp/middleware"
	"messagesApp/repositories"
	"messagesApp/services"
	"messagesApp/workers"

	"github.com/labstack/echo/v4"
)

func main() {
	dbConnector := services.NewDBConnector("localhost", 1433, "goDB")
	db, err := dbConnector.Connect()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	messageRepository := repositories.NewLocalMessageRepository(db)

	// Initialize worker pool
	// workerPool := workers.NewSimpleInsertWorkerPool(messageRepository, 30)
	workerPool := workers.NewBatchInsertWorkerPool(messageRepository, 50, 20)
	workerPool.Start()
	defer workerPool.Stop()

	// Create handlers
	// insertHandler := handlers.NewDbInsertHandler(messageRepository) // less efficient
	insertHandler := handlers.NewDbWOrkerPoolInsertHandler(workerPool) // satisfies requirements
	// insertHandler := handlers.NewDummyInsertHandler() // just dummy
	getByUuidHandler := handlers.NewDbGetByUuidHandler(messageRepository)
	getByUuidPartHandler := handlers.NewDbGetByUuidPartHandler(messageRepository)

	// Echo isntance
	e := echo.New()

	// Create middleware
	rateLimiter := middleware.NewRateLimiterMiddleware(100, 200)
	stats := &middleware.SpecificRequestStats{}

	// Add middleware
	e.Use(middleware.NewSpecificRequestStatsMiddleware(stats))

	// Setup routes
	e.POST("/api/message", insertHandler.Insert)
	e.GET("/api/message", getByUuidHandler.Get, rateLimiter)
	e.GET("/api/messageWithUuidPart", getByUuidPartHandler.Get)

	e.GET("/stats", func(c echo.Context) error {
		return c.JSON(http.StatusOK, stats)
	})

	// Server configuration
	e.Server = &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  12 * time.Second,
	}

	e.Logger.Fatal(e.StartServer(e.Server))
}
