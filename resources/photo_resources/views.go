package photo_resources

import "time"

type PhotoCreateResponse struct {
	Id        int64     `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserId    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

//
//type UserLoginResponse struct {
//	Token string `json:"token"`
//}
//
//type UserUpdateResponse struct {
//	Id        int64     `json:"id"`
//	Email     string    `json:"email"`
//	Username  string    `json:"username"`
//	Age       int8      `json:"age"`
//	UpdatedAt time.Time `json:"updated_at"`
//}
