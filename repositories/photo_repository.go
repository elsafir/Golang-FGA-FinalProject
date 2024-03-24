package repositories

import (
	models "Golang-FGA-FinalProject/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PhotoRepo interface {
	FindAll() (*[]models.Photo, error)
	FindById(id uint) (*models.Photo, error)
	FindByIdAndAuthId(id, authId uint) (*models.Photo, error)
	Create(photo *models.Photo) (*models.Photo, error)
	Update(photo *models.Photo) (*models.Photo, error)
	Delete(photo *models.Photo) (*models.Photo, error)
}

type photoRepo struct {
	db *gorm.DB
}

func NewPhotoRepo(db *gorm.DB) PhotoRepo {
	return &photoRepo{
		db: db,
	}
}

func (p *photoRepo) FindAll() (*[]models.Photo, error) {
	var photos []models.Photo
	err := p.db.Preload(clause.Associations).Find(&photos).Error
	return &photos, err
}

func (p *photoRepo) FindById(id uint) (*models.Photo, error) {
	var photo models.Photo
	err := p.db.First(&photo, id).Error
	return &photo, err

}
func (p *photoRepo) FindByIdAndAuthId(id, authId uint) (*models.Photo, error) {
	var photo models.Photo
	err := p.db.Where("user_id=?", authId).First(&photo, id).Error
	return &photo, err
}

func (p *photoRepo) Create(photo *models.Photo) (*models.Photo, error) {
	err := p.db.Create(&photo).Error
	return photo, err
}

func (p *photoRepo) Update(photo *models.Photo) (*models.Photo, error) {
	err := p.db.Save(&photo).Error
	return photo, err
}

func (p *photoRepo) Delete(photo *models.Photo) (*models.Photo, error) {
	err := p.db.Delete(&photo).Error
	return photo, err
}
