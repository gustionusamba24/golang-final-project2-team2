package UserService

import (
	"golang-final-project2-team2/domains/UserDomain"
	"golang-final-project2-team2/utils/ErrorUtils"
	"golang-final-project2-team2/utils/Helpers"
)

var UserService userServiceRepo = &userService{}

type userServiceRepo interface {
	CreateUser(*UserDomain.User) (*UserDomain.UserCreateResponse, ErrorUtils.MessageErr)
	//UpdateProduct(*product_domain.Product) (*product_domain.Product, ErrorUtils.MessageErr)
	//GetProducts() ([]*product_domain.Product, ErrorUtils.MessageErr)
	//DeleteProduct(int64) ErrorUtils.MessageErr
}

type userService struct{}

func (u *userService) CreateUser(userReq *UserDomain.User) (*UserDomain.UserCreateResponse, ErrorUtils.MessageErr) {
	err := userReq.Validate()

	if err != nil {
		return nil, err
	}
	newPass, err := Helpers.HashPass(userReq.Password)
	if err != nil {
		return nil, err
	}

	userReq.Password = *newPass

	user, err := UserDomain.UserDomain.CreateUser(userReq)

	if err != nil {
		return nil, err
	}

	return &UserDomain.UserCreateResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
	}, nil
}

//func (p *productService) UpdateProduct(productReq *product_domain.Product) (*product_domain.Product, ErrorUtils.MessageErr) {
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
//func (p *productService) GetProducts() ([]*product_domain.Product, ErrorUtils.MessageErr) {
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
//func (p *productService) DeleteProduct(productId int64) ErrorUtils.MessageErr {
//
//	err := product_domain.ProductDomain.DeleteProduct(productId)
//
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
