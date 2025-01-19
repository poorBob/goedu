package workers

import (
	"log"
	"messagesApp/models"
	"messagesApp/repositories"
)

type SimpleInsertWorkerPool struct {
	jobs    chan models.Message
	repo    repositories.MessageRepository
	workers int
}

func NewSimpleInsertWorkerPool(repo repositories.MessageRepository, workers int) InsertWorkerPool {
	return &SimpleInsertWorkerPool{
		jobs:    make(chan models.Message, 200),
		repo:    repo,
		workers: workers,
	}
}

func (wp *SimpleInsertWorkerPool) Start() {
	for i := 0; i < wp.workers; i++ {
		go wp.worker(i)
	}
}

func (wp *SimpleInsertWorkerPool) worker(id int) {
	for msg := range wp.jobs {
		_, err := wp.repo.InsertMessage(msg)
		if err != nil {
			log.Printf("Worker %d: Failed to insert message %s: %v\n", id, msg.Uuid, err)
		}
	}
}

func (wp *SimpleInsertWorkerPool) Stop() {
	close(wp.jobs)
}

func (wp *SimpleInsertWorkerPool) AddJob(msg models.Message) {
	wp.jobs <- msg
}
