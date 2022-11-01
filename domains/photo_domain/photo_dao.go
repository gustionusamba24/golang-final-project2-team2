package photo_domain

import (
	"golang-final-project2-team2/db"
	"golang-final-project2-team2/resources/photo_resources"
	"golang-final-project2-team2/utils/error_formats"
	"golang-final-project2-team2/utils/error_utils"
)

var PhotoDomain photoDomainRepo = &photoDomain{}

const (
	queryCreatePhoto = `INSERT INTO photos ( title, caption, photo_url, user_id) 
	VALUES($1, $2, $3, $4) RETURNING id,title, caption, photo_url, user_id, created_at`

	queryGetPhotos = `
select photos.id as id, title, caption, photo_url, user_id, photos.created_at  as created_at, photos.updated_at as updated_at,
       users.email as email, users.username as username  from photos left join users on users.id = photos.user_id;
	`
	//queryUserLogin  = `SELECT * from users where email = $1`
	//queryUserUpdate = `UPDATE users set updated_at = now(), email = $1, username = $2 where id = $3 RETURNING id,username,email, password, age, created_at, updated_at`
	//queryUserDelete = `UPDATE users SET  deleted_at = now() where id = $1`
	//queryUserById   = `SELECT * from users where id = $1 and deleted_at is NULL`
)

type photoDomainRepo interface {
	CreatePhoto(*photo_resources.PhotoCreateRequest, string) (*photo_resources.PhotoCreateResponse, error_utils.MessageErr)
	GetPhotos() (*[]photo_resources.PhotosGetResponse, error_utils.MessageErr)
}

type photoDomain struct {
}

func (u *photoDomain) CreatePhoto(photoReq *photo_resources.PhotoCreateRequest, userId string) (*photo_resources.PhotoCreateResponse, error_utils.MessageErr) {
	dbInstance := db.GetDB()
	row := dbInstance.QueryRow(queryCreatePhoto, photoReq.Title, photoReq.Caption, photoReq.PhotoUrl, userId)
	if row.Err() != nil {
		return nil, error_utils.NewBadRequest(row.Err().Error())
	}

	var photo photo_resources.PhotoCreateResponse

	err := row.Scan(&photo.Id, &photo.Title, &photo.Caption, &photo.PhotoUrl, &photo.UserId, &photo.CreatedAt)

	if err != nil {
		return nil, error_formats.ParseError(err)
	}

	return &photo, nil
}

func (u *photoDomain) GetPhotos() (*[]photo_resources.PhotosGetResponse, error_utils.MessageErr) {
	dbInstance := db.GetDB()
	rows, err := dbInstance.Query(queryGetPhotos)
	if err != nil {
		return nil, error_utils.NewBadRequest(err.Error())
	}

	var photos []photo_resources.PhotosGetResponse

	for rows.Next() {
		var photo photo_resources.PhotosGetResponse
		var photoUser photo_resources.PhotosUserGetResponse
		err = rows.Scan(&photo.Id, &photo.Title, &photo.Caption, &photo.PhotoUrl, &photo.UserId, &photo.CreatedAt, &photo.UpdatedAt, &photoUser.Email, &photoUser.Username)
		photo.User = &photoUser
		if err != nil {
			return nil, error_formats.ParseError(err)
		}

		photos = append(photos, photo)
	}

	return &photos, nil
}

//
//func (u *userDomain) UserCheckIsExists(id int64) bool {
//	dbInstance := db.GetDB()
//	row := dbInstance.QueryRow(queryUserById, id)
//	var user User
//	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Age, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
//	if row.Err() != nil || err != nil {
//		return false
//	}
//	return true
//}
//
//func (u *userDomain) UserLogin(userReq *user_resources.UserLoginRequest) (*User, error_utils.MessageErr) {
//	dbInstance := db.GetDB()
//	row := dbInstance.QueryRow(queryUserLogin, userReq.Email)
//	if row.Err() != nil {
//		return nil, error_utils.NewBadRequest(row.Err().Error())
//	}
//
//	var user User
//
//	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Age, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
//
//	if err != nil {
//		return nil, error_utils.NewBadRequest(err.Error())
//	}
//	return &user, nil
//}
//
//func (u *userDomain) UserUpdate(id string, userReq *user_resources.UserUpdateRequest) (*User, error_utils.MessageErr) {
//	dbInstance := db.GetDB()
//	row := dbInstance.QueryRow(queryUserUpdate, userReq.Email, userReq.Username, id)
//	if row.Err() != nil {
//		return nil, error_utils.NewBadRequest(row.Err().Error())
//	}
//
//	var user User
//
//	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Age, &user.CreatedAt, &user.UpdatedAt)
//
//	if err != nil {
//		return nil, error_utils.NewBadRequest(err.Error())
//	}
//	return &user, nil
//}
//func (u *userDomain) UserDelete(id string) error_utils.MessageErr {
//	dbInstance := db.GetDB()
//	row := dbInstance.QueryRow(queryUserDelete, id)
//	if row.Err() != nil {
//		return error_utils.NewBadRequest(row.Err().Error())
//	}
//	return nil
//}

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
