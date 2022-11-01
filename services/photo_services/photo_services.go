package photo_services

import (
	"golang-final-project2-team2/domains/photo_domain"
	"golang-final-project2-team2/resources/photo_resources"
	"golang-final-project2-team2/utils/error_utils"
	"golang-final-project2-team2/utils/helpers"
)

var PhotoService photoServiceRepo = &photoService{}

type photoServiceRepo interface {
	CreatePhoto(*photo_resources.PhotoCreateRequest, string) (*photo_resources.PhotoCreateResponse, error_utils.MessageErr)
	GetPhotos() (*[]photo_resources.PhotosGetResponse, error_utils.MessageErr)
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

//func (u *userService) UserLogin(userReq *user_resources.UserLoginRequest) (*user_resources.UserLoginResponse, error_utils.MessageErr) {
//	err := helpers.ValidateRequest(userReq)
//
//	if err != nil {
//		return nil, err
//	}
//	user, err := user_domain.UserDomain.UserLogin(userReq)
//
//	if err != nil {
//		return nil, err
//	}
//
//	if valid := helpers.ComparePass([]byte(user.Password), []byte(userReq.Password)); !valid {
//		return nil, error_utils.NewBadRequest("invalid credential")
//	}
//
//	token, err := helpers.GenerateToken(user)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return &user_resources.UserLoginResponse{
//		Token: *token,
//	}, nil
//}
//
//func (u *userService) UserUpdate(id string, userReq *user_resources.UserUpdateRequest) (*user_resources.UserUpdateResponse, error_utils.MessageErr) {
//	err := helpers.ValidateRequest(userReq)
//
//	if err != nil {
//		return nil, err
//	}
//	user, err := user_domain.UserDomain.UserUpdate(id, userReq)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return &user_resources.UserUpdateResponse{
//		Id:        user.Id,
//		Email:     user.Email,
//		Username:  user.Username,
//		Age:       user.Age,
//		UpdatedAt: user.UpdatedAt,
//	}, nil
//}
//
//func (u *userService) UserDelete(id string) error_utils.MessageErr {
//	err := user_domain.UserDomain.UserDelete(id)
//
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

//func (p *productService) UpdateProduct(productReq *product_domain.Product) (*product_domain.Product, error_utils.MessageErr) {
//	err := productReq.Validate()
//
//	if err != nil {
//		return nil, err
//	}
//
//	product, err := product_domain.ProductDomain.UpdateProduct(productReq)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return product, nil
//}
//
//func (p *productService) GetProducts() ([]*product_domain.Product, error_utils.MessageErr) {
//
//	products, err := product_domain.ProductDomain.GetProducts()
//
//	if err != nil {
//		return nil, err
//	}
//
//	return products, nil
//}
//
//func (p *productService) DeleteProduct(productId int64) error_utils.MessageErr {
//
//	err := product_domain.ProductDomain.DeleteProduct(productId)
//
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
