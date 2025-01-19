package services

import (
	"database/sql"
)

type DBConnector interface {
	Connect() (*sql.DB, error)
	Disconnect() error
}
