package user_domain

import (
	"golang-final-project2-team2/db"
	"golang-final-project2-team2/resources/user_resources"
	"golang-final-project2-team2/utils/error_formats"
	"golang-final-project2-team2/utils/error_utils"
)

var UserDomain userDomainRepo = &userDomain{}

const (
	queryCreateUser = `INSERT INTO users ( username, email, password, age) 
	VALUES($1, $2, $3, $4) RETURNING id, username, email, age`
	queryUserLogin  = `SELECT * from users where email = $1`
	queryUserUpdate = `UPDATE users set updated_at = now(), email = $1, username = $2 where id = $3 RETURNING id,username,email, password, age, created_at, updated_at`
	queryUserDelete = `UPDATE users SET  deleted_at = now() where id = $1`
	queryUserById   = `SELECT * from users where id = $1 and deleted_at is NULL`
)

type userDomainRepo interface {
	UserRegister(*user_resources.UserRegisterRequest) (*User, error_utils.MessageErr)
	UserLogin(*user_resources.UserLoginRequest) (*User, error_utils.MessageErr)
	UserUpdate(string, *user_resources.UserUpdateRequest) (*User, error_utils.MessageErr)
	UserDelete(string) error_utils.MessageErr
	UserCheckIsExists(int64) bool
}

type userDomain struct {
}

func (u *userDomain) UserRegister(userReq *user_resources.UserRegisterRequest) (*User, error_utils.MessageErr) {
	dbInstance := db.GetDB()
	row := dbInstance.QueryRow(queryCreateUser, userReq.Username, userReq.Email, userReq.Password, userReq.Age)
	if row.Err() != nil {
		return nil, error_utils.NewBadRequest(row.Err().Error())
	}

	var user User

	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Age)

	if err != nil {
		return nil, error_formats.ParseError(err)
	}

	return &user, nil
}

func (u *userDomain) UserCheckIsExists(id int64) bool {
	dbInstance := db.GetDB()
	row := dbInstance.QueryRow(queryUserById, id)
	var user User
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Age, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if row.Err() != nil || err != nil {
		return false
	}
	return true
}

func (u *userDomain) UserLogin(userReq *user_resources.UserLoginRequest) (*User, error_utils.MessageErr) {
	dbInstance := db.GetDB()
	row := dbInstance.QueryRow(queryUserLogin, userReq.Email)
	if row.Err() != nil {
		return nil, error_utils.NewBadRequest(row.Err().Error())
	}

	var user User

	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Age, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)

	if err != nil {
		return nil, error_utils.NewBadRequest(err.Error())
	}
	return &user, nil
}

func (u *userDomain) UserUpdate(id string, userReq *user_resources.UserUpdateRequest) (*User, error_utils.MessageErr) {
	dbInstance := db.GetDB()
	row := dbInstance.QueryRow(queryUserUpdate, userReq.Email, userReq.Username, id)
	if row.Err() != nil {
		return nil, error_utils.NewBadRequest(row.Err().Error())
	}

	var user User

	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Age, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, error_utils.NewBadRequest(err.Error())
	}
	return &user, nil
}
func (u *userDomain) UserDelete(id string) error_utils.MessageErr {
	dbInstance := db.GetDB()
	row := dbInstance.QueryRow(queryUserDelete, id)
	if row.Err() != nil {
		return error_utils.NewBadRequest(row.Err().Error())
	}
	return nil
}

//func (p *productDomain) GetProducts() ([]*Product, error_utils.MessageErr) {
//	db := db.GetDB()
//
//	rows, err := db.Query(queryGetProducts)
//
//	if err != nil {
//		return nil, error_formats.ParseError(err)
//	}
//
//	var products []*Product
//	for rows.Next() {
//		var product Product
//		err = rows.Scan(&product.Id, &product.Name, &product.Price, &product.Stock, &product.CreatedAt)
//
//		if err != nil {
//			return nil, error_formats.ParseError(err)
//		}
//
//		products = append(products, &product)
//	}
//
//	return products, nil
//}
//
//func (p *productDomain) UpdateProduct(productReq *Product) (*Product, error_utils.MessageErr) {
//	db := db.GetDB()
//
//	row := db.QueryRow(queryUpdateProduct, productReq.Id, productReq.Name, productReq.Price, productReq.Stock)
//
//	var product Product
//
//	err := row.Scan(&product.Id, &product.Name, &product.Price, &product.Stock, &product.CreatedAt)
//
//	if err != nil {
//		return nil, error_formats.ParseError(err)
//	}
//
//	return &product, nil
//}
//
//func (p *productDomain) DeleteProduct(productId int64) error_utils.MessageErr {
//	db := db.GetDB()
//
//	_, err := db.Exec(queryDeleteProduct, productId)
//
//	if err != nil {
//		return error_formats.ParseError(err)
//	}
//
//	return nil
//}
