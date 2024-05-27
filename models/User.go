package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FirstName string `gorm:"not null" json:"first_name"`
	LastName  string `gorm:"not null" json:"last_name"`
	Email     string `gorm:"not null;uniqueIndex" json:"email"`
	Username  string `gorm:"uniqueIndex" json:"username"`
	Password  string `json:"password"`
	Tasks     []Task `json:"tasks"`
}
