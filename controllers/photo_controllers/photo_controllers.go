package photo_controllers

import (
	"github.com/gin-gonic/gin"
	"golang-final-project2-team2/resources/photo_resources"
	"golang-final-project2-team2/services/photo_services"
	"golang-final-project2-team2/utils/error_utils"
	"golang-final-project2-team2/utils/success_utils"
	"net/http"
)

func CreatePhoto(c *gin.Context) {
	var photoReq photo_resources.PhotoCreateRequest
	userIdToken := c.MustGet("user_id")
	if err := c.ShouldBindJSON(&photoReq); err != nil {
		theErr := error_utils.NewUnprocessibleEntityError(err.Error())
		c.JSON(theErr.Status(), theErr)
		return
	}

	photo, err := photo_services.PhotoService.CreatePhoto(&photoReq, userIdToken.(string))

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

func UpdatePhoto(c *gin.Context) {
	var photoReq photo_resources.PhotoUpdateRequest

	userIdToken := c.MustGet("user_id")
	photoIdParam := c.Param("photoId")

	if err := c.ShouldBindJSON(&photoReq); err != nil {
		theErr := error_utils.NewUnprocessibleEntityError(err.Error())
		c.JSON(theErr.Status(), theErr)
		return
	}
	user, err := photo_services.PhotoService.UpdatePhoto(&photoReq, userIdToken.(string), photoIdParam)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeletePhoto(c *gin.Context) {
	userIdToken := c.MustGet("user_id")
	photoIdParam := c.Param("photoId")

	err := photo_services.PhotoService.DeletePhoto(userIdToken.(string), photoIdParam)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, success_utils.Success("Your account has been successfully deleted"))
}
