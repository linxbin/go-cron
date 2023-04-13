package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/linxbin/cron-service/internal/service/task"
	"github.com/linxbin/cron-service/pkg/app"
	"github.com/linxbin/cron-service/pkg/convert"
	"github.com/linxbin/cron-service/pkg/errcode"
)

type Task struct{}

func NewTask() Task {
	return Task{}
}

func (t *Task) Create(c *gin.Context) {
	param := task.FormRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	if err := task.Create(&param); err != nil {
		response.ToErrorResponse(errcode.ErrorCreateTaskFail)
		return
	}

	response.ToResponse(gin.H{})
}

func (t *Task) Update(c *gin.Context) {
	params := task.UpdateTaskRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	if err := task.Update(&params); err != nil {
		response.ToErrorResponse(errcode.ErrorUpdateTaskFail)
		return
	}

	response.ToResponse(gin.H{})
}

func (t *Task) List(c *gin.Context) {
	params := task.ListRequest{Name: convert.StrTo(c.Param("name")).String(), Status: uint8(convert.StrTo(c.Param("status")).MustInt())}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := task.Count(&task.CountRequest{Name: params.Name, Status: params.Status})
	if err != nil {
		response.ToErrorResponse(errcode.ErrorCountTaskFail)
		return
	}

	tags, err := task.List(&params, &pager)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorGetTaskListFail)
		return
	}

	response.ToResponseList(tags, totalRows)
}

func (t *Task) Delete(c *gin.Context) {
	params := task.IDRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	if err := task.Delete(&params); err != nil {
		response.ToErrorResponse(errcode.ErrorDeleteTaskFail)
		return
	}

	response.ToResponse(gin.H{})
}

func (t *Task) Detail(c *gin.Context) {
	params := task.IDRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	if t, err := task.Detail(params.ID); err != nil {
		response.ToErrorResponse(errcode.ErrorTaskNotFound)
	} else {
		response.ToResponse(gin.H{"data": t})
	}
}

func (t *Task) Enable(c *gin.Context) {
	params := task.IDRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	if err := task.Enable(params.ID); err != nil {
		response.ToErrorResponse(errcode.ErrorTaskEnable)
	} else {
		response.ToResponse(gin.H{})
	}
}

func (t *Task) Disable(c *gin.Context) {
	params := task.IDRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	if err := task.Disable(params.ID); err != nil {
		response.ToErrorResponse(errcode.ErrorTaskDisable)
	} else {
		response.ToResponse(gin.H{})
	}
}
