package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(50);unique;not null"`
	Password string `gorm:"type:varchar(100);not null"`
	Name     string `gorm:"type:varchar(100)"`
}
