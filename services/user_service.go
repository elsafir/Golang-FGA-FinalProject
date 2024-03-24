package services

import (
	"Golang-FGA-FinalProject/helpers"
	"Golang-FGA-FinalProject/params"
	"Golang-FGA-FinalProject/repositories"
	"net/http"
)

type UserService struct {
	userRepo repositories.UserRepo
}

func NewUserService(userRepo repositories.UserRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (u *UserService) Create(request *params.RegisterUser) *params.Response {
	_, err := u.userRepo.FindByEmail(request.Email)
	if err == nil {
		// if user found
		return &params.Response{
			Status: http.StatusBadRequest,
			Error:  "EMAIL ALREADY REGISTERED",
		}
	}

	user, err := u.userRepo.FindByUsername(request.Username)
	if err == nil {
		// if user found
		return &params.Response{
			Status: http.StatusBadRequest,
			Error:  "USERNAME MUST BE UNIQUE",
		}
	}

	// if email not registered
	// create new user
	hashedPass, _ := helpers.HashPassword(request.Password)

	user.Username = request.Username
	user.Email = request.Email
	user.Password = hashedPass
	user.Age = request.Age

	user, err = u.userRepo.Create(user)
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusCreated,
		Message: "Register success",
		Data: params.UserResponse{
			Age:      user.Age,
			Email:    user.Email,
			ID:       int(user.ID),
			Username: user.Username,
		},
	}
}

func (u *UserService) Login(request *params.LoginUser) *params.Response {
	user, err := u.userRepo.FindByEmail(request.Email)
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "EMAIL NOT REGISTERED",
			AdditionalInfo: err.Error(),
		}
	}

	isOk := helpers.ComparePassword(user.Password, request.Password)
	if !isOk {
		return &params.Response{
			Status: http.StatusBadRequest,
			Error:  "CREDENTIAL NOT MATCH",
		}
	}

	token := helpers.GenerateToken(user.ID, user.Email)

	// if user found
	return &params.Response{
		Status:  http.StatusOK,
		Message: "Login Success",
		Data: params.UserResponse{
			Token: token,
		},
	}
}

func (u *UserService) Update(request *params.UpdateUser) *params.Response {
	user, err := u.userRepo.FindById(request.ID)
	if err != nil {
		return &params.Response{
			Status:         http.StatusNotFound,
			Error:          "USER NOT FOUND",
			AdditionalInfo: err.Error(),
		}
	}

	user.Username = request.Username
	user.Email = request.Email

	user, err = u.userRepo.Update(user)
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "Your account has been successfully updated",
		Data: params.UserResponse{
			ID:        int(user.ID),
			Email:     user.Email,
			Username:  user.Username,
			Age:       user.Age,
			UpdatedAt: user.UpdatedAt,
		},
	}
}

func (u *UserService) Delete(id uint) *params.Response {
	user, err := u.userRepo.FindById(id)
	if err != nil {
		return &params.Response{
			Status:         http.StatusNotFound,
			Error:          "USER NOT FOUND",
			AdditionalInfo: err.Error(),
		}
	}

	_, err = u.userRepo.Delete(user)
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "Your account has been successfully deleted",
	}
}
