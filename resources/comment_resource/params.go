package comment_resource

type CommentCreateRequest struct {
	Message string `json:"message" validate:"required"`
	PhotoId int64  `json:"photo_id" validate:"required"`
}

type CommentUpdateRequest struct {
	Message string `json:"message" validate:"required"`
}
