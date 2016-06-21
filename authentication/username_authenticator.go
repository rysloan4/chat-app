package authentication

import (
	"chat/data"
	"log"
)

// UsernameAuthenticator is an authenticator that authenticates on chat username
type UsernameAuthenticator struct {
	storageManager data.StorageManager
}

// NewUserNameAuthenticator returns a new authenticator
func NewUserNameAuthenticator(storageManager data.StorageManager) Authenticator {
	return &UsernameAuthenticator{storageManager}
}

// Authenticate authenticates a user
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
