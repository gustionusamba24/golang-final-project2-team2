package UserDomain

import (
	"golang-final-project2-team2/db"
	"golang-final-project2-team2/utils/ErrorFormats"
	"golang-final-project2-team2/utils/ErrorUtils"
	"log"
)

var UserDomain userDomainRepo = &userDomain{}

const (
	queryCreateUser = `INSERT INTO users ( username, email, password, age) 
	VALUES($1, $2, $3, $4) RETURNING id, username, email, age`
	//queryGetProducts   = `SELECT id, name, price, stock, created_at from products ORDER BY id ASC`
	//queryUpdateProduct = `
	//	UPDATE products
	//	SET name = $2,
	//	price = $3,
	//	stock = $4
	//	WHERE id = $1
	//	RETURNING id, name, price, stock, created_at`
	//queryDeleteProduct = `DELETE from products WHERE id = $1`
)

type userDomainRepo interface {
	CreateUser(*User) (*User, ErrorUtils.MessageErr)
	//UpdateProduct(*Product) (*Product, ErrorUtils.MessageErr)
	//GetProducts() ([]*Product, ErrorUtils.MessageErr)
	//DeleteProduct(int64) ErrorUtils.MessageErr
	Close()
}

type userDomain struct {
}

func (u *userDomain) CreateUser(userReq *User) (*User, ErrorUtils.MessageErr) {
	dbInstance := db.GetDB()
	row := dbInstance.QueryRow(queryCreateUser, userReq.Username, userReq.Email, userReq.Password, userReq.Age)
	if row.Err() != nil {
		return nil, ErrorUtils.NewBadRequest(row.Err().Error())
	}

	var user User

	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Age)

	if err != nil {
		return nil, ErrorFormats.ParseError(err)
	}

	return &user, nil
}

//func (p *productDomain) GetProducts() ([]*Product, ErrorUtils.MessageErr) {
//	db := db.GetDB()
//
//	rows, err := db.Query(queryGetProducts)
//
//	if err != nil {
//		return nil, ErrorFormats.ParseError(err)
//	}
//
//	var products []*Product
//	for rows.Next() {
//		var product Product
//		err = rows.Scan(&product.Id, &product.Name, &product.Price, &product.Stock, &product.CreatedAt)
//
//		if err != nil {
//			return nil, ErrorFormats.ParseError(err)
//		}
//
//		products = append(products, &product)
//	}
//
//	return products, nil
//}
//
//func (p *productDomain) UpdateProduct(productReq *Product) (*Product, ErrorUtils.MessageErr) {
//	db := db.GetDB()
//
//	row := db.QueryRow(queryUpdateProduct, productReq.Id, productReq.Name, productReq.Price, productReq.Stock)
//
//	var product Product
//
//	err := row.Scan(&product.Id, &product.Name, &product.Price, &product.Stock, &product.CreatedAt)
//
//	if err != nil {
//		return nil, ErrorFormats.ParseError(err)
//	}
//
//	return &product, nil
//}
//
//func (p *productDomain) DeleteProduct(productId int64) ErrorUtils.MessageErr {
//	db := db.GetDB()
//
//	_, err := db.Exec(queryDeleteProduct, productId)
//
//	if err != nil {
//		return ErrorFormats.ParseError(err)
//	}
//
//	return nil
//}

func (u *userDomain) Close() {
	err := db.GetDB().Close()
	if err != nil {
		log.Fatal(err.Error())
	}
}
