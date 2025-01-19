package factory

import (
	"calculateSummary/models"
	"calculateSummary/services"
)

func SummaryServiceFactory(transcations []models.Transaction, serviceType string) services.SummaryService {

	if serviceType == "simple" {
		return services.NewSimpleSummaryService(transcations)
	}

	panic("Unknown summary service type: " + serviceType)
}
