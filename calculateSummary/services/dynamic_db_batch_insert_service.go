package services

import (
	"calculateSummary/models"
	"database/sql"
	"fmt"
	"strings"
)

type DynamicDBBatchInsertService struct {
	db *sql.DB
}

func NewDynamicDBBatchInsertService(db *sql.DB) *DynamicDBBatchInsertService {
	return &DynamicDBBatchInsertService{db: db}
}

func (s *DynamicDBBatchInsertService) Insert(transactions []models.Transaction, batchSize int) error {
	maxRowsPerBatch := 140
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for i := 0; i < len(transactions); {
		end := i + batchSize
		if end > len(transactions) {
			end = len(transactions)
		}

		// Extra division in case of large barch: batchsize > maxRowsPerBatch
		for j := i; j < end; j += maxRowsPerBatch {
			batchEnd := j + maxRowsPerBatch
			if batchEnd > end {
				batchEnd = end
			}

			batch := transactions[j:batchEnd]

			query := "INSERT INTO [Transaction] (DateTime, Foundation, FoundationMID, FoundationTID, MerchantId, MerchantName, MerchantCity, MerchantStreet, POSId, SelfServicePOS, POSTransactionId, CashierId, DonationAmount, SaleAmount, DonationHash) VALUES "

			values := []interface{}{}
			placeholders := []string{}

			for k, t := range batch {
				paramIndex := j + k
				placeholders = append(placeholders, fmt.Sprintf(
					"(@DateTime%d, @Foundation%d, @FoundationMID%d, @FoundationTID%d, @MerchantId%d, @MerchantName%d, @MerchantCity%d, @MerchantStreet%d, @POSId%d, @SelfServicePOS%d, @POSTransactionId%d, @CashierId%d, @DonationAmount%d, @SaleAmount%d, @DonationHash%d)",
					paramIndex, paramIndex, paramIndex, paramIndex, paramIndex, paramIndex, paramIndex, paramIndex, paramIndex, paramIndex, paramIndex, paramIndex, paramIndex, paramIndex, paramIndex))

				values = append(values,
					sql.Named(fmt.Sprintf("DateTime%d", paramIndex), t.DateTime),
					sql.Named(fmt.Sprintf("Foundation%d", paramIndex), t.Foundation),
					sql.Named(fmt.Sprintf("FoundationMID%d", paramIndex), t.FoundationMID),
					sql.Named(fmt.Sprintf("FoundationTID%d", paramIndex), t.FoundationTID),
					sql.Named(fmt.Sprintf("MerchantId%d", paramIndex), t.MerchantId),
					sql.Named(fmt.Sprintf("MerchantName%d", paramIndex), t.MerchantName),
					sql.Named(fmt.Sprintf("MerchantCity%d", paramIndex), t.MerchantCity),
					sql.Named(fmt.Sprintf("MerchantStreet%d", paramIndex), t.MerchantStreet),
					sql.Named(fmt.Sprintf("POSId%d", paramIndex), t.POSId),
					sql.Named(fmt.Sprintf("SelfServicePOS%d", paramIndex), t.SelfServicePOS),
					sql.Named(fmt.Sprintf("POSTransactionId%d", paramIndex), t.POSTransactionId),
					sql.Named(fmt.Sprintf("CashierId%d", paramIndex), t.CashierId),
					sql.Named(fmt.Sprintf("DonationAmount%d", paramIndex), t.DonationAmount),
					sql.Named(fmt.Sprintf("SaleAmount%d", paramIndex), t.SaleAmount),
					sql.Named(fmt.Sprintf("DonationHash%d", paramIndex), t.DonationHash),
				)
			}

			query += strings.Join(placeholders, ", ")

			// Execute the batch insert
			_, err := tx.Exec(query, values...)
			if err != nil {
				return err
			}
		}

		i = end
	}

	return tx.Commit()
}

func (d *DynamicDBBatchInsertService) String() string {
	return "DynamicDBBatchInsertService"
}
