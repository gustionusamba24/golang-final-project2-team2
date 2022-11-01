package photo_services

import (
	"golang-final-project2-team2/domains/photo_domain"
	"golang-final-project2-team2/resources/photo_resources"
	"golang-final-project2-team2/utils/error_formats"
	"golang-final-project2-team2/utils/error_utils"
	"golang-final-project2-team2/utils/helpers"
	"strconv"
)

var PhotoService photoServiceRepo = &photoService{}

type photoServiceRepo interface {
	CreatePhoto(*photo_resources.PhotoCreateRequest, string) (*photo_resources.PhotoCreateResponse, error_utils.MessageErr)
	GetPhotos() (*[]photo_resources.PhotosGetResponse, error_utils.MessageErr)
	UpdatePhoto(*photo_resources.PhotoUpdateRequest, string, string) (*photo_resources.PhotoUpdateResponse, error_utils.MessageErr)
	DeletePhoto(string, string) error_utils.MessageErr
	//UserLogin(*user_resources.UserLoginRequest) (*user_resources.UserLoginResponse, error_utils.MessageErr)
	//UserUpdate(string, *user_resources.UserUpdateRequest) (*user_resources.UserUpdateResponse, error_utils.MessageErr)
	//UserDelete(string) error_utils.MessageErr
	//UpdateProduct(*product_domain.Product) (*product_domain.Product, error_utils.MessageErr)
	//GetProducts() ([]*product_domain.Product, error_utils.MessageErr)
	//DeleteProduct(int64) error_utils.MessageErr
}

type photoService struct{}

func (u *photoService) CreatePhoto(photoReq *photo_resources.PhotoCreateRequest, userId string) (*photo_resources.PhotoCreateResponse, error_utils.MessageErr) {
	err := helpers.ValidateRequest(photoReq)

	if err != nil {
		return nil, err
	}

	photo, err := photo_domain.PhotoDomain.CreatePhoto(photoReq, userId)

	if err != nil {
		return nil, err
	}

	return photo, nil
}
func (u *photoService) GetPhotos() (*[]photo_resources.PhotosGetResponse, error_utils.MessageErr) {
	photos, err := photo_domain.PhotoDomain.GetPhotos()

	if err != nil {
		return nil, err
	}

	return photos, nil
}

func (u *photoService) UpdatePhoto(request *photo_resources.PhotoUpdateRequest, userId string, photoId string) (*photo_resources.PhotoUpdateResponse, error_utils.MessageErr) {
	photo, err := photo_domain.PhotoDomain.GetPhoto(photoId)
	if err != nil {
		return nil, err
	}
	if strconv.FormatInt(photo.UserId, 10) != userId {
		return nil, error_formats.NoAuthorization()
	}
	updatedPhoto, err := photo_domain.PhotoDomain.UpdatePhoto(request, photoId)

	if err != nil {
		return nil, err
	}

	return updatedPhoto, nil
}

func (u *photoService) DeletePhoto(userId string, photoId string) error_utils.MessageErr {
	photo, err := photo_domain.PhotoDomain.GetPhoto(photoId)
	if err != nil {
		return err
	}
	if strconv.FormatInt(photo.UserId, 10) != userId {
		return error_formats.NoAuthorization()
	}
	err = photo_domain.PhotoDomain.DeletePhoto(photoId)

	if err != nil {
		return err
	}
	return nil
}
