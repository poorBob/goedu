package services

import (
	"time"
)

type SummaryService interface {
	SumDonationsForFoundation(foundation string) (float32, error)
	GetDonationSumAndMinMaxDateForFoundation(foundation string) (float32, time.Time, time.Time, error)
	SumDonationsForDate(date time.Time) (float32, error)
}
