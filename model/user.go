package models

import (
	"time"
)

type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Email    string `gorm:"type:varchar(255);unique_index;not null" json:"email"`
	Username string `gorm:"type:varchar(255);unique_index;not null" json:"username"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
	Age      int   `gorm:"size:11;not null" json:"age"`

	// User Has Many Photos
	// User can posts many photos
	Photos []Photo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// User Has Many SocialMedias
	// User can have many social media accounts
	SocialMedias []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// User Has Many Comments
	// User can have many comments on each photo
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
