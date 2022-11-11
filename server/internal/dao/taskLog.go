package dao

import (
	"github.com/linxbin/corn-service/internal/model"
	"github.com/linxbin/corn-service/pkg/app"
	"strings"
	"time"
)

type TaskLogForm struct {
	Id         uint32
	TaskId     uint32
	Name       string `binding:"Required;MaxSize(32)"`
	Spec       string
	Command    string `binding:"Required;MaxSize(256)"`
	Timeout    uint16 `binding:"Range(0,86400)"`
	RetryTimes uint8
	Status     uint8 `binding:"oneof=0 1 2"`
	StartTime  time.Time
	EndTime    time.Time
	Result     string
}

func (d *Dao) CreateTaskLog(form TaskLogForm) (uint32, error) {
	taskLog := model.TaskLog{
		TaskId:     form.TaskId,
		Name:       form.Name,
		Spec:       form.Spec,
		Command:    strings.TrimSpace(form.Command),
		Timeout:    form.Timeout,
		RetryTimes: form.RetryTimes,
		Status:     form.Status,
		StartTime:  form.StartTime,
		EndTime:    form.EndTime,
		Result:     form.Result,
		Model:      &model.Model{Created: time.Now(), Updated: time.Now()},
	}

	err := taskLog.Create(d.engine)
	if err != nil {
		return 0, err
	}

	return taskLog.ID, nil
}

func (d *Dao) UpdateTaskLog(id uint32, commonMap model.CommonMap) error {
	taskLog := model.TaskLog{
		Model: &model.Model{ID: id},
	}

	return taskLog.Update(d.engine, commonMap)
}

func (d *Dao) CountTaskLog(taskId uint32) (int, error) {
	taskLog := model.TaskLog{
		TaskId: taskId,
	}
	return taskLog.Count(d.engine, taskId)
}

func (d *Dao) TaskLogList(taskId uint32, page, pageSize int) ([]*model.TaskLog, error) {
	taskLog := model.TaskLog{TaskId: taskId}
	pageOffset := app.GetPageOffset(page, pageSize)
	return taskLog.List(d.engine, taskId, pageOffset, pageSize)
}

func (d *Dao) TaskLogDetail(id uint32) (model.TaskLog, error) {
	var taskLog model.TaskLog
	return taskLog.Detail(d.engine, id)
}
