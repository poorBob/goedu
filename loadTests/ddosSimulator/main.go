package main

import (
	"ddosSimulator/results"
	"ddosSimulator/services"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
)

const (
	postMessageUrl            = "http://localhost:8080/api/message"
	getMessageUrl             = "http://localhost:8080/api/message"
	getMessageWIthUuidPartUrl = "http://localhost:8080/api/messageWithUuidPart"
)

func main() {
	start := time.Now()
	concurrency := 100
	requestsPerWorker := 10
	requestCount := concurrency * requestsPerWorker

	postResultChan := make(chan results.PostMessageResult, concurrency*requestsPerWorker)
	getMessageResultChan := make(chan results.GetMessageResult, concurrency*requestsPerWorker)
	getMessagesWithUuidPartResultChan := make(chan results.GetMessagesWithUuidPartResult, concurrency*requestsPerWorker)

	requestService := services.NewLocalRequestsService()

	var wg sync.WaitGroup
	var postWg sync.WaitGroup
	var getMessageWg sync.WaitGroup
	var getMessagesWithUuidPartWg sync.WaitGroup

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for j := 0; j < requestsPerWorker; j++ {
				fmt.Printf("Worker %d, Request %d\n", workerID, j)
				uuid := uuid.New().String()

				postWg.Add(1)
				go func(workerID, requestID int, uuid string) {
					defer postWg.Done()
					requestService.PostMessage(postMessageUrl, workerID, requestID, uuid, postResultChan)
				}(workerID, j, uuid)

				getMessageWg.Add(1)
				go func(workerID, requestID int, uuid string) {
					defer getMessageWg.Done()
					requestService.GetMessage(getMessageUrl, workerID, requestID, "123e4567-e89b-12d3-a456-426655440000", getMessageResultChan)
				}(workerID, j, uuid)

				// getMessagesWithUuidPartWg.Add(1)
				// go func(workerID, requestID int, uuidPart string) {
				// 	defer getMessagesWithUuidPartWg.Done()
				// 	requestService.GetMessagesWithUuidPart(getMessageWIthUuidPartUrl, workerID, requestID, uuidPart, getMessagesWithUuidPartResultChan)
				// }(workerID, j, uuid[:3])
			}
		}(i)
	}

	go func() {
		wg.Wait()
		postWg.Wait()
		close(postResultChan)
		getMessageWg.Wait()
		close(getMessageResultChan)
		getMessagesWithUuidPartWg.Wait()
		close(getMessagesWithUuidPartResultChan)
	}()

	fmt.Println("All requests completed.")

	postMessageSuccessCount := 0
	postMessageFailResponse := []string{}
	for result := range postResultChan {
		if result.Error == nil && result.ResponseCode == http.StatusCreated {
			postMessageSuccessCount++
		} else {
			postMessageFailResponse = append(postMessageFailResponse, result.Error.Error())
		}
	}

	if postMessageSuccessCount == requestCount {
		fmt.Println("All postMessage requests were successful.")
	} else {
		fmt.Printf("Only %d out of %d postMessage requests were successful.\n", postMessageSuccessCount, requestCount)
		fmt.Println("Failed responses: ", postMessageFailResponse)
	}

	getMessageSuccessCount := 0
	getMessageFailResponse := []string{}
	for result := range getMessageResultChan {
		if result.Error == nil && result.ResponseCode == http.StatusOK {
			getMessageSuccessCount++
		} else {
			if result.Error != nil {
				getMessageFailResponse = append(getMessageFailResponse, result.Error.Error())
			} else {
				getMessageFailResponse = append(getMessageFailResponse, fmt.Sprintf("Unexpected status code: %d", result.ResponseCode))
			}
		}
	}

	if getMessageSuccessCount == requestCount {
		fmt.Println("All getMessage requests were successful.")
	} else {
		fmt.Printf("Only %d out of %d getMessage requests were successful.\n", getMessageSuccessCount, requestCount)
		fmt.Println("Failed responses: ", getMessageFailResponse)
	}

	getMessagesWithUuidPartSuccessCount := 0
	getMessagesWithUuidPartFailResponse := []string{}
	for result := range getMessagesWithUuidPartResultChan {
		if result.Error == nil && result.ResponseCode == http.StatusOK {
			getMessagesWithUuidPartSuccessCount++
		} else {
			if result.Error != nil {
				getMessagesWithUuidPartFailResponse = append(getMessagesWithUuidPartFailResponse, result.Error.Error())
			} else {
				getMessagesWithUuidPartFailResponse = append(getMessagesWithUuidPartFailResponse, fmt.Sprintf("Unexpected status code: %d", result.ResponseCode))
			}
		}
	}

	if getMessagesWithUuidPartSuccessCount == requestCount {
		fmt.Println("All getMessagesWithUuidPart requests were successful.")
	} else {
		fmt.Printf("Only %d out of %d getMessagesWithUuidPart requests were successful.\n", getMessagesWithUuidPartSuccessCount, requestCount)
		fmt.Println("Failed responses: ", getMessagesWithUuidPartFailResponse)
	}

	end := time.Now()
	fmt.Printf("Total time: %v, request count: %d\n", end.Sub(start), requestCount)
}
