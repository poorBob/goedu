package models

type ClimbingGymCreate struct {
	Name    string        `json:"name"`
	Email   string        `json:"email"`
	Address AddressCreate `gorm:"foreignKey:ClimbingGymID" json:"address"`
}
