package user

import (
	"github.com/linxbin/cron-service/internal/model"
	"github.com/linxbin/cron-service/pkg/app"
)

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
