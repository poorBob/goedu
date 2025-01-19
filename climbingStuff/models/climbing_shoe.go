package models

type ClimbingShoe struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Brand string `gorm:"size:255" json:"brand"`
	Model string `gorm:"size:255" json:"model"`
	Size  int    `json:"size"`
}

func (ClimbingShoe) TableName() string {
	return "ClimbingShoes"
}
