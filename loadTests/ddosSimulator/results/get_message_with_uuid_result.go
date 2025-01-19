package results

import "ddosSimulator/models"

type GetMessagesWithUuidPartResult struct {
	ResponseCode int
	Messages     []models.Message
	Error        error
}
