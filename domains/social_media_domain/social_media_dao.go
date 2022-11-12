package social_media_domain

import (
	"golang-final-project2-team2/db"
	"golang-final-project2-team2/resources/social_media_resource"
	"golang-final-project2-team2/utils/error_utils"
)

var SocialMediaDomain socialMediaDomainRepo = &socialMediaDomain{}

const (
	queryCreateSocialMedia = `INSERT INTO social_medias (name, social_media_url) VALUES ($1, $2) RETURNING id, name, social_media_url, user_id, created_at;`

	queryGetSocialMedias = `SELECT social_medias.id AS id, name, social_media_url, user_id, social_medias.created_at AS created_at, social_medias.updated_at AS updated_at, users.id, users.username, users.profile_image_url FROM social_medias LEFT JOIN users ON social_medias.user_id = users.id;`

	queryGetSocialMedia = `SELECT * FROM social_medias WHERE id = $1;`

	queryUpdateSocialMedia = `UPDATE social_medias SET name = $1, social_media_url = $2 WHERE id = $3 RETURNING id, name, social_media_url, user_id, updated_at;`

	queryDeleteSocialMedia = `DELETE FROM social_medias WHERE id = $1;`
)

type socialMediaDomainRepo interface {
	CreateSocialMedia(request *social_media_resource.SocialMediaCreateRequest, userId string) (*social_media_resource.SocialMediaCreateResponse, error_utils.MessageErr)
	GetSocialMedias() (*[]social_media_resource.SocialMediaGetResponse, error_utils.MessageErr)
	GetSocialMedia(id string) (*social_media_resource.SocialMediaGetResponse, error_utils.MessageErr)
	UpdateSocialMedia(request *social_media_resource.SocialMediaUpdateRequest, userId string) (*social_media_resource.SocialMediaUpdateResponse, error_utils.MessageErr)
	DeleteSocialMedia(id string) error_utils.MessageErr
}

type socialMediaDomain struct{}

func (u *socialMediaDomain) CreateSocialMedia(socialMediaReq *social_media_resource.SocialMediaCreateRequest, userId string) (*social_media_resource.SocialMediaCreateResponse, error_utils.MessageErr) {
	panic("implement me")
	dbInstance := db.GetDB()
	row := dbInstance.QueryRow(queryCreateSocialMedia, socialMediaReq.Name, socialMediaReq.SocialMediaUrl)
	if row.Err() != nil {
		return nil, error_utils.NewBadRequest(row.Err().Error())
	}

	var socialMedia social_media_resource.SocialMediaCreateResponse

	err := row.Scan(&socialMedia.Id, &socialMedia.Name, &socialMedia.SocialMediaUrl, &socialMedia.UserId, &socialMedia.CreatedAt)

	if err != nil {
		return nil, error_utils.NewBadRequest(err.Error())
	}

	return &socialMedia, nil
}

func (u *socialMediaDomain) GetSocialMedias() (*[]social_media_resource.SocialMediaGetResponse, error_utils.MessageErr) {
	panic("implement me")
	dbInstance := db.GetDB()
	rows, err := dbInstance.Query(queryGetSocialMedias)
	if err != nil {
		return nil, error_utils.NewBadRequest(err.Error())
	}

	var socialMedias []social_media_resource.SocialMediaGetResponse

	for rows.Next() {
		var socialMedia social_media_resource.SocialMediaGetResponse
		err := rows.Scan(&socialMedia.Id, &socialMedia.Name, &socialMedia.SocialMediaUrl, &socialMedia.UserId, &socialMedia.CreatedAt, &socialMedia.UpdatedAt, &socialMedia.User.Id, &socialMedia.User.Username, &socialMedia.User.ProfileImageUrl)
		if err != nil {
			return nil, error_utils.NewBadRequest(err.Error())
		}
		socialMedias = append(socialMedias, socialMedia)
	}

	return &socialMedias, nil
}

func (u *socialMediaDomain) GetSocialMedia(id string) (*social_media_resource.SocialMediaGetResponse, error_utils.MessageErr) {
	panic("implement me")
	dbInstance := db.GetDB()
	row := dbInstance.QueryRow(queryGetSocialMedia, id)
	if row.Err() != nil {
		return nil, error_utils.NewBadRequest(row.Err().Error())
	}

	var socialMedia social_media_resource.SocialMediaGetResponse

	err := row.Scan(&socialMedia.Id, &socialMedia.Name, &socialMedia.SocialMediaUrl, &socialMedia.UserId, &socialMedia.CreatedAt, &socialMedia.UpdatedAt, &socialMedia.User.Id, &socialMedia.User.Username, &socialMedia.User.ProfileImageUrl)

	if err != nil {
		return nil, error_utils.NewBadRequest(err.Error())
	}

	return &socialMedia, nil
}

func (u *socialMediaDomain) UpdateSocialMedia(socialMediaReq *social_media_resource.SocialMediaUpdateRequest, userId string) (*social_media_resource.SocialMediaUpdateResponse, error_utils.MessageErr) {
	panic("implement me")
	dbInstance := db.GetDB()
	row := dbInstance.QueryRow(queryUpdateSocialMedia, socialMediaReq.Name, socialMediaReq.SocialMediaUrl, socialMediaReq.Id)
	if row.Err() != nil {
		return nil, error_utils.NewBadRequest(row.Err().Error())
	}

	var socialMedia social_media_resource.SocialMediaUpdateResponse

	err := row.Scan(&socialMedia.Id, &socialMedia.Name, &socialMedia.SocialMediaUrl, &socialMedia.UserId, &socialMedia.UpdatedAt)

	if err != nil {
		return nil, error_utils.NewBadRequest(err.Error())
	}

	return &socialMedia, nil
}

func (u *socialMediaDomain) DeleteSocialMedia(id string) error_utils.MessageErr {
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

	rows, err := dbInstance.Query(queryDeleteSocialMedia, id)
	if rows.Err() != nil {
		return error_utils.NewBadRequest(rows.Err().Error())
	}
	rows.Close()

	rows, err = dbInstance.Query(queryDeleteSocialMedia, id)
	if rows.Err() != nil {
		return error_utils.NewBadRequest(err.Error())
	}
	rows.Close()

	return nil
}

