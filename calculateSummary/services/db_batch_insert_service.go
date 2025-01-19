package services

import "calculateSummary/models"

type DBBatchInsertService interface {
	Insert(transactions []models.Transaction, batchSize int) error
}
