package controllers

import (
	"Golang-FGA-FinalProject/helpers"
	"Golang-FGA-FinalProject/params"
	"Golang-FGA-FinalProject/services"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type SocialMediaController struct {
	socialMediaService services.SocialMediaService
}

func NewSocialMediaController(service *services.SocialMediaService) *SocialMediaController {
	return &SocialMediaController{
		socialMediaService: *service,
	}
}

func (s *SocialMediaController) CreateSocialMedia(ctx *gin.Context) {
	var request params.CreateSocialMedia
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		helpers.ResponseStatusBadRequest(ctx, "BAD REQUEST")
		return
	}

	AuthId := ctx.MustGet("id").(float64)

	request.UserID = uint(AuthId)

	_, err = govalidator.ValidateStruct(&request)
	if err != nil {
		helpers.ResponseStatusBadRequestErrorValidation(ctx, err.Error())
		return
	}

	response := s.socialMediaService.Create(&request)
	ctx.JSON(response.Status, response)
}

func (s *SocialMediaController) GetSocialMedias(ctx *gin.Context) {
	AuthId := ctx.MustGet("id").(float64)
	response := s.socialMediaService.FindAllByAuthId(uint(AuthId))
	ctx.JSON(response.Status, response)
}

func (s *SocialMediaController) UpdateSocialMedia(ctx *gin.Context) {
	var request params.UpdateSocialMedia

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		helpers.ResponseStatusBadRequest(ctx, "BAD REQUEST")
		return
	}

	socialMediaIdString := ctx.Param("socialMediaId")
	id, err := strconv.Atoi(socialMediaIdString)
	if err != nil {
		helpers.ResponseStatusBadRequest(ctx, "BAD REQUEST")
		return
	}

	AuthId := ctx.MustGet("id").(float64)

	request.ID = uint(id)
	request.UserID = uint(AuthId)

	_, err = govalidator.ValidateStruct(&request)
	if err != nil {
		helpers.ResponseStatusBadRequestErrorValidation(ctx, err.Error())
		return
	}

	response := s.socialMediaService.Update(&request)
	ctx.JSON(response.Status, response)
}

func (s *SocialMediaController) DeleteSocialMedia(ctx *gin.Context) {
	socialMediaIdString := ctx.Param("socialMediaId")
	id, err := strconv.Atoi(socialMediaIdString)
	if err != nil {
		helpers.ResponseStatusBadRequestCannotConvertId(ctx, err.Error())
		return
	}

	AuthId := ctx.MustGet("id").(float64)

	response := s.socialMediaService.Delete(uint(id), uint(AuthId))
	ctx.JSON(response.Status, response)
}
