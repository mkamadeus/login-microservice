package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username	string `gorm:"uniqueIndex;not null" json:"username"`
	Password	string `json:"password"`
	Name			string `json:"name"`
}

type UserResponse struct {
	Name			string `json:"name"`
	Username	string `json:"username`
}
