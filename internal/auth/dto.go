package auth

type CheckUserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
