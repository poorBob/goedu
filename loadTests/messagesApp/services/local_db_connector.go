package services

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb" // MS SQL Server Driver
)

type LocalBbConnector struct {
	server string
	port   int
	dbName string
	db     *sql.DB
}

func NewDBConnector(server string, port int, dbName string) DBConnector {
	return &LocalBbConnector{
		server: server,
		port:   port,
		dbName: dbName}
}

func (c *LocalBbConnector) Connect() (*sql.DB, error) {
	server := c.server
	port := c.port
	database := c.dbName

	// Connection string with Windows Authentication
	connectionString := fmt.Sprintf("server=%s;port=%d;database=%s;trusted_connection=yes", server, port, database)

	var err error
	localDb, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	// Connection test
	err = localDb.Ping()
	if err != nil {
		log.Fatal("Error pinging database: ", err.Error())
	}

	c.db = localDb

	return localDb, nil
}

func (c *LocalBbConnector) Disconnect() error {
	return c.db.Close()
}
