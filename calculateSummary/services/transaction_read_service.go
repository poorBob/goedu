package services

import "calculateSummary/models"

type TransactionReadService interface {
	ReadTransactions() ([]models.Transaction, error)
}
