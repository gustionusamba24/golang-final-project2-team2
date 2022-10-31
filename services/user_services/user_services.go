package user_services

import (
	"golang-final-project2-team2/domains/user_domain"
	"golang-final-project2-team2/resources/user_resources"
	"golang-final-project2-team2/utils/error_utils"
	"golang-final-project2-team2/utils/helpers"
)

var UserService userServiceRepo = &userService{}

type userServiceRepo interface {
	UserRegister(*user_resources.UserRegisterRequest) (*user_resources.UserRegisterResponse, error_utils.MessageErr)
	UserLogin(*user_resources.UserLoginRequest) (*user_resources.UserLoginResponse, error_utils.MessageErr)
	UserUpdate(string, *user_resources.UserUpdateRequest) (*user_resources.UserUpdateResponse, error_utils.MessageErr)
	UserDelete(string) error_utils.MessageErr
	//UpdateProduct(*product_domain.Product) (*product_domain.Product, error_utils.MessageErr)
	//GetProducts() ([]*product_domain.Product, error_utils.MessageErr)
	//DeleteProduct(int64) error_utils.MessageErr
}

type userService struct{}

func (u *userService) UserRegister(userReq *user_resources.UserRegisterRequest) (*user_resources.UserRegisterResponse, error_utils.MessageErr) {
	err := helpers.ValidateRequest(userReq)

	if err != nil {
		return nil, err
	}
	newPass, err := helpers.HashPass(userReq.Password)
	if err != nil {
		return nil, err
	}

	userReq.Password = *newPass

	user, err := user_domain.UserDomain.UserRegister(userReq)

	if err != nil {
		return nil, err
	}

	return &user_resources.UserRegisterResponse{
		Age:      user.Age,
		Email:    user.Email,
		Id:       user.Id,
		Username: user.Username,
	}, nil
}

func (u *userService) UserLogin(userReq *user_resources.UserLoginRequest) (*user_resources.UserLoginResponse, error_utils.MessageErr) {
	err := helpers.ValidateRequest(userReq)

	if err != nil {
		return nil, err
	}
	user, err := user_domain.UserDomain.UserLogin(userReq)

	if err != nil {
		return nil, err
	}

	if valid := helpers.ComparePass([]byte(user.Password), []byte(userReq.Password)); !valid {
		return nil, error_utils.NewBadRequest("invalid credential")
	}

	token, err := helpers.GenerateToken(user)

	if err != nil {
		return nil, err
	}

	return &user_resources.UserLoginResponse{
		Token: *token,
	}, nil
}

func (u *userService) UserUpdate(id string, userReq *user_resources.UserUpdateRequest) (*user_resources.UserUpdateResponse, error_utils.MessageErr) {
	err := helpers.ValidateRequest(userReq)

	if err != nil {
		return nil, err
	}
	user, err := user_domain.UserDomain.UserUpdate(id, userReq)

	if err != nil {
		return nil, err
	}

	return &user_resources.UserUpdateResponse{
		Id:        user.Id,
		Email:     user.Email,
		Username:  user.Username,
		Age:       user.Age,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (u *userService) UserDelete(id string) error_utils.MessageErr {
	err := user_domain.UserDomain.UserDelete(id)

	if err != nil {
		return err
	}

	return nil
}

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
