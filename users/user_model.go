package users

import (
	"awesomeProject/shared"
)

type User struct {
	shared.Model
	FirstName string `gorm:"size;250:not null;" json:"first_name"`
	LastName string `gorm:"size;250:not null;" json:"last_name"`
	Email string `gorm:"email:size;1000:not null;" json:"email"`
}