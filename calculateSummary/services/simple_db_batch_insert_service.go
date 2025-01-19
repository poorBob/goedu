package services

import (
	"calculateSummary/models"
	"database/sql"
	"log"
)

type SimpleDBBatchInsertService struct {
	db *sql.DB
}

func NewSimpleDBBatchInsertService(db *sql.DB) *SimpleDBBatchInsertService {
	return &SimpleDBBatchInsertService{db: db}
}

func (s *SimpleDBBatchInsertService) Insert(transactions []models.Transaction, batchSize int) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare(`
        INSERT INTO [Transaction] (
            DateTime,
            Foundation,
            FoundationMID,
            FoundationTID,
            MerchantId,
            MerchantName,
            MerchantCity,
            MerchantStreet,
            POSId,
            SelfServicePOS,
            POSTransactionId,
            CashierId,
            DonationAmount,
            SaleAmount,
            DonationHash
        ) VALUES 
        (@DateTime, @Foundation, @FoundationMID, @FoundationTID, @MerchantId, @MerchantName, @MerchantCity, @MerchantStreet, @POSId, @SelfServicePOS, @POSTransactionId, @CashierId, @DonationAmount, @SaleAmount, @DonationHash)
    `)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for i := 0; i < len(transactions); i += batchSize {
		end := i + batchSize
		if end > len(transactions) {
			end = len(transactions)
		}
		batch := transactions[i:end]

		log.Printf("Executing batch with %d transactions", len(batch))

		for _, t := range batch {
			_, err := stmt.Exec(
				sql.Named("DateTime", t.DateTime),
				sql.Named("Foundation", t.Foundation),
				sql.Named("FoundationMID", t.FoundationMID),
				sql.Named("FoundationTID", t.FoundationTID),
				sql.Named("MerchantId", t.MerchantId),
				sql.Named("MerchantName", t.MerchantName),
				sql.Named("MerchantCity", t.MerchantCity),
				sql.Named("MerchantStreet", t.MerchantStreet),
				sql.Named("POSId", t.POSId),
				sql.Named("SelfServicePOS", t.SelfServicePOS),
				sql.Named("POSTransactionId", t.POSTransactionId),
				sql.Named("CashierId", t.CashierId),
				sql.Named("DonationAmount", t.DonationAmount),
				sql.Named("SaleAmount", t.SaleAmount),
				sql.Named("DonationHash", t.DonationHash),
			)
			if err != nil {
				return err
			}
		}
	}

	return tx.Commit()
}

func (s *SimpleDBBatchInsertService) String() string {
	return "SimpleDBBatchInsertService"
}
