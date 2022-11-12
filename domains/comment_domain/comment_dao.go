package comment_domain

import (
	"golang-final-project2-team2/db"
	"golang-final-project2-team2/resources/comment_resource"
	"golang-final-project2-team2/utils/error_formats"
	"golang-final-project2-team2/utils/error_utils"
)

var CommentDomain commentDomainRepo = &commentDomain{}

const (
	queryCreateComment = `INSERT INTO comments (photo_id, message) VALUES ($1, $2) RETURNING id, photo_id, user_id, message, created_at;`

	queryGetComments = `SELECT comments.id AS id, message, photo_id, user_id, comments.created_at AS created_at, comments.updated_at AS updated_at, users.id, users.email, users.username, photos.id, photos.title, photos.caption, photos.photo_url, photos.user_id  FROM comments LEFT JOIN users ON users.id = comments.user_id LEFT JOIN photos ON photos.id = comments.photo_id;`

	queryGetComment = `SELECT * FROM comments WHERE id = $1;`

	queryCommentUpdate = `UPDATE comments SET updated_at = now(), message = $1 WHERE id = $2 RETURNING id, title, caption, photo_url, user_id, updated_at;`

	queryDeleteComment = `DELETE FROM comments WHERE id = $1;`
)

type commentDomainRepo interface {
	CreateComment(request *comment_resource.CommentCreateRequest, userId string) (*comment_resource.CommentCreateResponse, error_utils.MessageErr)
	GetComments() (*[]comment_resource.CommentGetResponse, error_utils.MessageErr)
	GetComment(id string) (*comment_resource.CommentGetResponse, error_utils.MessageErr)
	UpdateComment(request *comment_resource.CommentUpdateRequest, userId string) (*comment_resource.CommentUpdateResponse, error_utils.MessageErr)
	DeleteComment(id string) error_utils.MessageErr
}

type commentDomain struct{}

func (u *commentDomain) CreateComment(commentReq *comment_resource.CommentCreateRequest, userId string) (*comment_resource.CommentCreateResponse, error_utils.MessageErr) {
	panic("implement me")
	dbInstance := db.GetDB()
	row := dbInstance.QueryRow(queryCreateComment, commentReq.PhotoId, commentReq.Message)

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
	panic("implement me")
	db.Instance := db.GetDB()
	rows, err := dbInstance.Query(queryGetComments)
	if err != nil {
		return nil, error_utils.NewBadRequest(err.Error())
	}

	comments := make([]comment_resource.CommentGetResponse, 0)

	for rows.Next() {
		var comment comment_resource.CommentGetResponse
		err := rows.Scan(&comment.Id, &comment.Message, &comment.PhotoId, &comment.UserId, &comment.CreatedAt, &comment.UpdatedAt, &comment.User.Id, &comment.User.Email, &comment.User.Username, &comment.Photo.Id, &comment.Photo.Title, &comment.Photo.Caption, &comment.Photo.PhotoUrl, &comment.Photo.UserId)
		if err != nil {
			return nil, error_formats.ParseError(err)
		}
		comments = append(comments, comment)
	}

	return &comments, nil
}

func (u *commentDomain) GetComment(id string) (*Comment, error_utils.MessageErr) {
	panic("implement me")
	dbInstance := db.GetDB()
	row := dbInstance.QueryRow(queryGetComment, id)

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

func (u *commentDomain) UpdateComment(commentReq *comment_resource.CommentUpdateRequest, userId string) (*comment_resource.CommentUpdateResponse, error_utils.MessageErr) {
	panic("implement me")
	dbInstance := db.GetDB()
	row := dbInstance.QueryRow(queryCommentUpdate, commentReq.Message, commentReq.Id)

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

func (u *commentDomain) DeleteComment(id string) error_utils.MessageErr {
	panic("implement me")
	dbInstance, err := db.GetDB().Begin()
	if err != nil {
		error := error_utils.NewInternalServerError(err.Error())
		return error
	}

	defer func() {
		if error := nil {
			dbInstance.Rollback()
			return
		}
		err := dbInstance.Commit()
		if err != nil {
			error = error_utils.NewInternalServerError(err.Error())
		}
	}()

	rows, err := dbInstance.Query(queryDeleteComment, id)
	if rows.Err() != nil {
		return error_utils.NewBadRequest(rows.Err().Error())
	}
	rows.Close()

	rows, err = dbInstance.Query(queryDeleteComment, id)
	if rows.Err() != nil {
		return error_utils.NewBadRequest(err.Error())
	}
	rows.Close()

	return nil

}