package social_media_controllers

import (
	"github.com/gin-gonic/gin"
	"golang-final-project2-team2/resources/social_media_resource"
	"golang-final-project2-team2/services/social_media_services"
	"golang-final-project2-team2/utils/error_utils"
	"net/http"
)

func CreateSocialMedia(c *gin.Context) {
	var socialMediaReq social_media_resource.SocialMediaCreateRequest
	userIdToken := c.MustGet("user_id")
	if err := c.ShouldBindJSON(&socialMediaReq); err != nil {
		theErr := error_utils.NewUnprocessibleEntityError(err.Error())
		c.JSON(theErr.Status(), theErr)
		return
	}

	socialMedia, err := social_media_services.SocialMediaService.CreateSocialMedia(&socialMediaReq, userIdToken.(string))

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, socialMedia)
}

func GetSocialMedias(c *gin.Context) {
	socialMedia, err := social_media_services.SocialMediaService.GetSocialMedias()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, socialMedia)
}

func UpdateSocialMedia(c *gin.Context) {
	var socialMediaReq social_media_resource.SocialMediaUpdateRequest

	userIdToken := c.MustGet("user_id")
	socialMediaIdParam := c.Param("socialMediaId")

	if err := c.ShouldBindJSON(&socialMediaReq); err != nil {
		theErr := error_utils.NewUnprocessibleEntityError(err.Error())
		c.JSON(theErr.Status(), theErr)
		return
	}

	socialMedia, err := social_media_services.SocialMediaService.UpdateSocialMedia(&socialMediaReq, userIdToken.(string), socialMediaIdParam)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, socialMedia)
}

func DeleteSocialMedia(c *gin.Context) {
	userIdToken := c.MustGet("user_id")
	socialMediaIdParam := c.Param("socialMediaId")

	err := social_media_services.SocialMediaService.DeleteSocialMedia(userIdToken.(string), socialMediaIdParam)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, "success")
}