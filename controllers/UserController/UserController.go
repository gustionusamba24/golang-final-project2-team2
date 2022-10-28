package UserController

import (
	"github.com/gin-gonic/gin"
	"golang-final-project2-team2/domains/UserDomain"
	"golang-final-project2-team2/services/UserService"
	"golang-final-project2-team2/utils/ErrorUtils"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var userReq UserDomain.User
	if err := c.ShouldBindJSON(&userReq); err != nil {
		theErr := ErrorUtils.NewUnprocessibleEntityError(err.Error())
		c.JSON(theErr.Status(), theErr)
		return
	}

	user, err := UserService.UserService.CreateUser(&userReq)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, user)
}

//func UpdateProduct(c *gin.Context) {
//	var productReq product_domain.Product
//
//	if err := c.ShouldBindJSON(&productReq); err != nil {
//		theErr := ErrorUtils.NewUnprocessibleEntityError("invalid json body")
//		c.JSON(theErr.Status(), theErr)
//		return
//	}
//
//	productId, err := productReq.GetProductIdParam(c)
//
//	if err != nil {
//		c.JSON(err.Status(), err)
//		return
//	}
//
//	productReq.Id = productId
//
//	product, err := product_service.ProductService.UpdateProduct(&productReq)
//
//	if err != nil {
//		c.JSON(err.Status(), err)
//		return
//	}
//
//	c.JSON(http.StatusOK, product)
//
//}
//
//func GetProducts(c *gin.Context) {
//	products, err := product_service.ProductService.GetProducts()
//
//	if err != nil {
//		c.JSON(err.Status(), err)
//		return
//	}
//
//	c.JSON(http.StatusOK, products)
//}
//
//func DeleteProduct(c *gin.Context) {
//	var productReq product_domain.Product
//
//	productId, err := productReq.GetProductIdParam(c)
//
//	if err != nil {
//		c.JSON(err.Status(), err)
//		return
//	}
//
//	err = product_service.ProductService.DeleteProduct(productId)
//
//	if err != nil {
//		c.JSON(err.Status(), err)
//		return
//	}
//
//	c.JSON(http.StatusNoContent, nil)
//}
