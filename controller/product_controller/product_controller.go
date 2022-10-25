package product_controller

import (
	"net/http"
	"products/domain/product_domain"
	"products/service/product_service"
	"products/utils/error_utils"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var productReq product_domain.Product

	if err := c.ShouldBindJSON(&productReq); err != nil {
		theErr := error_utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}

	product, err := product_service.ProductService.CreateProduct(&productReq)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, product)
}

func UpdateProduct(c *gin.Context) {
	var productReq product_domain.Product

	if err := c.ShouldBindJSON(&productReq); err != nil {
		theErr := error_utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}

	productId, err := productReq.GetProductIdParam(c)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	productReq.Id = productId

	product, err := product_service.ProductService.UpdateProduct(&productReq)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, product)

}

func GetProducts(c *gin.Context) {
	products, err := product_service.ProductService.GetProducts()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, products)
}

func DeleteProduct(c *gin.Context) {
	var productReq product_domain.Product

	productId, err := productReq.GetProductIdParam(c)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	err = product_service.ProductService.DeleteProduct(productId)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
