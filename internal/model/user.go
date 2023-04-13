package model

import (
	"errors"
	"github.com/linxbin/cron-service/global"
	"github.com/linxbin/cron-service/pkg/util"
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

func (u *User) FindByUsername(username string) *User {
	if err := global.DBEngine.Where("username = ?", username).First(&u).Error; err != nil {
		return nil
	}
	return u
}

func (u *User) Match(password string) (*User, error) {
	user := u.FindByUsername(u.Username)
	if user == nil {
		return nil, errors.New("user not exist")
	}

	if user.Password != encryptUserPassword(password, user.Salt) {
		return nil, errors.New("username or password not match")
	}
	return user, nil
}

// 密码加密
func encryptUserPassword(password, salt string) string {
	return util.EncodeMD5(password + salt)
}

func (u *User) Count() (int, error) {
	var count int
	if err := global.DBEngine.Model(&u).Where("is_del != ?", IsDelete).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (u *User) List(pageOffset, pageSize int) ([]*User, error) {
	var users []*User
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		global.DBEngine.Offset(pageOffset).Limit(pageSize)
	}

	if err = global.DBEngine.Where("is_del != ?", IsDelete).Order("id desc").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *User) Delete() error {
	return global.DBEngine.Where("id = ? AND is_del != ?", u.Model.ID, IsDelete).Delete(&u).Error
}

func (u *User) Create() (*User, error) {
	err := global.DBEngine.Create(u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}
