package models

type Address struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	Street        string `gorm:"size:255" json:"street"`
	City          string `gorm:"size:255" json:"city"`
	PostalCode    string `gorm:"size:20" json:"postal_code"`
	Country       string `gorm:"size:255" json:"country"`
	ClimbingGymID uint   `gorm:"uniqueIndex" json:"climbing_gym_id"`
}

func (Address) TableName() string {
	return "Addresses"
}
