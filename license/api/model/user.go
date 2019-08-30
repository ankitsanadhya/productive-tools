package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name string
}

type Profile struct {
	gorm.Model
	Name      string
	User      User `gorm:"foreignkey:UserRefer"` // use UserRefer as foreign key
	UserRefer uint
}
