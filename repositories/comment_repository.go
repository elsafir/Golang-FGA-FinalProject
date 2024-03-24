package repositories

import (
	models "Golang-FGA-FinalProject/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CommentRepo interface {
	FindAll() (*[]models.Comment, error)
	FindById(id uint) (*models.Comment, error)
	FindByIdAndAuthId(id, authId uint) (*models.Comment, error)
	Create(comment *models.Comment) (*models.Comment, error)
	Update(comment *models.Comment) (*models.Comment, error)
	Delete(comment *models.Comment) (*models.Comment, error)
}

type commentRepo struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) CommentRepo {
	return &commentRepo{
		db: db,
	}
}

func (c *commentRepo) FindAll() (*[]models.Comment, error) {
	var comments []models.Comment
	err := c.db.Preload(clause.Associations).Find(&comments).Error
	return &comments, err
}

func (c *commentRepo) FindById(id uint) (*models.Comment, error) {
	var comment models.Comment
	err := c.db.First(&comment, id).Error
	return &comment, err
}

func (c *commentRepo) FindByIdAndAuthId(id, authId uint) (*models.Comment, error) {
	var comment models.Comment
	err := c.db.Where("user_id=?", authId).First(&comment, id).Error
	return &comment, err
}

func (c *commentRepo) Create(comment *models.Comment) (*models.Comment, error) {
	err := c.db.Create(&comment).Error
	return comment, err
}

func (c *commentRepo) Update(comment *models.Comment) (*models.Comment, error) {
	err := c.db.Save(&comment).Error
	return comment, err
}

func (c *commentRepo) Delete(comment *models.Comment) (*models.Comment, error) {
	err := c.db.Delete(&comment).Error
	return comment, err
}
