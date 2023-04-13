package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/linxbin/cron-service/global"
	"github.com/linxbin/cron-service/internal/service/user"
	"github.com/linxbin/cron-service/pkg/app"
	"github.com/linxbin/cron-service/pkg/errcode"
)

type User struct{}

func NewUser() User {
	return User{}
}

func (u *User) Login(c *gin.Context) {
	param := user.LoginRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	userInfo, err := user.Login(&param)
	if err != nil {
		global.Logger.Errorf("svc.user login err: %v", err)
		response.ToErrorResponse(errcode.LoginError)
		return
	}

	response.ToResponse(userInfo)
}

func (u *User) Info(c *gin.Context) {
	userId, exist := c.Get("userId")
	if !exist {
		userId = nil
	}
	username, exist := c.Get("username")
	if !exist {
		username = nil
	}
	userInfo := &user.Info{
		UserId:   userId,
		Username: username,
		RoleId:   "admin",
	}
	response := app.NewResponse(c)
	response.ToResponse(userInfo)
}

func (u *User) Add(c *gin.Context) {
	param := user.AddRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	if err := user.Add(&param); err != nil {
		global.Logger.Errorf("svc.CreateTask err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateTaskFail)
		return
	}

	response.ToResponse(gin.H{})
}

func (u *User) List(c *gin.Context) {
	response := app.NewResponse(c)
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := user.Count()
	if err != nil {
		response.ToErrorResponse(errcode.ErrorCountTaskFail)
		return
	}

	tags, err := user.List(&pager)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorTaskLogListFail)
		return
	}

	response.ToResponseList(tags, totalRows)
}

func (u *User) Delete(c *gin.Context) {
	params := user.IDRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	if err := user.Delete(&params); err != nil {
		response.ToErrorResponse(errcode.ErrorDeleteTaskFail)
		return
	}

	response.ToResponse(gin.H{})
}
