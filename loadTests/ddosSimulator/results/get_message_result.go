package results

import "ddosSimulator/models"

type GetMessageResult struct {
	ResponseCode int
	Message      *models.Message
	Error        error
}
