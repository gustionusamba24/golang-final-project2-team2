package product_domain

import (
	"database/sql"
	"log"
	"products/db"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Fatalf("an error %s was not expected when opening a stub database connection", err)
	}
	return db, mock
}

func TestCreateProduct_Success(t *testing.T) {
	pgDB, mock := NewMock()
	_ = mock
	db.SetDB(pgDB)

	defer func() {
		ProductDomain.Close()
	}()

	query := regexp.QuoteMeta(`INSERT INTO products (name, price, stock) 
	VALUES($1, $2, $3) RETURNING id, name, price, stock, created_at`)

	p := &Product{
		Id:        1,
		Name:      "Product Test",
		Price:     10.10,
		Stock:     5,
		CreatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{
		"id",
		"name",
		"price",
		"stock",
		"created_at",
	}).AddRow(p.Id, p.Name, p.Price, p.Stock, p.CreatedAt)

	mock.ExpectQuery(query).WithArgs(p.Name, p.Price, p.Stock).WillReturnRows(rows)

	product, err := ProductDomain.CreateProduct(p)

	assert.NotNil(t, product)

	assert.Nil(t, err)
}

func TestCreateProduct_Error(t *testing.T) {
	pgDB, mock := NewMock()
	_ = mock
	db.SetDB(pgDB)

	defer func() {
		ProductDomain.Close()
	}()

	query := regexp.QuoteMeta(`INSERT INTO products (name, price, stock) 
	VALUES($1, $2, $3) RETURNING id, name, price, stock, created_at`)

	p := &Product{
		Id:        1,
		Name:      "Product Test",
		Price:     10.10,
		Stock:     5,
		CreatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{
		"id",
		"name",
		"price",
		"stock",
		"created_at",
	}).AddRow(p.Id, p.Name, p.Price, p.Stock, p.CreatedAt)

	mock.ExpectQuery(query).WithArgs("random text").WillReturnRows(rows)

	product, err := ProductDomain.CreateProduct(p)

	assert.Nil(t, product)

	assert.NotNil(t, err)
}
