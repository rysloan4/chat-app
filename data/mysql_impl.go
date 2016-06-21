package data

import (
	"chat/core"
	"time"

	"database/sql"
	// to pass linter
	_ "github.com/go-sql-driver/mysql"
	"github.com/nu7hatch/gouuid"
	"log"
)

const (
	// QueryGetUserByUsername to get user by username
	QueryGetUserByUsername = "SELECT * FROM USER WHERE username = ?;"

	// QueryGetMessages to get messages created after a certain time
	QueryGetMessages = "SELECT * FROM MESSAGE WHERE to_username = ? AND created_at >= ?;"

	// QueryInsertMessage to insert a message
	QueryInsertMessage = "INSERT INTO MESSAGE (uuid, from_username, to_username, content, created_at)" +
		" VALUES (?, ?, ?, ?, ?);"

	// QueryInsertUser to insert a user
	QueryInsertUser = "INSERT INTO USER (uuid, username, last_seen, created_at) VALUES (?, ?, ?, ?);"

	// QueryUpdateUserLastSeen to update when a user was last seen
	QueryUpdateUserLastSeen = "UPDATE USER SET last_seen = ? WHERE username = ?"
)

// MysqlStorageManager is a struct for a mysql storage manageer
type MysqlStorageManager struct {
	mysqlSession *sql.DB
}

// NewMysqlStorageManager returns a new mysql storage manager
func NewMysqlStorageManager(connectionString string) (StorageManager, error) {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &MysqlStorageManager{db}, nil
}

// CleanUp cleans up db connection
func (msm *MysqlStorageManager) CleanUp() {}

// InsertMessage inserts a message
func (msm *MysqlStorageManager) InsertMessage(m *core.Message) error {
	msgUUID, err := uuid.NewV4()
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = msm.mysqlSession.Query(QueryInsertMessage, msgUUID.String(), m.FromUsername, m.ToUsername, m.Content, time.Now())
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// GetMessages gets messages
func (msm *MysqlStorageManager) GetMessages(from time.Time, username string) ([]*core.Message, error) {
	var uuid, toUsername, fromUsername, content string
	var createdAt time.Time

	messages := []*core.Message{}

	rows, _ := msm.mysqlSession.Query(QueryGetMessages, username, from)

	for rows.Next() {
		rows.Scan(&uuid, &fromUsername, &toUsername, &content, &createdAt)
		message := core.Message{
			UUID:         uuid,
			FromUsername: fromUsername,
			ToUsername:   toUsername,
			Content:      content,
			CreatedAt:    createdAt,
		}
		messages = append(messages, &message)
	}

	return messages, nil
}

// GetUserByUsername gets a user by username
func (msm *MysqlStorageManager) GetUserByUsername(u string) (*core.User, error) {
	var uuid, username string
	var lastSeen, createdAt time.Time

	result := msm.mysqlSession.QueryRow(QueryGetUserByUsername, u)
	result.Scan(&uuid, &username, &lastSeen, &createdAt)

	user := core.User{
		UUID:      uuid,
		Username:  username,
		CreatedAt: createdAt,
		LastSeen:  lastSeen,
	}

	return &user, nil
}

// InsertUser inserts a user
func (msm *MysqlStorageManager) InsertUser(u *core.User) (*core.User, error) {
	userUUID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	_, err = msm.mysqlSession.Query(QueryInsertUser, userUUID.String(), u.Username, time.Now(), time.Now())
	if err != nil {
		return nil, err
	}

	result, err := msm.GetUserByUsername(u.Username)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateUserLastSeen updates when a user was last seen
func (msm *MysqlStorageManager) UpdateUserLastSeen(username string, lastSeen time.Time) error {
	_, err := msm.mysqlSession.Query(QueryUpdateUserLastSeen, lastSeen, username)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
