package services

import (
	models "Golang-FGA-FinalProject/model"
	"Golang-FGA-FinalProject/params"
	"Golang-FGA-FinalProject/repositories"
	"net/http"
)

type CommentService struct {
	commentRepo repositories.CommentRepo
}

func NewCommentService(commentRepo repositories.CommentRepo) *CommentService {
	return &CommentService{
		commentRepo: commentRepo,
	}
}

func (c *CommentService) FindAll() *params.Response {
	comments, err := c.commentRepo.FindAll()
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		}
	}

	if len(*comments) == 0 {
		return &params.Response{
			Status: http.StatusNotFound,
			Error:  "DATA IS EMPTY",
		}
	}

	var commentResponses []params.CommentResponse
	for _, comment := range *comments {
		commentResponses = append(commentResponses, params.CommentResponse{
			ID:        int(comment.ID),
			Message:   comment.Message,
			PhotoID:   int(comment.PhotoID),
			UserID:    int(comment.UserID),
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
			User: &params.UserResponse{
				ID:       int(comment.User.ID),
				Email:    comment.User.Email,
				Username: comment.User.Username,
			},
			Photo: &params.PhotoResponse{
				ID:       int(comment.Photo.ID),
				Title:    comment.Photo.Title,
				Caption:  comment.Photo.Caption,
				PhotoURL: comment.Photo.PhotoURL,
				UserID:   int(comment.Photo.UserID),
			},
		})
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "Success retrieve comments all data",
		Data:    &commentResponses,
	}
}

func (c *CommentService) Create(request *params.CreateComment) *params.Response {
	modelComment := models.Comment{
		Message: request.Message,
		PhotoID: uint(request.PhotoID),
		UserID:  request.UserID,
	}

	comment, err := c.commentRepo.Create(&modelComment)
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "Success create comment",
		Data: &params.CommentResponse{
			ID:        int(comment.ID),
			Message:   comment.Message,
			UserID:    int(comment.UserID),
			PhotoID:   int(comment.PhotoID),
			CreatedAt: comment.CreatedAt,
		},
	}
}

func (c *CommentService) Update(request *params.UpdateComment) *params.Response {
	comment, err := c.commentRepo.FindByIdAndAuthId(request.ID, request.UserID)
	if err != nil {
		return &params.Response{
			Status:         http.StatusUnauthorized,
			Error:          "UNAUTHORIZED",
			AdditionalInfo: "You can't change this data",
		}
	}

	// update message
	comment.Message = request.Message

	comment, err = c.commentRepo.Update(comment)
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "Your comment has been successfully updated",
		Data: &params.CommentResponse{
			ID:        int(comment.ID),
			Message:   comment.Message,
			PhotoID:   int(comment.PhotoID),
			UserID:    int(comment.UserID),
			UpdatedAt: comment.UpdatedAt,
		},
	}
}

func (c *CommentService) Delete(commentId, authId uint) *params.Response {
	comment, err := c.commentRepo.FindByIdAndAuthId(commentId, authId)
	if err != nil {
		return &params.Response{
			Status:         http.StatusUnauthorized,
			Error:          "UNAUTHORIZED",
			AdditionalInfo: "You can't delete this data",
		}
	}

	_, err = c.commentRepo.Delete(comment)
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "Your comment has been successfully deleted",
	}
}
