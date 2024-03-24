package repositories

import (
	models "Golang-FGA-FinalProject/model"

	"gorm.io/gorm"
)

type UserRepo interface {
	FindById(id uint) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	Create(user *models.User) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(user *models.User) (*models.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) FindById(id uint) (*models.User, error) {
	var user models.User
	err := u.db.First(&user, id).Error
	return &user, err
}

func (u *userRepo) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := u.db.Where("email=?", email).Take(&user).Error
	return &user, err
}

func (u *userRepo) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := u.db.Where("username=?", username).Take(&user).Error
	return &user, err
}

func (u *userRepo) Create(user *models.User) (*models.User, error) {
	err := u.db.Create(&user).Error
	return user, err
}

func (u *userRepo) Update(user *models.User) (*models.User, error) {
	err := u.db.Save(&user).Error
	return user, err
}

func (u *userRepo) Delete(user *models.User) (*models.User, error) {
	err := u.db.Delete(&user).Error
	return user, err
}
