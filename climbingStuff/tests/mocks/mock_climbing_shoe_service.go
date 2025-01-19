package mocks

import (
	"climbingStuff/models"
)

type MockClimbingShoeService struct {
	MockGetAll     func() ([]models.ClimbingShoe, error)
	MockGetByBrand func(brand string) ([]models.ClimbingShoe, error)
	MockAdd        func(shoe models.ClimbingShoe) (int64, error)
}

func (m *MockClimbingShoeService) GetAll() ([]models.ClimbingShoe, error) {
	return m.MockGetAll()
}

func (m *MockClimbingShoeService) GetByBrand(brand string) ([]models.ClimbingShoe, error) {
	return m.MockGetByBrand(brand)
}

func (m *MockClimbingShoeService) Add(shoe models.ClimbingShoe) (int64, error) {
	return m.MockAdd(shoe)
}
