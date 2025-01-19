package services

import "climbingStuff/models"

type ClimbingGymService interface {
	GetAll() ([]models.ClimbingGym, error)
	GetByID(id uint) (*models.ClimbingGym, error)
	GetByCity(city string) ([]models.ClimbingGym, error)
	Insert(gym *models.ClimbingGym) error
	DeleteByID(id uint) error
}
