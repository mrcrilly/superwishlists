package login

// Login is a model representing a user's credentials.
// A user, or Profile, can have multiple credentials such
// as username/password combinations, or an API key.
type Login struct {
	ID           string
	OwnerID      string
	Username     string
	PasswordHash string
	Enabled      bool
}
