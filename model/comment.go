package models

import "time"

type Comment struct {
	ID      uint   `gorm:"primary_key" json:"id"`
	Message string `gorm:"not null" json:"message"`

	// Comment Has One User
	UserID uint `gorm:"type:int;not null;" json:"user_id"`
	User   User
	// Comment Has One Photo
	PhotoID uint `gorm:"type:int;not null;" json:"photo_id"`
	Photo   Photo

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
