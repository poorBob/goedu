package models

import (
	"time"
)

type Transaction struct {
	DateTime         time.Time
	Foundation       string
	FoundationMID    string
	FoundationTID    string
	MerchantId       int
	MerchantName     string
	MerchantCity     string
	MerchantStreet   string
	POSId            int
	SelfServicePOS   bool
	POSTransactionId int
	CashierId        int
	DonationAmount   float32
	SaleAmount       float32
	DonationHash     string
}
