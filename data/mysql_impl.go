package data

import (
	"chat/core"
	"time"

	 "database/sql"
	 _ "github.com/go-sql-driver/mysql"
	"github.com/nu7hatch/gouuid"
)

const (
	QueryGetUserByUsername = "SELECT * FROM USER WHERE username = ?;"

	QueryGetMessages = "SELECT * FROM MESSAGE WHERE to_uuid = ? AND CREATED_AT >= ?;"

	QueryInsertMessage = "INSERT INTO MESSAGE (uuid, from_uuid, to_uuid, content)" +
		"(?, ?, ?, ?);"

	QueryInsertUser = "INSERT INTO USER (uuid, username) VALUES (?, ?);"
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
	_, err := msm.mysqlSession.Query(QueryInsertMessage, m.UUID, m.FromUUID, m.ToUUID, m.Content)
	if err != nil {
		return err
	}
	return nil
}

func (msm *MysqlStorageManager) GetMessages(from time.Time, userUUID string) ([]*core.Message, error) {
	return nil, nil
}

func (msm *MysqlStorageManager) GetUserByUsername(u string) (*core.User, error) {
	var uuid, username string
	var lastSeen, createdAt []uint8

	result := msm.mysqlSession.QueryRow(QueryGetUserByUsername, u)
	result.Scan(&uuid, &username, &lastSeen, &createdAt)

	user := core.User{
		UUID: uuid,
		Username: username,
	}

	return &user, nil
}

func (msm *MysqlStorageManager) InsertUser(u *core.User) (*core.User, error) {
	userUUID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	_, err = msm.mysqlSession.Query(QueryInsertUser, userUUID.String(), u.Username)
	if err != nil {
		return nil, err
	}

	result, err := msm.GetUserByUsername(u.Username)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (msm *MysqlStorageManager) UpdateUser(u *core.User) error {
	return nil
}