package services

import "calculateSummary/models"

type DBService interface {
	CreateTransactionTable() error
	GetTransactions() ([]models.Transaction, error)
	InsertTransactions([]models.Transaction) error
	DeleteTransactions() error
}
