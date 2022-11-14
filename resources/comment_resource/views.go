package comment_resource

import "time"

type CommentCreateResponse struct {
	Id        int64     `json:"id"`
	Message   string    `json:"message"`
	PhotoId   int64     `json:"photo_id"`
	UserId    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentUserGetResponse struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type CommentGetPhotoResponse struct {
	Id       int64  `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserId   int64  `json:"user_id"`
}

type CommentGetResponse struct {
	Id        int64                    `json:"id"`
	Message   string                   `json:"message"`
	PhotoId   int64                    `json:"photo_id"`
	UserId    int64                    `json:"user_id"`
	CreatedAt time.Time                `json:"created_at"`
	UpdatedAt time.Time                `json:"updated_at"`
	User      *CommentUserGetResponse  `json:"user"`
	Photo     *CommentGetPhotoResponse `json:"photo"`
}

type CommentUpdateResponse struct {
	Id        int64     `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserId    int64     `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}
