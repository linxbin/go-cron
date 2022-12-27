package dao

import (
	"github.com/jinzhu/gorm"
	"strings"
	"time"

	"github.com/linxbin/cron-service/internal/model"
	"github.com/linxbin/cron-service/pkg/app"
)

type TaskForm struct {
	Id            uint32
	Name          string `binding:"Required;MaxSize(32)"`
	Spec          string
	Command       string `binding:"Required;MaxSize(256)"`
	Timeout       uint16 `binding:"Range(0,86400)"`
	RetryTimes    uint8
	RetryInterval uint16
	Remark        string
	Status        uint8 `binding:"oneof=0 1"`
}

func (d *Dao) CreateTask(form TaskForm) (*model.Task, error) {
	task := model.Task{
		Name:          form.Name,
		Spec:          form.Spec,
		Command:       strings.TrimSpace(form.Command),
		Timeout:       form.Timeout,
		RetryTimes:    form.RetryTimes,
		RetryInterval: form.RetryInterval,
		Status:        form.Status,
		Remark:        form.Remark,
		Model:         &model.Model{Created: time.Now(), Updated: time.Now()},
	}

	err := task.Create(d.engine)

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (d *Dao) UpdateTask(id uint32, form TaskForm) error {
	task := model.Task{
		Model: &model.Model{ID: id},
	}
	values := model.CommonMap{
		"name":           form.Name,
		"spec":           form.Spec,
		"command":        form.Command,
		"timeout":        form.Timeout,
		"retry_times":    form.RetryTimes,
		"retry_interval": form.RetryInterval,
		"remark":         form.Remark,
		"status":         form.Status,
	}

	return task.Update(d.engine, values)
}

func (d *Dao) CountTask(name string, status uint8) (int, error) {
	task := model.Task{
		Name:   name,
		Status: status,
	}
	return task.Count(d.engine)
}

func (d *Dao) TaskList(name string, status uint8, page, pageSize int) ([]*model.TaskList, error) {
	task := model.Task{Name: name, Status: status}
	pageOffset := app.GetPageOffset(page, pageSize)
	return task.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) TaskActiveList(page, pageSize int) ([]*model.Task, error) {
	task := model.Task{}
	pageOffset := app.GetPageOffset(page, pageSize)
	return task.ActiveList(d.engine, pageOffset, pageSize)
}

func (d *Dao) DeleteTask(id uint32) error {
	err := d.engine.Transaction(func(tx *gorm.DB) error {
		task := model.Task{Model: &model.Model{ID: id}}
		err := task.Delete(d.engine)
		if err != nil {
			return err
		}
		var taskLog model.TaskLog
		err = taskLog.BatchDelete(d.engine, id)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (d *Dao) TaskDetail(id uint32) (*model.Task, error) {
	var task model.Task
	return task.Detail(d.engine, id)
}

func (d *Dao) EnableTask(id uint32) error {
	task := model.Task{
		Model: &model.Model{ID: id},
	}
	values := model.CommonMap{
		"status": model.TaskStatusEnable,
	}
	return task.Update(d.engine, values)
}

func (d *Dao) DisableTask(id uint32) error {
	task := model.Task{
		Model: &model.Model{ID: id},
	}
	values := model.CommonMap{
		"status": model.TaskStatusDisable,
	}
	return task.Update(d.engine, values)
}
