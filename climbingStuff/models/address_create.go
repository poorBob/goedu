package models

type AddressCreate struct {
	Street     string `gorm:"size:255" json:"street"`
	City       string `gorm:"size:255" json:"city"`
	PostalCode string `gorm:"size:20" json:"postal_code"`
	Country    string `gorm:"size:255" json:"country"`
}
