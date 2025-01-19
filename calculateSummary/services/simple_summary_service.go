package services

import (
	"calculateSummary/models"
	"time"
)

type SimpleSummaryService struct {
	transactions []models.Transaction
}

func NewSimpleSummaryService(t []models.Transaction) SummaryService {
	return &SimpleSummaryService{transactions: t}
}

func (s *SimpleSummaryService) SumDonationsForFoundation(foundation string) (float32, error) {
	var sum float32
	for _, transaction := range s.transactions {
		if transaction.Foundation == foundation {
			sum += transaction.DonationAmount
		}
	}
	return sum, nil
}

func (s *SimpleSummaryService) GetDonationSumAndMinMaxDateForFoundation(foundation string) (float32, time.Time, time.Time, error) {
	var sum float32
	var minDate time.Time
	var maxDate time.Time
	for _, transaction := range s.transactions {
		if transaction.Foundation == foundation {
			sum += transaction.DonationAmount
			if minDate.IsZero() || transaction.DateTime.Before(minDate) {
				minDate = transaction.DateTime
			}
			if maxDate.IsZero() || transaction.DateTime.After(maxDate) {
				maxDate = transaction.DateTime
			}
		}
	}
	return sum, minDate, maxDate, nil
}

func (s *SimpleSummaryService) SumDonationsForDate(date time.Time) (float32, error) {
	var sum float32
	for _, transaction := range s.transactions {
		if transaction.DateTime.Year() == date.Year() &&
			transaction.DateTime.Month() == date.Month() &&
			transaction.DateTime.Day() == date.Day() {
			sum += transaction.DonationAmount
		}
	}
	return sum, nil
}
