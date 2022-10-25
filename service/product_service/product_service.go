package product_service

import (
	"products/domain/product_domain"
	"products/utils/error_utils"
)

var ProductService productServiceRepo = &productService{}

type productServiceRepo interface {
	CreateProduct(*product_domain.Product) (*product_domain.Product, error_utils.MessageErr)
	UpdateProduct(*product_domain.Product) (*product_domain.Product, error_utils.MessageErr)
	GetProducts() ([]*product_domain.Product, error_utils.MessageErr)
	DeleteProduct(int64) error_utils.MessageErr
}

type productService struct{}

func (p *productService) CreateProduct(productReq *product_domain.Product) (*product_domain.Product, error_utils.MessageErr) {
	err := productReq.Validate()

	if err != nil {
		return nil, err
	}

	product, err := product_domain.ProductDomain.CreateProduct(productReq)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *productService) UpdateProduct(productReq *product_domain.Product) (*product_domain.Product, error_utils.MessageErr) {
	err := productReq.Validate()

	if err != nil {
		return nil, err
	}

	product, err := product_domain.ProductDomain.UpdateProduct(productReq)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *productService) GetProducts() ([]*product_domain.Product, error_utils.MessageErr) {

	products, err := product_domain.ProductDomain.GetProducts()

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *productService) DeleteProduct(productId int64) error_utils.MessageErr {

	err := product_domain.ProductDomain.DeleteProduct(productId)

	if err != nil {
		return err
	}

	return nil
}
