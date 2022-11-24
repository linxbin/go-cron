package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	*Model
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) GetOneByUsername(db *gorm.DB) error {
	if err := db.Where("username = ?", u.Username).First(&u).Error; err != nil {
		return err
	}
	return nil
}
