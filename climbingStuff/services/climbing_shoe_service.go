package services

import "climbingStuff/models"

type ClimbingShoeService interface {
	GetAll() ([]models.ClimbingShoe, error)
	GetByBrand(brand string) ([]models.ClimbingShoe, error)
	Add(shoe models.ClimbingShoe) (int64, error)
}
