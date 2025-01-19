package services

import (
	"bytes"
	"ddosSimulator/models"
	"ddosSimulator/results"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type RequestsService interface {
	PostMessage(url string, workerID, requestID int, uuid string, resultChan chan<- results.PostMessageResult)
	GetMessage(url string, workerID, requestID int, uuid string, resultChan chan<- results.GetMessageResult)
	GetMessagesWithUuidPart(url string, workerID, requestID int, uuidPart string, resultChan chan<- results.GetMessagesWithUuidPartResult)
}

type LocalRequestsService struct{}

func NewLocalRequestsService() RequestsService {
	return &LocalRequestsService{}
}

func (l *LocalRequestsService) PostMessage(url string, workerID, requestID int, uuid string, resultChan chan<- results.PostMessageResult) {
	msg := models.Message{
		Uuid:     uuid,
		DateTime: time.Now(),
		Content:  fmt.Sprintf("Message from worker %d, request %d", workerID, requestID),
	}

	jsonData, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Worker %d, Request %d postMessage: Error marshaling message: %v\n", workerID, requestID, err)
		resultChan <- results.PostMessageResult{ResponseCode: 0, Error: err}
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Worker %d, Request %d postMessage: Error sending request: %v\n", workerID, requestID, err)
		resultChan <- results.PostMessageResult{ResponseCode: 0, Error: err}
		return
	}
	defer resp.Body.Close()

	log.Printf("Worker %d, Request %d: postMessage Status: %s\n", workerID, requestID, resp.Status)
	resultChan <- results.PostMessageResult{ResponseCode: resp.StatusCode, Error: nil}
}

func (l *LocalRequestsService) GetMessage(url string, workerID, requestID int, uuid string, resultChan chan<- results.GetMessageResult) {
	fullURL := fmt.Sprintf("%s?uuid=%s", url, uuid)

	resp, err := http.Get(fullURL)
	if err != nil {
		log.Printf("Worker %d, Request %d: Error sending GET request: %v\n", workerID, requestID, err)
		resultChan <- results.GetMessageResult{ResponseCode: 0, Message: nil, Error: err}
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests {
		log.Printf("Worker %d, Request %d getMessage: Received 429 Too Many Requests\n", workerID, requestID)
		resultChan <- results.GetMessageResult{ResponseCode: resp.StatusCode, Message: nil, Error: fmt.Errorf("rate limit exceeded")}
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("Worker %d, Request %d getMessage: Received unexpected status code: %d\n", workerID, requestID, resp.StatusCode)
		resultChan <- results.GetMessageResult{ResponseCode: resp.StatusCode, Message: nil, Error: fmt.Errorf("unexpected status code")}
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Worker %d, Request %d getMessage: Error reading response body: %v\n", workerID, requestID, err)
		resultChan <- results.GetMessageResult{ResponseCode: resp.StatusCode, Message: nil, Error: err}
		return
	}

	var msg models.Message
	err = json.Unmarshal(body, &msg)
	if err != nil {
		log.Printf("Worker %d, Request %d getMessage: Error unmarshaling response: %v\n", workerID, requestID, err)
		resultChan <- results.GetMessageResult{ResponseCode: resp.StatusCode, Message: nil, Error: err}
		return
	}

	// log.Printf("Worker %d, Request %d to %s: Received message: %+v\n", workerID, requestID, fullURL, msg)
	resultChan <- results.GetMessageResult{ResponseCode: resp.StatusCode, Message: &msg, Error: nil}
}

func (l *LocalRequestsService) GetMessagesWithUuidPart(url string, workerID, requestID int, uuidPart string, resultChan chan<- results.GetMessagesWithUuidPartResult) {
	fullURL := fmt.Sprintf("%s?uuidPart=%s", url, uuidPart)

	resp, err := http.Get(fullURL)
	if err != nil {
		log.Printf("Worker %d, Request %d: Error sending GET request: %v\n", workerID, requestID, err)
		resultChan <- results.GetMessagesWithUuidPartResult{ResponseCode: 0, Messages: nil, Error: err}
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Worker %d, Request %d: Received unexpected status code: %d\n", workerID, requestID, resp.StatusCode)
		resultChan <- results.GetMessagesWithUuidPartResult{ResponseCode: resp.StatusCode, Messages: nil, Error: fmt.Errorf("unexpected status code")}
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Worker %d, Request %d: Error reading response body: %v\n", workerID, requestID, err)
		resultChan <- results.GetMessagesWithUuidPartResult{ResponseCode: resp.StatusCode, Messages: nil, Error: err}
		return
	}

	var msgs []models.Message
	err = json.Unmarshal(body, &msgs)
	if err != nil {
		log.Printf("Worker %d, Request %d: Error unmarshaling response: %v\n", workerID, requestID, err)
		resultChan <- results.GetMessagesWithUuidPartResult{ResponseCode: resp.StatusCode, Messages: nil, Error: err}
		return
	}

	// log.Printf("Worker %d, Request %d to %s: Received messages: %+v\n", workerID, requestID, fullURL, msgs)
	resultChan <- results.GetMessagesWithUuidPartResult{ResponseCode: resp.StatusCode, Messages: msgs, Error: nil}
}
