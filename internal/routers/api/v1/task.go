package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/linxbin/cron-service/global"
	"github.com/linxbin/cron-service/internal/service"
	"github.com/linxbin/cron-service/pkg/app"
	"github.com/linxbin/cron-service/pkg/convert"
	"github.com/linxbin/cron-service/pkg/errcode"
)

type Task struct{}

func NewTask() Task {
	return Task{}
}

func (t *Task) Create(c *gin.Context) {
	param := service.TaskFormRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	if err := svc.CreateTask(&param); err != nil {
		global.Logger.Errorf("svc.CreateTask err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateTaskFail)
		return
	}

	response.ToResponse(gin.H{})
}

func (t *Task) Update(c *gin.Context) {
	params := service.UpDateTaskRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	if err := svc.UpdateTask(&params); err != nil {
		global.Logger.Errorf("svc.UpdateTask err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTaskFail)
		return
	}

	response.ToResponse(gin.H{})
}

func (t *Task) List(c *gin.Context) {
	params := service.TaskListRequest{Name: convert.StrTo(c.Param("name")).String(), Status: uint8(convert.StrTo(c.Param("status")).MustInt())}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountTask(&service.CountTaskRequest{Name: params.Name, Status: params.Status})
	if err != nil {
		global.Logger.Errorf("svc.CountTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountTaskFail)
		return
	}

	tags, err := svc.TaskList(&params, &pager)
	if err != nil {
		global.Logger.Errorf("svc.GetTagList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTaskListFail)
		return
	}

	response.ToResponseList(tags, totalRows)
}

func (t *Task) Delete(c *gin.Context) {
	params := service.TaskRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	if err := svc.DeleteTask(&params); err != nil {
		global.Logger.Errorf("svc.UpdateTask err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteTaskFail)
		return
	}

	response.ToResponse(gin.H{})

}

func (t *Task) Detail(c *gin.Context) {
	params := service.TaskRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	if task, err := svc.TaskDetail(params.ID); err != nil {
		global.Logger.Errorf("svc.TaskDetail err: %v", err)
		response.ToErrorResponse(errcode.ErrorTaskNotFound)
	} else {
		response.ToResponse(gin.H{"data": task})
	}
}

func (t *Task) Enable(c *gin.Context) {
	params := service.TaskRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	if err := svc.EnableTask(params.ID); err != nil {
		global.Logger.Errorf("svc.EnableTask err: %v", err)
		response.ToErrorResponse(errcode.ErrorTaskEnable)
	} else {
		response.ToResponse(gin.H{})
	}
}

func (t *Task) Disable(c *gin.Context) {
	params := service.TaskRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	if err := svc.DisableTask(params.ID); err != nil {
		global.Logger.Errorf("svc.DisableTask err: %v", err)
		response.ToErrorResponse(errcode.ErrorTaskDisable)
	} else {
		response.ToResponse(gin.H{})
	}
}
