package users

type CreateUserRequest struct {
	Username         string `json:"username"`
	Password         string `json:"password"`
	Role             string `json:"role"`
	Firstname        string `json:"firstname"`
	Lastname         string `json:"lastname"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	Department       string `json:"department"`
	ProfileImagePath string
	Files            []string
}
