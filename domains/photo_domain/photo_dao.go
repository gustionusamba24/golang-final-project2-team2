package photo_domain

import (
	"golang-final-project2-team2/db"
	"golang-final-project2-team2/resources/photo_resources"
	"golang-final-project2-team2/utils/error_formats"
	"golang-final-project2-team2/utils/error_utils"
)

var PhotoDomain photoDomainRepo = &photoDomain{}

const (
	queryCreatePhoto = `INSERT INTO photos ( title, caption, photo_url, user_id) 
	VALUES($1, $2, $3, $4) RETURNING id,title, caption, photo_url, user_id, created_at`

	queryGetPhotos = `
select photos.id as id, title, caption, photo_url, user_id, photos.created_at  as created_at, photos.updated_at as updated_at,
       users.email as email, users.username as username from photos left join users on users.id = photos.user_id;
	`

	queryGetPhoto = `select * from photos where id = $1`

	queryDeleteCommentsByPhotoId = `DELETE from comments where photo_id = $1`
	queryDeletePhoto             = `DELETE from photos where id = $1`

	queryPhotoUpdate = `UPDATE photos set updated_at = now(), title = $1, caption = $2, photo_url = $3 where id = $4 RETURNING id,title,caption, photo_url, user_id, updated_at`
)

type photoDomainRepo interface {
	CreatePhoto(*photo_resources.PhotoCreateRequest, string) (*photo_resources.PhotoCreateResponse, error_utils.MessageErr)
	GetPhotos() (*[]photo_resources.PhotosGetResponse, error_utils.MessageErr)
	GetPhoto(string) (*Photo, error_utils.MessageErr)
	UpdatePhoto(*photo_resources.PhotoUpdateRequest, string) (*photo_resources.PhotoUpdateResponse, error_utils.MessageErr)
	DeletePhoto(string) error_utils.MessageErr
}

type photoDomain struct {
}

func (u *photoDomain) CreatePhoto(photoReq *photo_resources.PhotoCreateRequest, userId string) (*photo_resources.PhotoCreateResponse, error_utils.MessageErr) {
	dbInstance := db.GetDB()
	row := dbInstance.QueryRow(queryCreatePhoto, photoReq.Title, photoReq.Caption, photoReq.PhotoUrl, userId)
	if row.Err() != nil {
		return nil, error_utils.NewBadRequest(row.Err().Error())
	}

	var photo photo_resources.PhotoCreateResponse

	err := row.Scan(&photo.Id, &photo.Title, &photo.Caption, &photo.PhotoUrl, &photo.UserId, &photo.CreatedAt)

	if err != nil {
		return nil, error_formats.ParseError(err)
	}

	return &photo, nil
}

func (u *photoDomain) GetPhotos() (*[]photo_resources.PhotosGetResponse, error_utils.MessageErr) {
	dbInstance := db.GetDB()
	rows, err := dbInstance.Query(queryGetPhotos)
	if err != nil {
		return nil, error_utils.NewBadRequest(err.Error())
	}

	photos := make([]photo_resources.PhotosGetResponse, 0)

	for rows.Next() {
		var photo photo_resources.PhotosGetResponse
		var photoUser photo_resources.PhotosUserGetResponse
		err = rows.Scan(&photo.Id, &photo.Title, &photo.Caption, &photo.PhotoUrl, &photo.UserId, &photo.CreatedAt, &photo.UpdatedAt, &photoUser.Email, &photoUser.Username)
		photo.User = &photoUser
		if err != nil {
			return nil, error_formats.ParseError(err)
		}

		photos = append(photos, photo)
	}

	return &photos, nil
}

func (u *photoDomain) GetPhoto(photoId string) (*Photo, error_utils.MessageErr) {
	dbInstance := db.GetDB()
	row := dbInstance.QueryRow(queryGetPhoto, photoId)
	if row.Err() != nil {
		return nil, error_utils.NewBadRequest(row.Err().Error())
	}

	var photo Photo

	err := row.Scan(&photo.Id, &photo.Title, &photo.Caption, &photo.PhotoUrl, &photo.UserId, &photo.CreatedAt, &photo.UpdatedAt)

	if err != nil {
		return nil, error_formats.ParseError(err)
	}
	return &photo, nil
}

func (u *photoDomain) UpdatePhoto(request *photo_resources.PhotoUpdateRequest, photoId string) (*photo_resources.PhotoUpdateResponse, error_utils.MessageErr) {
	dbInstance := db.GetDB()
	row := dbInstance.QueryRow(queryPhotoUpdate, request.Title, request.Caption, request.PhotoUrl, photoId)
	if row.Err() != nil {
		return nil, error_utils.NewBadRequest(row.Err().Error())
	}

	var photo photo_resources.PhotoUpdateResponse

	err := row.Scan(&photo.Id, &photo.Title, &photo.Caption, &photo.PhotoUrl, &photo.UserId, &photo.UpdatedAt)

	if err != nil {
		return nil, error_formats.ParseError(err)
	}
	return &photo, nil
}

func (u *photoDomain) DeletePhoto(photoId string) (error error_utils.MessageErr) {
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

	rows, err := dbInstance.Query(queryDeleteCommentsByPhotoId, photoId)
	if rows.Err() != nil {
		return error_utils.NewBadRequest(err.Error())
	}
	rows.Close()

	rows, err = dbInstance.Query(queryDeletePhoto, photoId)
	if rows.Err() != nil {
		return error_utils.NewBadRequest(err.Error())
	}
	rows.Close()

	return nil
}
