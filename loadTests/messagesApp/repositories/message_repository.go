package repositories

import "messagesApp/models"

type MessageRepository interface {
	InsertMessage(message models.Message) (int64, error)
	InsertMessagesBatch(messages []models.Message) error
	GetMessageByUuid(uuid string) (models.Message, error)
	GetMessagesByUuidPart(uuidPart string) ([]models.Message, error)
	GetMessages() ([]models.Message, error)
	DeleteMessages() error
}
