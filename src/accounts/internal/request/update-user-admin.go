package request

type UpdateUserRequest struct {
	Email    string  `json:"email"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Balance  float64 `json:"balance"`
	Admin    bool    `json:"admin"`
}
