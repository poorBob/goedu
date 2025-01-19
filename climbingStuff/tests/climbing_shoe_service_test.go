package tests

import (
	"climbingStuff/models"
	"climbingStuff/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllClimbingShoes(t *testing.T) {
	mockService := &mocks.MockClimbingShoeService{
		MockGetAll: func() ([]models.ClimbingShoe, error) {
			return []models.ClimbingShoe{
				{ID: 1, Brand: "La Sportiva", Model: "Solution", Size: 42},
				{ID: 2, Brand: "Scarpa", Model: "Drago", Size: 41},
			}, nil
		},
	}

	shoes, err := mockService.GetAll()

	assert.NoError(t, err)
	assert.Len(t, shoes, 2)
	assert.Equal(t, "La Sportiva", shoes[0].Brand)
	assert.Equal(t, "Solution", shoes[0].Model)
}

func TestGetClimbingShoesByBrand(t *testing.T) {
	mockService := &mocks.MockClimbingShoeService{
		MockGetByBrand: func(brand string) ([]models.ClimbingShoe, error) {
			if brand == "Scarpa" {
				return []models.ClimbingShoe{
					{ID: 2, Brand: "Scarpa", Model: "Drago", Size: 41},
				}, nil
			}
			return []models.ClimbingShoe{}, nil
		},
	}

	shoes, err := mockService.GetByBrand("Scarpa")

	assert.NoError(t, err)
	assert.Len(t, shoes, 1)
	assert.Equal(t, "Drago", shoes[0].Model)
}

func TestAddClimbingShoe(t *testing.T) {
	mockService := &mocks.MockClimbingShoeService{
		MockAdd: func(shoe models.ClimbingShoe) (int64, error) {
			return 3, nil
		},
	}

	id, err := mockService.Add(models.ClimbingShoe{
		Brand: "Evolv",
		Model: "Phantom",
		Size:  42,
	})

	assert.NoError(t, err)
	assert.Equal(t, int64(3), id)
}
