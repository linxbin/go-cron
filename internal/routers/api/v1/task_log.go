package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/linxbin/corn-service/global"
	"github.com/linxbin/corn-service/internal/service"
	"github.com/linxbin/corn-service/pkg/app"
	"github.com/linxbin/corn-service/pkg/convert"
	"github.com/linxbin/corn-service/pkg/errcode"
)

type TaskLog struct{}

func NewTaskLog() TaskLog {
	return TaskLog{}
}

func (tl TaskLog) List(c *gin.Context) {
	params := service.TaskLogListRequest{TaskId: convert.StrTo(c.Param("task_id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountTaskLog(params.TaskId)
	if err != nil {
		global.Logger.Errorf("svc.CountTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountTaskFail)
		return
	}

	tags, err := svc.TaskLogList(&params, &pager)
	if err != nil {
		global.Logger.Errorf("svc.GetTagLogList err: %v", err)
		response.ToErrorResponse(errcode.ErrorTaskLogListFail)
		return
	}

	response.ToResponseList(tags, totalRows)

}

func (tl TaskLog) Detail(c *gin.Context) {
	params := service.TaskLogDetailRequest{Id: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	taskLog, err := svc.TaskLogDetail(params.Id)
	if err != nil {
		global.Logger.Errorf("svc.GetTaskLogDetail err: %v", err)
		response.ToErrorResponse(errcode.ErrorTaskLogDetailFail)
		return
	}
	response.ToResponse(taskLog)
}
