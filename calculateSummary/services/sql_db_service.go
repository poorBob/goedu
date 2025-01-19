package services

import (
	"calculateSummary/models"
	"database/sql"
)

type SqlDbService struct {
	db *sql.DB
}

func NewSqlDbService(db *sql.DB) *SqlDbService {
	return &SqlDbService{db: db}
}

func (s *SqlDbService) CreateTransactionTable() error {
	_, err := s.db.Exec(`IF NOT EXISTS (SELECT * FROM sys.tables WHERE name = 'Transaction')
BEGIN
    CREATE TABLE [Transaction] (
        DateTime DATETIME NOT NULL,
        Foundation NVARCHAR(255) NOT NULL,
        FoundationMID NVARCHAR(255) NOT NULL,
        FoundationTID NVARCHAR(255) NOT NULL,
        MerchantId INT NOT NULL,
        MerchantName NVARCHAR(255) NOT NULL,
        MerchantCity NVARCHAR(255),
        MerchantStreet NVARCHAR(255),
        POSId INT NOT NULL,
        SelfServicePOS BIT NOT NULL,
        POSTransactionId INT NOT NULL,
        CashierId INT,
        DonationAmount FLOAT NOT NULL,
        SaleAmount FLOAT NOT NULL,
        DonationHash NVARCHAR(255)
    );
END`)
	if err != nil {
		return err
	}

	return nil
}

func (s *SqlDbService) GetTransactions() ([]models.Transaction, error) {

	var transactions []models.Transaction

	rows, err := s.db.Query("SELECT * FROM [Transaction]")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var t models.Transaction

		err := rows.Scan(
			&t.DateTime,
			&t.Foundation,
			&t.FoundationMID,
			&t.FoundationTID,
			&t.MerchantId,
			&t.MerchantName,
			&t.MerchantCity,
			&t.MerchantStreet,
			&t.POSId,
			&t.SelfServicePOS,
			&t.POSTransactionId,
			&t.CashierId,
			&t.DonationAmount,
			&t.SaleAmount,
			&t.DonationHash,
		)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, t)
	}

	return transactions, nil
}

func (s *SqlDbService) InsertTransactions(transactions []models.Transaction) error {

	for _, t := range transactions {
		_, err := s.db.Exec(`INSERT INTO [Transaction] VALUES (
			@DateTime,
			@Foundation,
			@FoundationMID,
			@FoundationTID,
			@MerchantId,
			@MerchantName,
			@MerchantCity,
			@MerchantStreet,
			@POSId,
			@SelfServicePOS,
			@POSTransactionId,
			@CashierId,
			@DonationAmount,
			@SaleAmount,
			@DonationHash
		)`,
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
	return nil
}

func (s *SqlDbService) DeleteTransactions() error {
	_, err := s.db.Exec("DELETE FROM [Transaction]")
	if err != nil {
		return err
	}
	return nil
}
