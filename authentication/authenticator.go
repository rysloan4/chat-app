package authentication

type Authenticator interface {
	Authenticate(interface{}) bool
}
