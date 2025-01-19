package services

import (
	"climbingStuff/models"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type GormClimbingGymService struct {
	db *gorm.DB
}

func NewClimbingGymService(connectionString string) ClimbingGymService {
	db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Simple Migration
	err = db.AutoMigrate(&models.ClimbingGym{}, &models.Address{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return &GormClimbingGymService{db: db}
}

func (s *GormClimbingGymService) GetAll() ([]models.ClimbingGym, error) {
	var gyms []models.ClimbingGym
	err := s.db.Preload("Address").Find(&gyms).Error
	return gyms, err
}

func (s *GormClimbingGymService) GetByID(id uint) (*models.ClimbingGym, error) {
	var gym models.ClimbingGym
	err := s.db.Preload("Address").First(&gym, id).Error
	if err != nil {
		return nil, err
	}
	return &gym, nil
}

func (s *GormClimbingGymService) GetByCity(city string) ([]models.ClimbingGym, error) {
	var gyms []models.ClimbingGym
	err := s.db.
		Preload("Address").
		Joins("JOIN Addresses ON Addresses.climbing_gym_id = ClimbingGyms.id").
		Where("Addresses.city = ?", city).
		Find(&gyms).Error
	return gyms, err
}

func (s *GormClimbingGymService) Insert(gym *models.ClimbingGym) error {
	return s.db.Create(&gym).Error
}

func (s *GormClimbingGymService) DeleteByID(id uint) error {
	return s.db.Delete(&models.ClimbingGym{}, id).Error
}
