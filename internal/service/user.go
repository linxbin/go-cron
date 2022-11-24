package service

import (
	"github.com/gin-gonic/gin"
	"github.com/linxbin/corn-service/pkg/app"
)

type UserLoginRequest struct {
	Username string `form:"username" binding:"required,min=0,max=32"`
	Password string `form:"password" binding:"required,min=0,max=255"`
}

type LoginInfo struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	Id       uint32 `json:"id"`
}

func (svc *Service) Login(request *UserLoginRequest) (*LoginInfo, error) {
	user, err := svc.dao.MatchUser(request.Username, request.Password)
	if err != nil {
		return nil, err
	}
	token, err := app.GenerateToken(user.ID, user.Username)
	if err != nil {
		return nil, err
	}
	u := &LoginInfo{
		Username: user.Username,
		Id:       user.ID,
		Token:    token,
	}
	return u, nil
}

type UserInfo struct {
	UserId   interface{} `json:"user_id"`
	Username interface{} `json:"username"`
	RoleId   string      `json:"role_id"`
}

func (svc *Service) GetUserInfo(c *gin.Context) *UserInfo {
	userId, exist := c.Get("userId")
	if !exist {
		userId = nil
	}
	username, exist := c.Get("username")
	if !exist {
		username = nil
	}

	return &UserInfo{
		UserId:   userId,
		Username: username,
		RoleId:   "admin",
	}
}
