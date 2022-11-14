package social_media_services

import (
	"golang-final-project2-team2/domains/social_media_domain"
	"golang-final-project2-team2/resources/social_media_resource"
	"golang-final-project2-team2/utils/error_formats"
	"golang-final-project2-team2/utils/error_utils"
	"golang-final-project2-team2/utils/helpers"
	"strconv"
)

var SocialMediaService socialMediaServiceRepo = &socialMediaService{}

type socialMediaServiceRepo interface {
	CreateSocialMedia(*social_media_resource.SocialMediaCreateRequest, string) (*social_media_resource.SocialMediaCreateResponse, error_utils.MessageErr)
	GetSocialMedias() (*[]social_media_resource.SocialMediaGetResponse, error_utils.MessageErr)
	UpdateSocialMedia(*social_media_resource.SocialMediaUpdateRequest, string, string) (*social_media_resource.SocialMediaUpdateResponse, error_utils.MessageErr)
	DeleteSocialMedia(string, string) error_utils.MessageErr
}

type socialMediaService struct{}

func (u *socialMediaService) CreateSocialMedia(socialMediaReq *social_media_resource.SocialMediaCreateRequest, userId string) (*social_media_resource.SocialMediaCreateResponse, error_utils.MessageErr) {
	err := helpers.ValidateRequest(socialMediaReq)

	if err != nil {
		return nil, err
	}

	socialMedia, err := social_media_domain.SocialMediaDomain.CreateSocialMedia(socialMediaReq, userId)

	if err != nil {
		return nil, err
	}

	return socialMedia, nil
}

func (u *socialMediaService) GetSocialMedias() (*[]social_media_resource.SocialMediaGetResponse, error_utils.MessageErr) {
	socialMedias, err := social_media_domain.SocialMediaDomain.GetSocialMedias()

	if err != nil {
		return nil, err
	}

	return socialMedias, nil
}

func (u *socialMediaService) UpdateSocialMedia(request *social_media_resource.SocialMediaUpdateRequest, userId string, socialMediaId string) (*social_media_resource.SocialMediaUpdateResponse, error_utils.MessageErr) {
	socialMedia, err := social_media_domain.SocialMediaDomain.GetSocialMedia(socialMediaId)
	if err != nil {
		return nil, err
	}

	if strconv.FormatInt(socialMedia.UserId, 10) != userId {
		return nil, error_formats.NoAuthorization()
	}

	updatedSocialMedia, err := social_media_domain.SocialMediaDomain.UpdateSocialMedia(request, socialMediaId)

	if err != nil {
		return nil, err
	}

	return updatedSocialMedia, nil
}

func (u *socialMediaService) DeleteSocialMedia(userId string, socialMediaId string) error_utils.MessageErr {
	socialMedia, err := social_media_domain.SocialMediaDomain.GetSocialMedia(socialMediaId)
	if err != nil {
		return err
	}

	if strconv.FormatInt(socialMedia.UserId, 10) != userId {
		return error_formats.NoAuthorization()
	}

	err = social_media_domain.SocialMediaDomain.DeleteSocialMedia(socialMediaId)

	if err != nil {
		return err
	}

	return nil
}
