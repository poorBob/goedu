package models

type ClimbingGym struct {
	ID      uint    `gorm:"primaryKey" json:"id"`
	Name    string  `gorm:"size:255" json:"name"`
	Email   string  `gorm:"unique" json:"email"`
	Address Address `gorm:"constraint:OnDelete:CASCADE;" json:"address"`
}

func (ClimbingGym) TableName() string {
	return "ClimbingGyms"
}
