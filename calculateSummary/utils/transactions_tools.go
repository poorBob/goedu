package utils

import (
	"calculateSummary/models"
	"time"
)

func GetDatesOnlyFromTransactions(transactions []models.Transaction) []time.Time {
	dateMap := make(map[string]struct{})
	var uniqueDates []time.Time

	for _, transaction := range transactions {
		dateOnly := transaction.DateTime.Truncate(24 * time.Hour)

		dateKey := dateOnly.Format("2000-01-02")

		if _, exists := dateMap[dateKey]; !exists {
			dateMap[dateKey] = struct{}{}
			uniqueDates = append(uniqueDates, dateOnly)
		}
	}

	return uniqueDates
}

func GetFoundationsFromTransactions(transactions []models.Transaction) []string {
	foundations := make(map[string]struct{})
	var uniqueFoundations []string

	for _, transaction := range transactions {
		if _, exists := foundations[transaction.Foundation]; !exists {
			foundations[transaction.Foundation] = struct{}{}
			uniqueFoundations = append(uniqueFoundations, transaction.Foundation)
		}
	}

	return uniqueFoundations
}
