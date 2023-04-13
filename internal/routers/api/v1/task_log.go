package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/linxbin/cron-service/internal/service/tasklog"
	"github.com/linxbin/cron-service/pkg/app"
	"github.com/linxbin/cron-service/pkg/convert"
	"github.com/linxbin/cron-service/pkg/errcode"
)

type TaskLog struct{}

func NewTaskLog() TaskLog {
	return TaskLog{}
}

func (tl TaskLog) List(c *gin.Context) {
	params := tasklog.ListRequest{TaskId: convert.StrTo(c.Param("task_id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := tasklog.Count(params.TaskId)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorCountTaskFail)
		return
	}

	tags, err := tasklog.List(&params, &pager)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorTaskLogListFail)
		return
	}

	response.ToResponseList(tags, totalRows)

}

func (tl TaskLog) Detail(c *gin.Context) {
	params := tasklog.DetailRequest{Id: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	taskLog, err := tasklog.Detail(params.Id)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorTaskLogDetailFail)
		return
	}
	response.ToResponse(taskLog)
}

func (tl TaskLog) Clear(c *gin.Context) {
	params := tasklog.ListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	err := tasklog.Clear(params.TaskId)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorTaskLogDetailFail)
		return
	}
	response.ToResponse(gin.H{})
}
