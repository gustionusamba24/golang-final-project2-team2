package comment_domain

import (
	"golang-final-project2-team2/db"
	"golang-final-project2-team2/resources/comment_resource"
	"golang-final-project2-team2/utils/error_formats"
	"golang-final-project2-team2/utils/error_utils"
)

var CommentDomain commentDomainRepo = &commentDomain{}

const (
	queryCreateComment = `INSERT INTO comments (photo_id, message, user_id) VALUES ($1, $2, $3) RETURNING id, photo_id, user_id, message, created_at;`

	queryGetComments = `SELECT comments.id AS id, message, photo_id, user_id, comments.created_at AS created_at, comments.updated_at AS updated_at, users.id, users.email, users.username, photos.id, photos.title, photos.caption, photos.photo_url, photos.user_id  FROM comments LEFT JOIN users ON users.id = comments.user_id LEFT JOIN photos ON photos.id = comments.photo_id;`

	queryGetComment = `SELECT * FROM comments WHERE id = $1;`

	queryCommentUpdate = `UPDATE comments SET updated_at = now(), message = $1 WHERE id = $2 RETURNING id, title, caption, photo_url, user_id, updated_at;`

	queryDeleteComment = `DELETE FROM comments WHERE id = $1;`
)

type commentDomainRepo interface {
	CreateComment(*comment_resource.CommentCreateRequest, string) (*comment_resource.CommentCreateResponse, error_utils.MessageErr)
	GetComments() (*[]comment_resource.CommentGetResponse, error_utils.MessageErr)
	GetComment(string) (*Comment, error_utils.MessageErr)
	UpdateComment(*comment_resource.CommentUpdateRequest, string) (*comment_resource.CommentUpdateResponse, error_utils.MessageErr)
	DeleteComment(string) error_utils.MessageErr
}

type commentDomain struct{}

func (u *commentDomain) CreateComment(commentReq *comment_resource.CommentCreateRequest, userId string) (*comment_resource.CommentCreateResponse, error_utils.MessageErr) {
	dbInstance := db.GetDB()
	row := dbInstance.QueryRow(queryCreateComment, commentReq.PhotoId, commentReq.Message, userId)
	if row.Err() != nil {
		return nil, error_utils.NewBadRequest(row.Err().Error())
	}

	var comment comment_resource.CommentCreateResponse

	err := row.Scan(&comment.Id, &comment.PhotoId, &comment.UserId, &comment.Message, &comment.CreatedAt)

	if err != nil {
		return nil, error_formats.ParseError(err)
	}

	return &comment, nil
}

func (u *commentDomain) GetComments() (*[]comment_resource.CommentGetResponse, error_utils.MessageErr) {
	dbInstance := db.GetDB()
	rows, err := dbInstance.Query(queryGetComments)
	if err != nil {
		return nil, error_utils.NewBadRequest(err.Error())
	}

	comments := make([]comment_resource.CommentGetResponse, 0)

	for rows.Next() {
		var comment comment_resource.CommentGetResponse
		var commentUser comment_resource.CommentUserGetResponse
		var commentPhoto comment_resource.CommentGetPhotoResponse
		err = rows.Scan(&comment.Id, &comment.Message, &comment.PhotoId, &comment.UserId, &comment.CreatedAt, &comment.UpdatedAt, &commentUser.Id, &commentUser.Email, &commentUser.Username, &commentPhoto.Id, &commentPhoto.Title, &commentPhoto.Caption, &commentPhoto.PhotoUrl, &commentPhoto.UserId)
		comment.User = &commentUser
		comment.Photo = &commentPhoto
		if err != nil {
			return nil, error_formats.ParseError(err)
		}
		comments = append(comments, comment)
	}

	return &comments, nil
}

func (u *commentDomain) GetComment(commentId string) (*Comment, error_utils.MessageErr) {
	dbInstance := db.GetDB()
	row := dbInstance.QueryRow(queryGetComment, commentId)
	if row.Err() != nil {
		return nil, error_utils.NewBadRequest(row.Err().Error())
	}

	var comment Comment

	err := row.Scan(&comment.Id, &comment.Message, &comment.PhotoId, &comment.UserId, &comment.CreatedAt, &comment.UpdatedAt)

	if err != nil {
		return nil, error_formats.ParseError(err)
	}
	return &comment, nil
}

func (u *commentDomain) UpdateComment(request *comment_resource.CommentUpdateRequest, commentId string) (*comment_resource.CommentUpdateResponse, error_utils.MessageErr) {
	dbInstance := db.GetDB()
	row := dbInstance.QueryRow(queryCommentUpdate, request.Message, commentId)
	if row.Err() != nil {
		return nil, error_utils.NewBadRequest(row.Err().Error())
	}

	var comment comment_resource.CommentUpdateResponse

	err := row.Scan(&comment.Id, &comment.Title, &comment.Caption, &comment.PhotoUrl, &comment.UserId, &comment.UpdatedAt)

	if err != nil {
		return nil, error_formats.ParseError(err)
	}
	return &comment, nil
}

func (u *commentDomain) DeleteComment(commentId string) (error error_utils.MessageErr) {
	dbInstance, err := db.GetDB().Begin()
	if err != nil {
		error = error_utils.NewInternalServerError(err.Error())
		return
	}

	defer func() {
		if error != nil {
			dbInstance.Rollback()
			return
		}
		err := dbInstance.Commit()
		if err != nil {
			error = error_utils.NewInternalServerError(err.Error())
		}
	}()

	rows, err := dbInstance.Query(queryDeleteComment, commentId)
	if rows.Err() != nil {
		return error_utils.NewBadRequest(err.Error())
	}
	rows.Close()

	return nil
}
