package comment_domain

import "time"

type Comment struct {
	Id        int64      `json:"id"`
	UserId    int64      `json:"user_id"`
	PhotoId   string     `json:"photo_id"`
	Message   string     `json:"message"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
