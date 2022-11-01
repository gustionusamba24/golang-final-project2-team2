package photo_resources

type PhotoCreateRequest struct {
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption" validate:"required"`
	PhotoUrl string `json:"photo_url" validate:"required"`
}

//
//type UserLoginRequest struct {
//	Email    string `json:"email" validate:"required,email"`
//	Password string `json:"password" validate:"required,min=6"`
//}
//
//type UserUpdateRequest struct {
//	Username string `json:"username" validate:"required"`
//	Email    string `json:"email" validate:"required,email"`
//}
