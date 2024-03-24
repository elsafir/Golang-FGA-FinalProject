package controllers

import (
	"Golang-FGA-FinalProject/helpers"
	"Golang-FGA-FinalProject/params"
	"Golang-FGA-FinalProject/services"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: *userService,
	}
}

func (u *UserController) Register(ctx *gin.Context) {
	var request params.RegisterUser

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		helpers.ResponseStatusBadRequest(ctx, err.Error())
		return
	}

	_, err = govalidator.ValidateStruct(&request)
	if err != nil {
		helpers.ResponseStatusBadRequestErrorValidation(ctx, err.Error())
		return
	}

	response := u.userService.Create(&request)
	ctx.JSON(response.Status, response)
}

func (u *UserController) Login(ctx *gin.Context) {
	var request params.LoginUser

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		helpers.ResponseStatusBadRequest(ctx, err.Error())
		return
	}

	_, err = govalidator.ValidateStruct(&request)
	if err != nil {
		helpers.ResponseStatusBadRequestErrorValidation(ctx, err.Error())
		return
	}

	response := u.userService.Login(&request)
	ctx.JSON(response.Status, response)
}

func (u *UserController) UpdateUser(ctx *gin.Context) {
	var request params.UpdateUser

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		helpers.ResponseStatusBadRequest(ctx, err.Error())
		return
	}

	userIdString := ctx.Param("userId")
	idFromParam, err := strconv.Atoi(userIdString)
	if err != nil {
		helpers.ResponseStatusBadRequestCannotConvertId(ctx, err.Error())
		return
	}

	// check id from token and from param
	AuthId := ctx.MustGet("id").(float64)
	if int(AuthId) != idFromParam {
		helpers.ResponseStatusUnauthorized(ctx, "You can't change this data")
		return
	}

	request.ID = uint(idFromParam)

	_, err = govalidator.ValidateStruct(&request)
	if err != nil {
		helpers.ResponseStatusBadRequestErrorValidation(ctx, err.Error())
		return
	}

	response := u.userService.Update(&request)
	ctx.JSON(response.Status, response)
}

func (u *UserController) DeleteUser(ctx *gin.Context) {
	userIdString := ctx.Param("userId")
	idFromParam, err := strconv.Atoi(userIdString)
	if err != nil {
		helpers.ResponseStatusBadRequestCannotConvertId(ctx, err.Error())
		return
	}

	// check id from token and from param
	AuthId := ctx.MustGet("id").(float64)
	if int(AuthId) != idFromParam {
		helpers.ResponseStatusUnauthorized(ctx, "You can't change this data")
		return
	}

	response := u.userService.Delete(uint(idFromParam))
	ctx.JSON(response.Status, response)
}
