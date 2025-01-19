package services

import (
	"climbingStuff/models"
	"database/sql"
	"fmt"
)

type SQLClimbingShoeService struct {
	db *sql.DB
}

func NewSQLClimbingShoeService(db *sql.DB) ClimbingShoeService {
	return &SQLClimbingShoeService{db: db}
}

func (s *SQLClimbingShoeService) GetAll() ([]models.ClimbingShoe, error) {
	var shoes []models.ClimbingShoe
	rows, err := s.db.Query("SELECT * FROM ClimbingShoes")
	if err != nil {
		return nil, fmt.Errorf("error while fetching shoes: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var shoe models.ClimbingShoe
		if err := rows.Scan(&shoe.ID, &shoe.Brand, &shoe.Model, &shoe.Size); err != nil {
			return nil, fmt.Errorf("error while scanning shoes: %v", err)
		}
		shoes = append(shoes, shoe)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error from rows: %v", err)
	}
	return shoes, nil
}

func (s *SQLClimbingShoeService) GetByBrand(brand string) ([]models.ClimbingShoe, error) {
	var shoes []models.ClimbingShoe
	rows, err := s.db.Query("SELECT * FROM ClimbingShoes WHERE Brand = @brand", sql.Named("brand", brand))
	if err != nil {
		return nil, fmt.Errorf("error while fetching shoes by brand: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var shoe models.ClimbingShoe
		if err := rows.Scan(&shoe.ID, &shoe.Brand, &shoe.Model, &shoe.Size); err != nil {
			return nil, fmt.Errorf("error while scanning shoes: %v", err)
		}
		shoes = append(shoes, shoe)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error from rows: %v", err)
	}
	return shoes, nil
}

func (s *SQLClimbingShoeService) Add(shoe models.ClimbingShoe) (int64, error) {
	var id int64
	exists, err := s.shoeExists(shoe)
	if err != nil {
		return -1, fmt.Errorf("error checking shoe existence: %v", err)
	}
	if exists {
		return id, nil
	}

	err = s.db.QueryRow(
		"INSERT INTO ClimbingShoes (Brand, Model, Size) OUTPUT INSERTED.ID VALUES (@brand, @model, @size)",
		sql.Named("brand", shoe.Brand),
		sql.Named("model", shoe.Model),
		sql.Named("size", shoe.Size),
	).Scan(&id)
	if err != nil {
		return -1, fmt.Errorf("error inserting shoe: %v", err)
	}
	return id, nil
}

func (s *SQLClimbingShoeService) shoeExists(shoe models.ClimbingShoe) (bool, error) {
	var id int64
	err := s.db.QueryRow(
		"SELECT ID FROM ClimbingShoes WHERE Brand = @brand AND Model = @model",
		sql.Named("brand", shoe.Brand),
		sql.Named("model", shoe.Model),
	).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("error checking shoe existence: %v", err)
	}
	return true, nil
}
