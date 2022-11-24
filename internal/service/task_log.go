package service

import (
	"github.com/linxbin/corn-service/internal/model"
	"github.com/linxbin/corn-service/pkg/app"
)

type TaskLogListRequest struct {
	TaskId uint32 `form:"task_id" binding:"required,gte=1"`
}

type TaskLogDetailRequest struct {
	Id uint32 `form:"id" binding:"required,gte=1"`
}

func (svc *Service) CountTaskLog(taskId uint32) (int, error) {
	return svc.dao.CountTaskLog(taskId)
}

func (svc *Service) TaskLogList(request *TaskLogListRequest, pager *app.Pager) ([]*model.TaskLog, error) {
	return svc.dao.TaskLogList(request.TaskId, pager.Page, pager.PageSize)
}

func (svc *Service) TaskLogDetail(id uint32) (model.TaskLog, error) {
	return svc.dao.TaskLogDetail(id)
}
