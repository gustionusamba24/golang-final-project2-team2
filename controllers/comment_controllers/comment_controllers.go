package comment_controllers

import (
	"github.com/gin-gonic/gin"
	"golang-final-project2-team2/resources/comment_resource"
	"golang-final-project2-team2/services/comment_services"
	"golang-final-project2-team2/utils/error_utils"
	"golang-final-project2-team2/utils/success_utils"
	"net/http"
)

func CreateComment(c *gin.Context) {
	var commentReq comment_resource.CommentCreateRequest
	userIdToken := c.MustGet("user_id")
	if err := c.ShouldBindJSON(&commentReq); err != nil {
		theErr := error_utils.NewUnprocessibleEntityError(err.Error())
		c.JSON(theErr.Status(), theErr)
		return
	}

	comment, err := comment_services.CommentService.CreateComment(&commentReq, userIdToken.(string))

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, comment)
}

func GetComments(c *gin.Context) {
	user, err := comment_services.CommentService.GetComments()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateComment(c *gin.Context) {
	var commentReq comment_resource.CommentUpdateRequest

	userIdToken := c.MustGet("user_id")
	commentIdParam := c.Param("commentId")

	if err := c.ShouldBindJSON(&commentReq); err != nil {
		theErr := error_utils.NewUnprocessibleEntityError(err.Error())
		c.JSON(theErr.Status(), theErr)
		return
	}

	comment, err := comment_services.CommentService.UpdateComment(&commentReq, userIdToken.(string), commentIdParam)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, comment)
}

func DeleteComment(c *gin.Context) {
	userIdToken := c.MustGet("user_id")
	commentIdParam := c.Param("commentId")

	err := comment_services.CommentService.DeleteComment(userIdToken.(string), commentIdParam)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, success_utils.Success("Delete comment successfully"))
}
