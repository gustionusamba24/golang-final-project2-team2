package product_controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"products/domain/product_domain"
	"products/service/product_service"
	"products/utils/error_utils"
	"testing"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	tm            = time.Now()
	createProduct func(*product_domain.Product) (*product_domain.Product, error_utils.MessageErr)
	updateProduct func(*product_domain.Product) (*product_domain.Product, error_utils.MessageErr)
	getProducts   func() ([]*product_domain.Product, error_utils.MessageErr)
	deleteProduct func(int64) error_utils.MessageErr
)

type productServiceMock struct{}

func (p *productServiceMock) CreateProduct(productReq *product_domain.Product) (*product_domain.Product, error_utils.MessageErr) {
	return createProduct(productReq)
}

func (p *productServiceMock) UpdateProduct(productReq *product_domain.Product) (*product_domain.Product, error_utils.MessageErr) {
	return updateProduct(productReq)
}

func (p *productServiceMock) GetProducts() ([]*product_domain.Product, error_utils.MessageErr) {
	return getProducts()
}

func (p *productServiceMock) DeleteProduct(productId int64) error_utils.MessageErr {
	return deleteProduct(productId)
}

func TestProductController_CreateProduct_Success(t *testing.T) {
	product_service.ProductService = &productServiceMock{}

	requestBody := `
		{
			"name": "Product Test",
			"price": 30.50,
			"stock": 100
		}
	`

	expectedVal := &product_domain.Product{
		Id:        1,
		Name:      "Product Test",
		Price:     10.10,
		Stock:     5,
		CreatedAt: tm,
	}

	createProduct = func(p *product_domain.Product) (*product_domain.Product, error_utils.MessageErr) {
		return expectedVal, nil
	}

	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBufferString(requestBody))

	router := gin.Default()

	gin.SetMode(gin.TestMode)

	rr := httptest.NewRecorder()

	router.POST("/products", CreateProduct)

	router.ServeHTTP(rr, req)

	result := rr.Result()

	data, err := ioutil.ReadAll(result.Body)

	defer result.Body.Close()

	require.Nil(t, err)

	var product product_domain.Product

	err = json.Unmarshal(data, &product)

	assert.Nil(t, err)
	assert.NotNil(t, product)

	assert.EqualValues(t, expectedVal.Id, product.Id)
	assert.EqualValues(t, expectedVal.Name, product.Name)
	assert.EqualValues(t, expectedVal.Price, product.Price)
	assert.EqualValues(t, expectedVal.Stock, product.Stock)
}

func TestProductController_CreateProduct_InvalidJSONBody(t *testing.T) {
	requestBody := `
		{
			"name": "name",
			"price": 30.50,
			"stock": 100
		
	`

	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBufferString(requestBody))

	router := gin.Default()

	gin.SetMode(gin.TestMode)

	rr := httptest.NewRecorder()

	router.POST("/products", CreateProduct)

	router.ServeHTTP(rr, req)

	result := rr.Result()

	data, err := ioutil.ReadAll(result.Body)

	defer result.Body.Close()

	require.Nil(t, err)

	var errData error_utils.MessageErrData

	err = json.Unmarshal(data, &errData)

	assert.Nil(t, err)
	assert.EqualValues(t, "invalid json body", errData.ErrMessage)
	assert.EqualValues(t, "invalid_request", errData.ErrError)
	assert.EqualValues(t, 422, errData.Status())
}

func TestProductController_CreateProduct_ServerError(t *testing.T) {
	product_service.ProductService = &productServiceMock{}

	requestBody := `
		{
			"name": "Product Test",
			"price": 30.50,
			"stock": 100
		}
	`

	createProduct = func(p *product_domain.Product) (*product_domain.Product, error_utils.MessageErr) {
		return nil, error_utils.NewInternalServerErrorr("something went wrong")
	}
	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBufferString(requestBody))

	router := gin.Default()

	gin.SetMode(gin.TestMode)

	rr := httptest.NewRecorder()

	router.POST("/products", CreateProduct)

	router.ServeHTTP(rr, req)

	result := rr.Result()

	data, err := ioutil.ReadAll(result.Body)

	defer result.Body.Close()

	require.Nil(t, err)

	var errData error_utils.MessageErrData

	err = json.Unmarshal(data, &errData)

	assert.Nil(t, err)
	assert.EqualValues(t, "something went wrong", errData.ErrMessage)
	assert.EqualValues(t, "server_error", errData.ErrError)
	assert.EqualValues(t, 500, errData.Status())
}


// Tugas Positif Case
func TestProductController_UpdateProduct_Success(t *testing.T) {
	product_service.ProductService = &productServiceMock{}

	requestBody := `
		{
			"id" : 1,
			"name": "Product Test2",
			"price": 50.50,
			"stock": 100
		}
	`

	expectedVal := &product_domain.Product{
		Id:        1,
		Name:      "Product Test2",
		Price:     50.50,
		Stock:     100,
		CreatedAt: tm,
	}

	updateProduct = func(p *product_domain.Product) (*product_domain.Product, error_utils.MessageErr) {
		return expectedVal, nil
	}

	req := httptest.NewRequest(http.MethodPut, "/products", bytes.NewBufferString(requestBody))

	router := gin.Default()

	gin.SetMode(gin.TestMode)

	rr := httptest.NewRecorder()

	router.PUT("/products", UpdateProduct)

	router.ServeHTTP(rr, req)

	result := rr.Result()

	data, err := ioutil.ReadAll(result.Body)

	defer result.Body.Close()

	require.Nil(t, err)

	var product product_domain.Product

	err = json.Unmarshal(data, &product)

	assert.Nil(t, err)
	assert.NotNil(t, product)

	assert.EqualValues(t, expectedVal.Id, product.Id)
	assert.EqualValues(t, expectedVal.Name, product.Name)
	assert.EqualValues(t, expectedVal.Price, product.Price)
	assert.EqualValues(t, expectedVal.Stock, product.Stock)
}

// Tugas Negatif Case
func TestProductController_UpdateProduct_ServerError(t *testing.T) {
	product_service.ProductService = &productServiceMock{}

	requestBody := `
		{
			"id" : 1,
			"name": "Product Test2",
			"price": 50.50,
			"stock": 100
		}
	`

	updateProduct = func(p *product_domain.Product) (*product_domain.Product, error_utils.MessageErr) {
		return nil, error_utils.NewInternalServerErrorr("something went wrong")
	}
	req := httptest.NewRequest(http.MethodPut, "/products", bytes.NewBufferString(requestBody))

	router := gin.Default()

	gin.SetMode(gin.TestMode)

	rr := httptest.NewRecorder()

	router.PUT("/products", UpdateProduct)

	router.ServeHTTP(rr, req)

	result := rr.Result()

	data, err := ioutil.ReadAll(result.Body)

	defer result.Body.Close()

	require.Nil(t, err)

	var errData error_utils.MessageErrData

	err = json.Unmarshal(data, &errData)

	assert.Nil(t, err)
	assert.EqualValues(t, "something went wrong", errData.ErrMessage)
	assert.EqualValues(t, "server_error", errData.ErrError)
	assert.EqualValues(t, 500, errData.Status())
}