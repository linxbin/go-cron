package dao

import (
	"errors"
	"github.com/linxbin/corn-service/internal/model"
	"github.com/linxbin/corn-service/pkg/util"
)

func (d *Dao) MatchUser(username, password string) (*model.User, error) {
	user := &model.User{
		Username: username,
	}
	if err := user.GetOneByUsername(d.engine); err != nil {
		return nil, err
	}
	if user.Password != d.encryptUserPassword(password, user.Salt) {
		return nil, errors.New("username or password not match")
	}
	return user, nil
}

// 密码加密
func (d *Dao) encryptUserPassword(password, salt string) string {
	return util.EncodeMD5(password + salt)
}
