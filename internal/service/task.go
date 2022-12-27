package service

import (
	"github.com/linxbin/cron-service/global"
	"github.com/linxbin/cron-service/internal/dao"
	"github.com/linxbin/cron-service/internal/model"
	"github.com/linxbin/cron-service/pkg/app"
)

type TaskFormRequest struct {
	Name          string `form:"name" binding:"required,min=0,max=32"`
	Spec          string `form:"spec" binding:"required,min=0,max=64"`
	Command       string `form:"command" binding:"required,min=0,max=255"`
	Timeout       uint16 `form:"timeout" binding:"required,gte=1,lte=86400"`
	RetryTimes    uint8  `form:"retryTimes" binding:"required,gte=0"`
	RetryInterval uint16 `form:"retryInterval" binding:"required,gte=1"`
	Remark        string `form:"remark" binding:"min=0,max=255"`
	Status        uint8  `form:"status" binding:"oneof=10 20"`
}

type UpDateTaskRequest struct {
	ID uint32 `form:"id" binding:"required"`
	TaskFormRequest
}

type TaskRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type TaskListRequest struct {
	Name   string `form:"name" binding:"max=100"`
	Status uint8  `form:"status,default=0" binding:"oneof=0 10 20"`
}

type CountTaskRequest struct {
	Name   string `form:"name" binding:"max=100"`
	Status uint8  `form:"status,default=1" binding:"oneof=0 1"`
}

func (svc *Service) CreateTask(request *TaskFormRequest) error {
	form := dao.TaskForm{
		Name:          request.Name,
		Spec:          request.Spec,
		Command:       request.Command,
		Timeout:       request.Timeout,
		RetryTimes:    request.RetryTimes,
		RetryInterval: request.RetryInterval,
		Remark:        request.Remark,
		Status:        request.Status,
	}

	task, err := svc.dao.CreateTask(form)
	if err != nil {
		return err
	}

	return svc.cron.AddTask(task)
}

func (svc *Service) UpdateTask(request *UpDateTaskRequest) error {
	task, err := svc.TaskDetail(request.ID)
	if err != nil || task.ID == 0 {
		global.Logger.Errorf("svc.UpdateTask err: %v", err)
		return err
	}
	form := dao.TaskForm{
		Name:          request.Name,
		Spec:          request.Spec,
		Command:       request.Command,
		Timeout:       request.Timeout,
		RetryTimes:    request.RetryTimes,
		RetryInterval: request.RetryInterval,
		Remark:        request.Remark,
		Status:        request.Status,
	}

	return svc.dao.UpdateTask(request.ID, form)
}

func (svc *Service) CountTask(request *CountTaskRequest) (int, error) {
	return svc.dao.CountTask(request.Name, request.Status)
}

func (svc *Service) TaskList(request *TaskListRequest, pager *app.Pager) ([]*model.TaskList, error) {
	return svc.dao.TaskList(request.Name, request.Status, pager.Page, pager.PageSize)
}

func (svc *Service) DeleteTask(param *TaskRequest) error {
	err := svc.dao.DeleteTask(param.ID)
	if err != nil {
		return err
	}
	svc.cron.RemoveTask(int(param.ID))
	return nil
}

func (svc *Service) TaskDetail(id uint32) (task *model.Task, err error) {
	return svc.dao.TaskDetail(id)
}

func (svc *Service) EnableTask(id uint32) error {
	err := svc.dao.EnableTask(id)
	if err != nil {
		return err
	}
	err = svc.addTaskToTimer(id)
	if err != nil {
		return err
	}
	return nil
}

// 添加任务到定时器
func (svc *Service) addTaskToTimer(id uint32) error {
	task, err := svc.dao.TaskDetail(id)
	if err != nil {
		return err
	}
	err = svc.cron.RemoveAndAddTask(task)
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) DisableTask(id uint32) error {
	err := svc.dao.DisableTask(id)
	if err != nil {
		return err
	}
	svc.cron.RemoveTask(int(id))
	return nil
}
