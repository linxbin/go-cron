package tasklog

import (
	"github.com/linxbin/cron-service/internal/model"
	"github.com/linxbin/cron-service/pkg/app"
)

type ListRequest struct {
	TaskId uint32 `form:"task_id" binding:"required,gte=1"`
}

type DetailRequest struct {
	Id uint32 `form:"id" binding:"required,gte=1"`
}

func Count(taskId uint32) (int, error) {
	tl := model.TaskLog{
		TaskId: taskId,
	}
	return tl.Count()
}

func List(request *ListRequest, pager *app.Pager) ([]*model.TaskLog, error) {
	tl := model.TaskLog{
		TaskId: request.TaskId,
	}
	return tl.List(pager.Page, pager.PageSize)
}

func Detail(id uint32) (*model.TaskLog, error) {
	tl := model.TaskLog{
		Model: &model.Model{ID: id},
	}
	return tl.Detail()
}

func Clear(taskId uint32) error {
	tl := model.TaskLog{
		TaskId: taskId,
	}
	return tl.Clear()
}
