package data

import (
	"chat/core"
	"time"

	 "database/sql"
	 _ "github.com/go-sql-driver/mysql"
	_"encoding/json"
)

const (
	QueryGetUser = "SELECT * FROM USER WHERE UUID = ?;"
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
	return nil
}

func (msm *MysqlStorageManager) GetMessages(from time.Time, until time.Time, userUUID string) ([]*core.Message, error) {
	return nil, nil
}

func (msm *MysqlStorageManager) GetUser(u string) (*core.User, error) {
	var uuid, username string
	var lastSeen, createdAt []uint8

	result := msm.mysqlSession.QueryRow(QueryGetUser, u)
	result.Scan(&uuid, &username, &lastSeen, &createdAt)

	user := core.User{
		UUID: uuid,
		Username: username,
	}

	return &user, nil
}

func (msm *MysqlStorageManager) InsertUser(u *core.User) error {
	return nil
}

func (msm *MysqlStorageManager) UpdateUser(u *core.User) error {
	return nil
}