package workers

import (
	"fmt"
	"log"
	"messagesApp/models"
	"messagesApp/repositories"
	"time"
)

type BatchInsertWorkerPool struct {
	jobs      chan models.Message
	repo      repositories.MessageRepository
	workers   int
	batchSize int
}

func NewBatchInsertWorkerPool(repo repositories.MessageRepository, workers, batchSize int) InsertWorkerPool {
	return &BatchInsertWorkerPool{
		jobs:      make(chan models.Message, 500),
		repo:      repo,
		workers:   workers,
		batchSize: batchSize,
	}
}

func (wp *BatchInsertWorkerPool) Start() {
	for i := 0; i < wp.workers; i++ {
		go wp.worker(i)
	}
}

func (wp *BatchInsertWorkerPool) Stop() {
	close(wp.jobs)
}

func (wp *BatchInsertWorkerPool) AddJob(msg models.Message) {
	wp.jobs <- msg
}

func (wp *BatchInsertWorkerPool) worker(id int) {
	buffer := make([]models.Message, 0, wp.batchSize)
	ticker := time.NewTicker(5 * time.Second) // Timeout evey 5 seconds
	defer ticker.Stop()

	for {
		select {
		case msg, ok := <-wp.jobs:
			if !ok {
				// Channel closed: immediately insert remaining messages
				if len(buffer) > 0 {
					wp.insertBatch(buffer)
				}
				fmt.Printf("Worker %d: Stopped!\n", id)
				return
			}

			buffer = append(buffer, msg)
			if len(buffer) >= wp.batchSize {
				wp.insertBatch(buffer)
				buffer = buffer[:0]
			}
		case <-ticker.C:
			// Handle timeout: insert remaining messages
			if len(buffer) > 0 {
				wp.insertBatch(buffer)
				buffer = buffer[:0]
			}
		}
	}
}

func (wp *BatchInsertWorkerPool) insertBatch(messages []models.Message) {
	if len(messages) == 0 {
		return
	}
	err := wp.repo.InsertMessagesBatch(messages)
	if err != nil {
		log.Printf("Batch insert failed: %v\n", err)
	}
}
