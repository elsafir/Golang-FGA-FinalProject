package controllers

import (
	"Golang-FGA-FinalProject/helpers"
	"Golang-FGA-FinalProject/params"
	"Golang-FGA-FinalProject/services"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type PhotoController struct {
	photoService services.PhotoService
}

func NewPhotoController(photoService *services.PhotoService) *PhotoController {
	return &PhotoController{
		photoService: *photoService,
	}
}

func (p *PhotoController) CreatePhoto(ctx *gin.Context) {
	var request params.CreatePhoto

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		helpers.ResponseStatusBadRequest(ctx, "BAD REQUEST")
		return
	}

	authId := ctx.MustGet("id").(float64)

	request.UserID = uint(authId)

	_, err = govalidator.ValidateStruct(&request)
	if err != nil {
		helpers.ResponseStatusBadRequestErrorValidation(ctx, err.Error())
		return
	}

	response := p.photoService.Create(&request)
	ctx.JSON(response.Status, response)
}

func (p *PhotoController) GetPhotos(ctx *gin.Context) {
	response := p.photoService.FindAll()
	ctx.JSON(response.Status, response)
}

func (p *PhotoController) UpdatePhoto(ctx *gin.Context) {
	var request params.UpdatePhoto

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		helpers.ResponseStatusBadRequest(ctx, "BAD REQUEST")
		return
	}

	photoIdString := ctx.Param("photoId")
	idFromParam, err := strconv.Atoi(photoIdString)
	if err != nil {
		helpers.ResponseStatusBadRequestCannotConvertId(ctx, err.Error())
		return
	}

	authId := ctx.MustGet("id").(float64)

	request.ID = uint(idFromParam)
	request.UserID = uint(authId)

	_, err = govalidator.ValidateStruct(&request)
	if err != nil {
		helpers.ResponseStatusBadRequestErrorValidation(ctx, err.Error())
		return
	}

	response := p.photoService.Update(&request)
	ctx.JSON(response.Status, response)
}

func (p *PhotoController) DeletePhoto(ctx *gin.Context) {
	photoIdString := ctx.Param("photoId")
	idFromParam, err := strconv.Atoi(photoIdString)
	if err != nil {
		helpers.ResponseStatusBadRequestCannotConvertId(ctx, err.Error())
		return
	}

	authId := ctx.MustGet("id").(float64)

	response := p.photoService.Delete(uint(idFromParam), uint(authId))
	ctx.JSON(response.Status, response)
}
