package user_resources

import "time"

type UserRegisterResponse struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int8   `json:"age"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type UserUpdateResponse struct {
	Id        int64     `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Age       int8      `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}
