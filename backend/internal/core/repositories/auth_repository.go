package repositories

type AuthRepository interface {
	AuthenticateUser(username string, password string) (string, error)
	//DeleteJWTToken(token string) error
}
