package user

import (
	"errors"
	"github.com/linxbin/cron-service/internal/model"
	"github.com/linxbin/cron-service/pkg/app"
	"github.com/linxbin/cron-service/pkg/util"
	"math/rand"
	"time"
)

type IDRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type AddRequest struct {
	UserName string `form:"username" binding:"required,min=0,max=32"`
	Password string `form:"password" binding:"required,min=0,max=64"`
}

type LoginRequest struct {
	Username string `form:"username" binding:"required,min=0,max=32"`
	Password string `form:"password" binding:"required,min=0,max=255"`
}

type LoginInfo struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	Id       uint32 `json:"id"`
}

func Login(request *LoginRequest) (*LoginInfo, error) {
	u := &model.User{
		Username: request.Username,
	}
	user, err := u.Match(request.Password)
	if err != nil {
		return nil, err
	}
	token, err := app.GenerateToken(user.ID, user.Username)
	if err != nil {
		return nil, err
	}
	uf := &LoginInfo{
		Username: user.Username,
		Id:       user.ID,
		Token:    token,
	}
	return uf, nil
}

type Info struct {
	UserId   interface{} `json:"user_id"`
	Username interface{} `json:"username"`
	RoleId   string      `json:"role_id"`
}

func Count() (int, error) {
	u := new(model.User)
	return u.Count()
}

func List(pager *app.Pager) ([]*model.User, error) {
	u := &model.User{}
	return u.List(pager.Page, pager.PageSize)
}

func Delete(param *IDRequest) error {
	u := &model.User{
		Model: &model.Model{ID: param.ID},
	}
	err := u.Delete()
	if err != nil {
		return err
	}
	return nil
}

func Add(request *AddRequest) error {
	user := &model.User{}
	user = user.FindByUsername(request.UserName)
	if user != nil {
		return errors.New("用户名已存在")
	}
	salt := randomString(10)
	pwd := util.EncodeMD5(request.Password + salt)
	u := &model.User{
		Username: request.UserName,
		Password: pwd,
		Salt:     salt,
		Model:    &model.Model{Created: time.Now(), Updated: time.Now()},
	}
	_, err := u.Create()
	if err != nil {
		return err
	}

	return nil
}

func randomString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
