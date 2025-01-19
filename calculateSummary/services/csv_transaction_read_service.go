package services

import (
	"bufio"
	"calculateSummary/models"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type CsvTransactionReadService struct {
	dirPath string
}

func NewCsvTransactionReadService(dirPath string) TransactionReadService {
	return &CsvTransactionReadService{dirPath: dirPath}
}

func (rs *CsvTransactionReadService) ReadTransactions() ([]models.Transaction, error) {
	dir, err := os.Open(rs.dirPath)
	if err != nil {
		log.Fatal(err)
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}

	prefix := "transactions_"

	var transactions []models.Transaction
	for _, file := range files {
		if strings.HasPrefix(file.Name(), prefix) {
			fullPath := filepath.Join(rs.dirPath, file.Name())
			ts, e := readCsvFile(fullPath)
			if e != nil {
				panic(e.Error())
			}
			transactions = append(transactions, ts...)
		}
	}

	return transactions, nil
}

func readCsvFile(filePath string) ([]models.Transaction, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	dateTimeLayout := "2006-01-02 15:04:05.0"

	var transactions []models.Transaction
	scanner := bufio.NewScanner(file)
	// skiping header
	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ";")

		parsedTime, err := time.Parse(dateTimeLayout, fields[0])
		if err != nil {
			fmt.Println("Error parsing date:", err)
			return nil, err
		}

		merchantId, err := strconv.Atoi(fields[4])
		if err != nil {
			return nil, err
		}

		posId, err := strconv.Atoi(fields[8])
		if err != nil {
			return nil, err
		}

		posTransactionId, err := strconv.Atoi(fields[10])
		if err != nil {
			return nil, err
		}

		cashierId, err := strconv.Atoi(fields[11])
		if err != nil {
			return nil, err
		}

		donationAmountInput := strings.Replace(fields[12], ",", ".", -1)
		pardedDonationAmount, err := strconv.ParseFloat(donationAmountInput, 32)
		if err != nil {
			return nil, err
		}
		donationAmount := float32(pardedDonationAmount)

		saleAmountInput := strings.Replace(fields[13], ",", ".", -1)
		parserdSaleAmount, err := strconv.ParseFloat(saleAmountInput, 32)
		if err != nil {
			return nil, err
		}
		saleAmount := float32(parserdSaleAmount)
		// TODO: error handling

		transaction := models.Transaction{
			DateTime:         parsedTime,
			Foundation:       fields[1],
			FoundationMID:    fields[2],
			FoundationTID:    fields[3],
			MerchantId:       merchantId,
			MerchantName:     fields[5],
			MerchantCity:     fields[6],
			MerchantStreet:   fields[7],
			POSId:            posId,
			SelfServicePOS:   fields[9] == "true",
			POSTransactionId: posTransactionId,
			CashierId:        cashierId,
			DonationAmount:   donationAmount,
			SaleAmount:       saleAmount,
			DonationHash:     fields[14],
		}
		transactions = append(transactions, transaction)

	}

	return transactions, nil
}

func findUniqueFoundations(transactions []models.Transaction) []string {
	foundations := make(map[string]bool)

	for _, transaction := range transactions {
		foundations[transaction.Foundation] = true
	}
	uniqueFoundations := make([]string, 0, len(foundations))
	for foundation := range foundations {
		uniqueFoundations = append(uniqueFoundations, foundation)
	}
	return uniqueFoundations
}
