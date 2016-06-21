package authentication

import (
	"chat/data"
	"log"
)

type UsernameAuthenticator struct {
	storageManager data.StorageManager
}

func NewUserNameAuthenticator(storageManager data.StorageManager) Authenticator {
	return &UsernameAuthenticator{storageManager}
}

func (a *UsernameAuthenticator) Authenticate(username interface{}) bool {
	user, err := a.storageManager.GetUserByUsername(username.(string))
	if err != nil {
		log.Println(err)
	}
	if user.UUID == "" {
		return false
	}
	return true
}