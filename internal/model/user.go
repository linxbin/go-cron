package model

import (
	"errors"
	"fmt"
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

func FindByUsername(username string) (*User, error) {
	user := &User{}
	if err := global.DBEngine.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) Match(password string) (*User, error) {
	user, err := FindByUsername(u.Username)
	fmt.Println(user)
	if err != nil {
		return nil, err
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
