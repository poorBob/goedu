package main

import (
	"calculateSummary/factory"
	"calculateSummary/services"
	"calculateSummary/utils"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/denisenkom/go-mssqldb" // MS SQL Server driver
)

func main() {
	db := connectToDb()
	defer db.Close()

	// Services initialization
	readService := services.NewCsvTransactionReadService("./input_data")
	dbService := GetDbService(db)

	transactions, err := readService.ReadTransactions()
	if err != nil {
		log.Fatal("Error pinging database: ", err.Error())
	}
	fmt.Println(len(transactions))

	err = dbService.DeleteTransactions()
	if err != nil {
		log.Fatal("Error deleting transactions: ", err.Error())
	}

	err = dbService.CreateTransactionTable()
	if err != nil {
		log.Fatal("Error creating transaction table: ", err.Error())
	}

	err = dbService.InsertTransactions(transactions)
	if err != nil {
		log.Fatal("Error inserting transactions: ", err.Error())
	}

	// Run some benchmarks
	batchInsertServices := []services.DBBatchInsertService{
		services.NewSimpleDBBatchInsertService(db),
		services.NewDynamicDBBatchInsertService(db),
	}

	for _, service := range batchInsertServices {
		err = dbService.DeleteTransactions()
		if err != nil {
			log.Fatal("Error deleting transactions: ", err.Error())
		}

		start := time.Now()
		err := service.Insert(transactions, 100)
		if err != nil {
			log.Printf("Failed to create transaction table: %v", err)
		} else {
			log.Println("Transaction table created successfully!")
		}
		duration := time.Since(start)
		fmt.Printf("Executing batch with the %s service. Execution time: %v\n", service, duration)
	}

	// Calculate some summaries
	simpleSummaryService := factory.SummaryServiceFactory(transactions, "simple")
	dbSummaryService := services.NewDBSummaryService(dbService)
	dates := utils.GetDatesOnlyFromTransactions(transactions)
	foundations := utils.GetFoundationsFromTransactions(transactions)

	summaryServices := []services.SummaryService{
		simpleSummaryService,
		dbSummaryService,
	}

	for _, service := range summaryServices {

		fmt.Println("----")
		fmt.Printf("Summaries for service %T\n", service)

		for _, date := range dates {
			s, _ := service.SumDonationsForDate(date)
			fmt.Printf("Sum of donations for date %v: %v\n", date, s)
		}

		for _, foundation := range foundations {
			s, err := service.SumDonationsForFoundation(foundation)
			if err != nil {
				log.Fatal("Error pinging database: ", err.Error())
			}
			fmt.Printf("Sum of donations for foundation %v: %v\n", foundation, s)
		}
		fmt.Println("----")
	}
}

func connectToDb() *sql.DB {
	server := "localhost"
	port := 1433
	database := "goDB"

	// Connection string with Windows Authentication
	connectionString := fmt.Sprintf("server=%s;port=%d;database=%s;trusted_connection=yes", server, port, database)

	var err error
	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	// Connection test
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database: ", err.Error())
	}

	fmt.Println("Connected with DB with use of Windows Authentication!")
	return db
}

func GetDbService(db *sql.DB) services.DBService {
	return services.NewSqlDbService(db)
}
