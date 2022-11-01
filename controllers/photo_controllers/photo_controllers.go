package photo_controllers

import (
	"github.com/gin-gonic/gin"
	"golang-final-project2-team2/resources/photo_resources"
	"golang-final-project2-team2/services/photo_services"
	"golang-final-project2-team2/utils/error_utils"
	"net/http"
)

func CreatePhoto(c *gin.Context) {
	var photoReq photo_resources.PhotoCreateRequest
	userIdParam := c.MustGet("user_id")
	if err := c.ShouldBindJSON(&photoReq); err != nil {
		theErr := error_utils.NewUnprocessibleEntityError(err.Error())
		c.JSON(theErr.Status(), theErr)
		return
	}

	photo, err := photo_services.PhotoService.CreatePhoto(&photoReq, userIdParam.(string))

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, photo)
}

func GetPhotos(c *gin.Context) {
	user, err := photo_services.PhotoService.GetPhotos()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, user)
}

//func UserLogin(c *gin.Context) {
//	var userReq user_resources.UserLoginRequest
//	if err := c.ShouldBindJSON(&userReq); err != nil {
//		theErr := error_utils.NewUnprocessibleEntityError(err.Error())
//		c.JSON(theErr.Status(), theErr)
//		return
//	}
//
//	user, err := user_services.UserService.UserLogin(&userReq)
//
//	if err != nil {
//		c.JSON(err.Status(), err)
//		return
//	}
//
//	c.JSON(http.StatusOK, user)
//}
//
//func UpdateUser(c *gin.Context) {
//	var userReq user_resources.UserUpdateRequest
//	userIdParam := c.Param("userId")
//	userIdToken := c.MustGet("user_id")
//	if userIdToken != userIdParam {
//		c.JSON(error_formats.NoAuthorization().Status(), error_formats.NoAuthorization())
//		return
//	}
//	if err := c.ShouldBindJSON(&userReq); err != nil {
//		theErr := error_utils.NewUnprocessibleEntityError(err.Error())
//		c.JSON(theErr.Status(), theErr)
//		return
//	}
//
//	user, err := user_services.UserService.UserUpdate(userIdParam, &userReq)
//
//	if err != nil {
//		c.JSON(err.Status(), err)
//		return
//	}
//
//	c.JSON(http.StatusOK, user)
//}
//
//func DeleteUser(c *gin.Context) {
//	userIdToken := c.MustGet("user_id")
//
//	err := user_services.UserService.UserDelete(userIdToken.(string))
//
//	if err != nil {
//		c.JSON(err.Status(), err)
//		return
//	}
//
//	c.JSON(http.StatusOK, success_utils.Success("Your account has been successfully deleted"))
//}

//func UpdateProduct(c *gin.Context) {
//	var productReq product_domain.Product
//
//	if err := c.ShouldBindJSON(&productReq); err != nil {
//		theErr := error_utils.NewUnprocessibleEntityError("invalid json body")
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
