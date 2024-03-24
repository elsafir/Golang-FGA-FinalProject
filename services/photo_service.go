package services

import (
	models "Golang-FGA-FinalProject/model"
	"Golang-FGA-FinalProject/params"
	"Golang-FGA-FinalProject/repositories"
	"net/http"
)

type PhotoService struct {
	photoRepo repositories.PhotoRepo
}

func NewPhotoService(photoRepo repositories.PhotoRepo) *PhotoService {
	return &PhotoService{
		photoRepo: photoRepo,
	}
}

func (p *PhotoService) Create(request *params.CreatePhoto) *params.Response {
	modelPhoto := models.Photo{
		Title:    request.Title,
		Caption:  request.Caption,
		PhotoURL: request.PhotoURL,
		UserID:   request.UserID,
	}

	photo, err := p.photoRepo.Create(&modelPhoto)
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusCreated,
		Message: "Success post a photo",
		Data: &params.PhotoResponse{
			ID:        int(photo.ID),
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoURL:  photo.PhotoURL,
			UserID:    int(photo.UserID),
			CreatedAt: photo.CreatedAt,
		},
	}
}

func (p *PhotoService) FindAll() *params.Response {
	photos, err := p.photoRepo.FindAll()
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		}
	}

	if len(*photos) == 0 {
		return &params.Response{
			Status: http.StatusNotFound,
			Error:  "DATA IS EMPTY",
		}
	}

	var photoResponses []params.PhotoResponse
	for _, photo := range *photos {
		photoResponses = append(photoResponses, params.PhotoResponse{
			ID:        int(photo.ID),
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoURL:  photo.PhotoURL,
			UserID:    int(photo.UserID),
			CreatedAt: photo.CreatedAt,
			UpdatedAt: photo.UpdatedAt,
			User: &params.UserResponse{
				Email:    photo.User.Email,
				Username: photo.User.Username,
			},
		})
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "Success retrieve all data",
		Data:    photoResponses,
	}
}

func (p *PhotoService) Update(request *params.UpdatePhoto) *params.Response {
	photo, err := p.photoRepo.FindByIdAndAuthId(request.ID, request.UserID)
	if err != nil {
		return &params.Response{
			Status:         http.StatusUnauthorized,
			Error:          "UNAUTHORIZED",
			AdditionalInfo: "You can't change this data",
		}
	}

	photo.Title = request.Title
	photo.Caption = request.Caption
	photo.PhotoURL = request.PhotoURL

	photo, err = p.photoRepo.Update(photo)
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "Your photo has been successfully updated",
		Data: &params.PhotoResponse{
			ID:        int(photo.ID),
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoURL:  photo.PhotoURL,
			UserID:    int(photo.UserID),
			UpdatedAt: photo.UpdatedAt,
		},
	}
}

func (p *PhotoService) Delete(photoId, authId uint) *params.Response {
	photo, err := p.photoRepo.FindByIdAndAuthId(photoId, authId)
	if err != nil {
		return &params.Response{
			Status:         http.StatusUnauthorized,
			Error:          "UNAUTHORIZED",
			AdditionalInfo: "You can't delete this data",
		}
	}

	_, err = p.photoRepo.Delete(photo)
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "Your photo has been successfully deleted",
	}
}

func (p *PhotoService) IsPhotoExist(id uint) bool {
	_, err := p.photoRepo.FindById(id)
	return err == nil
}
