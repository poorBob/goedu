package mocks

import (
	"climbingStuff/models"
	"fmt"
)

type MockClimbingGymService struct {
	Gyms []models.ClimbingGym
}

func (m *MockClimbingGymService) GetAll() ([]models.ClimbingGym, error) {
	return m.Gyms, nil
}

func (m *MockClimbingGymService) GetByID(id uint) (*models.ClimbingGym, error) {
	for _, gym := range m.Gyms {
		if gym.ID == id {
			return &gym, nil
		}
	}
	return nil, fmt.Errorf("gym not found")
}

func (m *MockClimbingGymService) GetByCity(city string) ([]models.ClimbingGym, error) {
	var gyms []models.ClimbingGym
	for _, gym := range m.Gyms {
		if gym.Address.City == city {
			gyms = append(gyms, gym)
		}
	}
	return gyms, nil
}

func (m *MockClimbingGymService) Insert(gym *models.ClimbingGym) error {
	m.Gyms = append(m.Gyms, *gym)
	return nil
}

func (m *MockClimbingGymService) DeleteByID(id uint) error {
	for i, gym := range m.Gyms {
		if gym.ID == id {
			m.Gyms = append(m.Gyms[:i], m.Gyms[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("gym not found")
}
