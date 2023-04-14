package task

import (
	"github.com/linxbin/cron-service/global"
	cron2 "github.com/linxbin/cron-service/internal/cron"
	"github.com/linxbin/cron-service/internal/model"
	"github.com/linxbin/cron-service/pkg/app"
	"strings"
	"time"
)

type FormRequest struct {
	Name          string `form:"name" binding:"required,min=0,max=32"`
	Spec          string `form:"spec" binding:"required,min=0,max=64"`
	Command       string `form:"command" binding:"required,min=0,max=255"`
	Timeout       uint16 `form:"timeout" binding:"gte=0,lte=86400"`
	RetryTimes    uint8  `form:"retryTimes" binding:"gte=0"`
	RetryInterval uint16 `form:"retryInterval" binding:"gte=0"`
	Remark        string `form:"remark" binding:"min=0,max=255"`
	Status        uint8  `form:"status" binding:"oneof=10 20"`
}

type UpdateTaskRequest struct {
	IDRequest
	FormRequest
}

type IDRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type ListRequest struct {
	Name   string `form:"name" binding:"max=100"`
	Status uint8  `form:"status,default=0" binding:"oneof=0 10 20"`
}

type CountRequest struct {
	Name   string `form:"name" binding:"max=100"`
	Status uint8  `form:"status,default=1" binding:"oneof=0 1"`
}

func Create(request *FormRequest) error {
	t := model.Task{
		Name:          request.Name,
		Spec:          request.Spec,
		Command:       strings.TrimSpace(request.Command),
		Timeout:       request.Timeout,
		RetryTimes:    request.RetryTimes,
		RetryInterval: request.RetryInterval,
		Status:        request.Status,
		Remark:        request.Remark,
		Model:         &model.Model{Created: time.Now(), Updated: time.Now()},
	}
	task, err := t.Create()
	if err != nil {
		return err
	}

	cron := cron2.NewCron()
	return cron.AddTask(task)
}

func Update(request *UpdateTaskRequest) error {
	t := &model.Task{
		Model: &model.Model{ID: request.ID},
	}
	task, err := t.Detail(request.ID)
	if err != nil || task.ID == 0 {
		global.Logger.Errorf("svc.UpdateTask err: %v", err)
		return err
	}
	task.Name = request.Name
	task.Command = request.Command
	task.Status = request.Status
	task.RetryTimes = request.RetryTimes
	task.RetryInterval = request.RetryInterval
	task.Timeout = request.Timeout
	tx := global.DBEngine.Begin()
	if err = global.DBEngine.Save(&task).Error; err != nil {
		tx.Rollback()
		return err
	}

	cron := cron2.NewCron()
	if task.IsEnable() {
		err = cron.RemoveAndAddTask(task)
	} else {
		cron.RemoveTask(int(task.ID))
	}

	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func Count(request *CountRequest) (int, error) {
	t := &model.Task{
		Name:   request.Name,
		Status: request.Status,
	}
	return t.Count()
}

func List(request *ListRequest, pager *app.Pager) ([]*model.TaskList, error) {
	t := &model.Task{
		Name:   request.Name,
		Status: request.Status,
	}
	return t.List(pager.Page, pager.PageSize)
}

func Delete(param *IDRequest) error {
	t := &model.Task{
		Model: &model.Model{ID: param.ID},
	}
	err := t.Delete()
	if err != nil {
		return err
	}
	cron := cron2.NewCron()
	cron.RemoveTask(int(param.ID))
	return nil
}

func Detail(id uint32) (task *model.Task, err error) {
	t := &model.Task{
		Model: &model.Model{ID: id},
	}
	return t.Detail(id)
}

func Enable(id uint32) error {
	t := &model.Task{
		Model: &model.Model{ID: id},
	}
	tx := global.DBEngine.Begin()
	err := t.Enable()
	if err != nil {
		tx.Rollback()
		return err
	}
	err = addTaskToTimer(id)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// 添加任务到定时器
func addTaskToTimer(id uint32) error {
	t := &model.Task{}
	task, err := t.Detail(id)
	if err != nil {
		return err
	}
	cron := cron2.NewCron()
	err = cron.RemoveAndAddTask(task)
	if err != nil {
		return err
	}
	return nil
}

func Disable(id uint32) error {
	t := &model.Task{
		Model: &model.Model{
			ID: id,
		},
	}
	err := t.Disable()
	if err != nil {
		return err
	}
	cron := cron2.NewCron()
	cron.RemoveTask(int(id))
	return nil
}
