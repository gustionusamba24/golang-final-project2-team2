package comment_services

import (
	"golang-final-project2-team2/domains/comment_domain"
	"golang-final-project2-team2/resources/comment_resource"
	"golang-final-project2-team2/utils/error_formats"
	"golang-final-project2-team2/utils/error_utils"
	"golang-final-project2-team2/utils/helpers"
	"strconv"
)

var CommentService commentServiceRepo = &commentService{}

type commentServiceRepo interface {
	CreateComment(*comment_resource.CommentCreateRequest, string) (*comment_resource.CommentCreateResponse, error_utils.MessageErr)
	GetComments(*[]comment_resource.CommentGetResponse, error_utils.MessageErr)
	UpdateComment(*comment_resource.CommentUpdateRequest, string, string) (*comment_resource.CommentUpdateResponse, error_utils.MessageErr)
	DeleteComment(string, string) error_utils.MessageErr
}

type commentService struct{}

func (u *commentService) CreateComment(commentReq *comment_resource.CommentCreateRequest, userId string) (*comment_resource.CommentCreateResponse, error_utils.MessageErr) {
	err := helpers.ValidateRequest(commentReq)

	if err != nil {
		return nil, err
	}

	comment, err := comment_domain.CommentDomain.CreateComment(commentReq, userId)

	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (u *commentService) GetComments() (*[]comment_resource.CommentGetResponse, error_utils.MessageErr) {
	comments, err := comment_domain.CommentDomain.GetComments()

	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (u *commentService) UpdateComment(request *comment_resource.CommentUpdateRequest, userId string, commentId string) (*comment_resource.CommentUpdateResponse, error_utils.MessageErr) {
	comment, err := comment_domain.CommentDomain.GetComment(commentId)
	if err != nil {
		return nil, err
	}

	if strconv.FormatInt(comment.UserId, 10) != userId {
		return nil, error_formats.NoAuthorization()
	}

	updatedComment, err := comment_domain.CommentDomain.UpdateComment(request, commentId)

	if err != nil {
		return nil, err
	}

	return updatedComment, nil
}

func (u *commentService) DeleteComment(userId string, commentId string) error_utils.MessageErr {
	comment, err := comment_domain.CommentDomain.GetComment(commentId)
	if err != nil {
		return err
	}

	if strconv.FormatInt(comment.UserId, 10) != userId {
		return error_formats.NoAuthorization()
	}

	err = comment_domain.CommentDomain.DeleteComment(commentId)

	if err != nil {
		return err
	}

	return nil
}
