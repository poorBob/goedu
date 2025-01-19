package workers

import "messagesApp/models"

type InsertWorkerPool interface {
	Start()
	Stop()
	AddJob(msg models.Message)
}
