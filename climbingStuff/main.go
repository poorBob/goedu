package main

import (
	"database/sql"
	"fmt"
	"log"

	"climbingStuff/app"
	"climbingStuff/config"
	"climbingStuff/handlers"
	"climbingStuff/services"
	"climbingStuff/utils"
)

func main() {
	configProvider, err := config.NewViperConfigProvider(".")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	connectionString := utils.BuildConnectionString(configProvider.GetConfig())

	db := connectToDb(connectionString)
	// Services init
	gymService := services.NewClimbingGymService(connectionString)
	shoeService := services.NewSQLClimbingShoeService(db)

	// Handlers init
	gymHandler := handlers.NewDefaultClimbingGymHandler(gymService)
	shoeHandler := handlers.NewDefaultClimbingShoeHandler(shoeService)

	app := app.NewApp(db, gymService, shoeService, gymHandler, shoeHandler)
	defer app.Close()

	if err := app.Run(":8080"); err != nil {
		log.Fatalf("Failed to run app: %v", err)
	}
}

func connectToDb(connectionString string) *sql.DB {
	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		panic(fmt.Errorf("error creating connection pool: %w", err))
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		panic(fmt.Errorf("error pinging database: %w", err))
	}
	return db
}
