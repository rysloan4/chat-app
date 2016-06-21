package authentication

// Authenticator is an interface for authenticating a user
type Authenticator interface {
	Authenticate(i interface{}) bool
}
