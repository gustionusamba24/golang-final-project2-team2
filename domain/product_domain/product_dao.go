package product_domain

import (
	"products/db"
	"products/utils/error_formats"
	"products/utils/error_utils"
)

var ProductDomain productDomainRepo = &productDomain{}

const (
	queryCreateProduct = `INSERT INTO products (name, price, stock) 
	VALUES($1, $2, $3) RETURNING id, name, price, stock, created_at`
	queryGetProducts   = `SELECT id, name, price, stock, created_at from products ORDER BY id ASC`
	queryUpdateProduct = `
		UPDATE products
		SET name = $2,
		price = $3,
		stock = $4
		WHERE id = $1
		RETURNING id, name, price, stock, created_at`
	queryDeleteProduct = `DELETE from products WHERE id = $1`
)

type productDomainRepo interface {
	CreateProduct(*Product) (*Product, error_utils.MessageErr)
	UpdateProduct(*Product) (*Product, error_utils.MessageErr)
	GetProducts() ([]*Product, error_utils.MessageErr)
	DeleteProduct(int64) error_utils.MessageErr
	Close()
}

type productDomain struct {
}

func (p *productDomain) CreateProduct(productReq *Product) (*Product, error_utils.MessageErr) {
	db := db.GetDB()

	row := db.QueryRow(queryCreateProduct, productReq.Name, productReq.Price, productReq.Stock)

	var product Product

	err := row.Scan(&product.Id, &product.Name, &product.Price, &product.Stock, &product.CreatedAt)

	if err != nil {
		return nil, error_formats.ParseError(err)
	}

	return &product, nil
}

func (p *productDomain) GetProducts() ([]*Product, error_utils.MessageErr) {
	db := db.GetDB()

	rows, err := db.Query(queryGetProducts)

	if err != nil {
		return nil, error_formats.ParseError(err)
	}

	var products []*Product
	for rows.Next() {
		var product Product
		err = rows.Scan(&product.Id, &product.Name, &product.Price, &product.Stock, &product.CreatedAt)

		if err != nil {
			return nil, error_formats.ParseError(err)
		}

		products = append(products, &product)
	}

	return products, nil
}

func (p *productDomain) UpdateProduct(productReq *Product) (*Product, error_utils.MessageErr) {
	db := db.GetDB()

	row := db.QueryRow(queryUpdateProduct, productReq.Id, productReq.Name, productReq.Price, productReq.Stock)

	var product Product

	err := row.Scan(&product.Id, &product.Name, &product.Price, &product.Stock, &product.CreatedAt)

	if err != nil {
		return nil, error_formats.ParseError(err)
	}

	return &product, nil
}

func (p *productDomain) DeleteProduct(productId int64) error_utils.MessageErr {
	db := db.GetDB()

	_, err := db.Exec(queryDeleteProduct, productId)

	if err != nil {
		return error_formats.ParseError(err)
	}

	return nil
}

func (u *productDomain) Close() {
	db.GetDB().Close()
}
