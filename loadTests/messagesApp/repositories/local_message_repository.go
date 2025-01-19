package repositories

import (
	"database/sql"
	"fmt"
	"messagesApp/models"
)

type LocalMessageRepository struct {
	db *sql.DB
}

func NewLocalMessageRepository(db *sql.DB) MessageRepository {
	repo := &LocalMessageRepository{db: db}

	err := repo.init()
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize Messages table: %v", err))
	}

	return repo
}

func (l *LocalMessageRepository) init() error {
	return checkAndCreateMessagesTable(l.db)
}

func (l *LocalMessageRepository) InsertMessage(message models.Message) (id int64, err error) {
	err = l.db.QueryRow(
		"INSERT INTO Messages (Uuid, DateTime, Content) VALUES (@Uuid, @DateTime, @Content); SELECT SCOPE_IDENTITY();",
		sql.Named("Uuid", message.Uuid),
		sql.Named("DateTime", message.DateTime),
		sql.Named("Content", message.Content),
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (l *LocalMessageRepository) InsertMessagesBatch(messages []models.Message) error {
	// insert with db transaction
	dbTx, err := l.db.Begin()
	if err != nil {
		return err
	}

	statement, err := dbTx.Prepare("INSERT INTO messages (uuid, content, DateTime) VALUES (@p1, @p2, @p3)")
	if err != nil {
		dbTx.Rollback()
		return err
	}
	defer statement.Close()

	for _, msg := range messages {
		if _, err := statement.Exec(sql.Named("p1", msg.Uuid), sql.Named("p2", msg.Content), sql.Named("p3", msg.DateTime)); err != nil {
			dbTx.Rollback()
			return err
		}
	}

	return dbTx.Commit()
}

func (l *LocalMessageRepository) GetMessageByUuid(uuid string) (msg models.Message, err error) {
	err = l.db.QueryRow(
		"SELECT ID, Uuid, DateTime, Content FROM Messages WHERE Uuid = @Uuid",
		sql.Named("Uuid", uuid),
	).Scan(&msg.ID, &msg.Uuid, &msg.DateTime, &msg.Content)

	if err != nil {
		if err == sql.ErrNoRows {
			return msg, nil
		}
		return msg, err
	}

	return msg, nil
}

func (l *LocalMessageRepository) GetMessagesByUuidPart(uuidPart string) ([]models.Message, error) {
	var messages []models.Message

	rows, err := l.db.Query("SELECT ID, Uuid, DateTime, Content FROM Messages WHERE Uuid LIKE @UuidPart", sql.Named("UuidPart", "%"+uuidPart+"%"))
	if err != nil {
		return messages, err
	}
	defer rows.Close()

	for rows.Next() {
		var msg models.Message
		err = rows.Scan(&msg.ID, &msg.Uuid, &msg.DateTime, &msg.Content)
		if err != nil {
			return messages, err
		}
		messages = append(messages, msg)
	}

	if err = rows.Err(); err != nil {
		return messages, err
	}

	return messages, nil
}

func (l *LocalMessageRepository) GetMessages() ([]models.Message, error) {
	panic("unimplemented")
}

func (l *LocalMessageRepository) DeleteMessages() error {
	panic("unimplemented")
}

func checkAndCreateMessagesTable(db *sql.DB) error {
	exists, err := tableExists(db, "Messages")

	if err != nil {
		return err
	}

	if !exists {
		return createMessagesTable(db)
	}

	return nil
}
func tableExists(db *sql.DB, tableName string) (bool, error) {
	var exists bool
	query := `
		SELECT CASE 
			WHEN OBJECT_ID(@tableName, 'U') IS NOT NULL THEN 1
			ELSE 0
		END`
	err := db.QueryRow(query, sql.Named("tableName", tableName)).Scan(&exists)

	if err != nil {
		return false, fmt.Errorf("error checking if table exists: %v", err)
	}

	return exists, nil
}

func createMessagesTable(db *sql.DB) error {
	query := `
		CREATE TABLE Messages (
			ID INT IDENTITY(1,1) PRIMARY KEY,
			Uuid NVARCHAR(36) NOT NULL,
			DateTime DATETIME NOT NULL,
			Content NVARCHAR(MAX) NOT NULL
		)`
	_, err := db.Exec(query)

	if err != nil {
		return fmt.Errorf("error creating table: %v", err)
	}

	return nil
}
