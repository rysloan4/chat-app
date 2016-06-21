package data

import (
	"chat/core"
	"time"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nu7hatch/gouuid"
	"log"
)

const (
	QueryGetUserByUsername = "SELECT * FROM USER WHERE username = ?;"

	QueryGetMessages = "SELECT * FROM MESSAGE WHERE to_username = ? AND created_at >= ?;"

	QueryInsertMessage = "INSERT INTO MESSAGE (uuid, from_username, to_username, content, created_at)" +
		" VALUES (?, ?, ?, ?, ?);"

	QueryInsertUser = "INSERT INTO USER (uuid, username, last_seen, created_at) VALUES (?, ?, ?, ?);"

	QueryUpdateUserLastSeen = "UPDATE USER SET last_seen = ? WHERE username = ?"
)

type MysqlStorageManager struct {
	mysqlSession *sql.DB
}

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

func (msm *MysqlStorageManager) CleanUp() {}

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

func (msm *MysqlStorageManager) UpdateUserLastSeen(username string, lastSeen time.Time) error {
	_, err := msm.mysqlSession.Query(QueryUpdateUserLastSeen, lastSeen, username)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
