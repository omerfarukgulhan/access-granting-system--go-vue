package requests

type SignInCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
