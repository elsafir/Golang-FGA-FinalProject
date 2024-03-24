package controllers

import (
	"Golang-FGA-FinalProject/helpers"
	"Golang-FGA-FinalProject/params"
	"Golang-FGA-FinalProject/services"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type CommentController struct {
	commentService services.CommentService
	photoService   services.PhotoService
}

func NewCommentController(commentService *services.CommentService, photoService *services.PhotoService) *CommentController {
	return &CommentController{
		commentService: *commentService,
		photoService:   *photoService,
	}
}

func (c *CommentController) CreateComment(ctx *gin.Context) {
	var request params.CreateComment

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		helpers.ResponseStatusBadRequest(ctx, err.Error())
		return
	}

	if request.PhotoID == 0 {
		helpers.ResponseStatusBadRequest(ctx, "Field Photo Id is required")
		return
	}

	isPhotoExist := c.photoService.IsPhotoExist(uint(request.PhotoID))
	if !isPhotoExist {
		helpers.ResponseStatusNotFound(ctx)
		return
	}

	authId := ctx.MustGet("id").(float64)

	request.UserID = uint(authId)

	_, err = govalidator.ValidateStruct(&request)
	if err != nil {
		helpers.ResponseStatusBadRequestErrorValidation(ctx, err.Error())
		return
	}

	response := c.commentService.Create(&request)
	ctx.JSON(response.Status, response)
}

func (c *CommentController) GetComments(ctx *gin.Context) {
	response := c.commentService.FindAll()
	ctx.JSON(response.Status, response)
}

func (c *CommentController) UpdateComment(ctx *gin.Context) {
	var request params.UpdateComment

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		helpers.ResponseStatusBadRequest(ctx, "BAD REQUEST")
		return
	}

	idFromParam, err := helpers.ConvertStringParamToInt(ctx, "commentId")
	if err != nil {
		helpers.ResponseStatusBadRequestCannotConvertId(ctx, err.Error())
		return
	}
	authId := ctx.MustGet("id").(float64)

	request.ID = uint(idFromParam)
	request.UserID = uint(authId)

	_, err = govalidator.ValidateStruct(request)
	if err != nil {
		helpers.ResponseStatusBadRequestErrorValidation(ctx, err.Error())
		return
	}

	response := c.commentService.Update(&request)
	ctx.JSON(response.Status, response)
}

func (c *CommentController) DeleteComment(ctx *gin.Context) {
	idFromParam, err := helpers.ConvertStringParamToInt(ctx, "commentId")
	if err != nil {
		helpers.ResponseStatusBadRequestCannotConvertId(ctx, err.Error())
		return
	}
	authId := ctx.MustGet("id").(float64)

	response := c.commentService.Delete(uint(idFromParam), uint(authId))
	ctx.JSON(response.Status, response)
}
