package services

import (
	"time"
)

type DBSummaryService struct {
	dbService DBService
}

func NewDBSummaryService(dbService DBService) *DBSummaryService {
	return &DBSummaryService{dbService: dbService}
}

func (s *DBSummaryService) SumDonationsForFoundation(foundation string) (float32, error) {
	transactions, err := s.dbService.GetTransactions()
	if err != nil {
		return 0, err
	}

	var sum float32
	for _, transaction := range transactions {
		if transaction.Foundation == foundation {
			sum += transaction.DonationAmount
		}
	}
	return sum, nil
}

func (s *DBSummaryService) GetDonationSumAndMinMaxDateForFoundation(foundation string) (float32, time.Time, time.Time, error) {
	transactions, err := s.dbService.GetTransactions()
	if err != nil {
		return 0, time.Time{}, time.Time{}, err
	}

	var sum float32
	var minDate time.Time
	var maxDate time.Time
	for _, transaction := range transactions {
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

func (s *DBSummaryService) SumDonationsForDate(date time.Time) (float32, error) {
	transactions, err := s.dbService.GetTransactions()
	if err != nil {
		return 0, err
	}

	var sum float32
	for _, transaction := range transactions {
		if transaction.DateTime.Year() == date.Year() &&
			transaction.DateTime.Month() == date.Month() &&
			transaction.DateTime.Day() == date.Day() {
			sum += transaction.DonationAmount
		}
	}
	return sum, nil
}
