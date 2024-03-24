package services

import (
	models "Golang-FGA-FinalProject/model"
	"Golang-FGA-FinalProject/params"
	"Golang-FGA-FinalProject/repositories"
	"fmt"
	"net/http"
)

type SocialMediaService struct {
	socialMediaRepo repositories.SocialMediaRepo
}

func NewSocialMediaService(socialMediaRepo repositories.SocialMediaRepo) *SocialMediaService {
	return &SocialMediaService{
		socialMediaRepo: socialMediaRepo,
	}
}

func (s *SocialMediaService) Create(request *params.CreateSocialMedia) *params.Response {
	modelSocialMedia := models.SocialMedia{
		Name:           request.Name,
		SocialMediaURL: request.SocialMediaURL,
		UserID:         request.UserID,
	}

	socialMedia, err := s.socialMediaRepo.Create(&modelSocialMedia)
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "Bad Request",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusCreated,
		Message: "Success create new social media",
		Data: &params.SocialMediaResponse{
			ID:             int(socialMedia.ID),
			Name:           socialMedia.Name,
			SocialMediaURL: socialMedia.SocialMediaURL,
			UserID:         int(socialMedia.UserID),
			CreatedAt:      socialMedia.CreatedAt,
		},
	}
}

func (s *SocialMediaService) FindAllByAuthId(authId uint) *params.Response {
	socialMedias, err := s.socialMediaRepo.FindAllByAuthId(authId)
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "Bad Request",
			AdditionalInfo: err.Error(),
		}
	}

	if len(*socialMedias) == 0 {
		return &params.Response{
			Status: http.StatusNotFound,
			Error:  "Data is Empty",
		}
	}

	var socialMediaResponses []params.SocialMediaResponse
	for _, socialMedia := range *socialMedias {
		socialMediaResponses = append(socialMediaResponses, params.SocialMediaResponse{
			ID:             int(socialMedia.ID),
			Name:           socialMedia.Name,
			SocialMediaURL: socialMedia.SocialMediaURL,
			UserID:         int(socialMedia.UserID),
			CreatedAt:      socialMedia.CreatedAt,
			UpdatedAt:      socialMedia.UpdatedAt,
			User: params.UserResponse{
				ID:       int(socialMedia.User.ID),
				Username: socialMedia.User.Username,
				Email:    socialMedia.User.Email,
			},
		})
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "Success retrieve all data",
		Data:    socialMediaResponses,
	}
}

func (s *SocialMediaService) Update(request *params.UpdateSocialMedia) *params.Response {
	socialMedia, err := s.socialMediaRepo.FindByIdAndAuthId(request.ID, request.UserID)
	if err != nil {
		return &params.Response{
			Status:         http.StatusUnauthorized,
			Error:          "UNAUTHORIZED",
			AdditionalInfo: "You can't change this data",
		}
	}

	socialMedia.Name = request.Name
	socialMedia.SocialMediaURL = request.SocialMediaURL

	socialMedia, err = s.socialMediaRepo.Update(socialMedia)
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("Update data social media with id %d success", request.ID),
		Data: &params.SocialMediaResponse{
			ID:             int(socialMedia.ID),
			Name:           socialMedia.Name,
			SocialMediaURL: socialMedia.SocialMediaURL,
			UserID:         int(socialMedia.UserID),
			UpdatedAt:      socialMedia.UpdatedAt,
		},
	}
}

func (s *SocialMediaService) Delete(socialMediaId, authId uint) *params.Response {
	socialMedia, err := s.socialMediaRepo.FindByIdAndAuthId(socialMediaId, authId)
	if err != nil {
		return &params.Response{
			Status:         http.StatusUnauthorized,
			Error:          "UNAUTHORIZED",
			AdditionalInfo: "You can't delete this data",
		}
	}

	_, err = s.socialMediaRepo.Delete(socialMedia)
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "Your social media has been successfully deleted",
	}
}
