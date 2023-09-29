package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model `gorm:"embedded"`
	Email      string `gorm:"unique;not null"`
	Password   string `gorm:"not null"`
}

func (u *User) Table() string {
	return "user"
}
