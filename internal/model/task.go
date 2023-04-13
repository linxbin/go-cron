package model

import (
	"github.com/linxbin/cron-service/global"
)

type Task struct {
	*Model
	Name          string `json:"name"`
	Spec          string `json:"spec"`
	Command       string `json:"command"`
	Timeout       uint16 `json:"timeout"`
	RetryTimes    uint8  `json:"retry_times"`
	RetryInterval uint16 `json:"retry_interval"`
	Remark        string `json:"remark"`
	Status        uint8  `json:"status"`
}

type TaskList struct {
	*Task
	IsEnable bool `json:"is_enable"`
}

const (
	TaskStatusEnable  = 10
	TaskStatusDisable = 20
)

func (t *Task) TableName() string {
	return "task"
}

func (t *Task) Create() (*Task, error) {
	err := global.DBEngine.Create(t).Error

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (t *Task) Update(values interface{}) (*Task, error) {
	if err := global.DBEngine.Model(t).Where("is_del != ?", IsDelete).Updates(values).Error; err != nil {
		return t, err
	}
	return t, nil
}

func (t *Task) Delete() error {
	return global.DBEngine.Where("id = ? AND is_del != ?", t.Model.ID, IsDelete).Delete(&t).Error
}

func (t *Task) Count() (int, error) {
	var count int
	query := global.DBEngine.Model(&t).Where("is_del != ?", IsDelete)
	if t.Name != "" {
		query.Where("name = ?", t.Name)
	}
	if t.Status != 0 {
		query.Where("status = ?", t.Status)
	}
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (t *Task) List(pageOffset, pageSize int) ([]*TaskList, error) {
	var tasks []*TaskList
	var err error
	query := global.DBEngine.Where("is_del != ?", IsDelete)
	if pageOffset >= 0 && pageSize > 0 {
		query.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		query.Where("name = ?", t.Name)
	}
	if t.Status != 0 {
		query.Where("status = ?", t.Status)
	}

	if err = query.Order("id desc").Find(&tasks).Error; err != nil {
		return nil, err
	}

	for i, item := range tasks {
		tasks[i].IsEnable = item.Status == TaskStatusEnable
	}

	return tasks, nil
}

func (t *Task) Detail(ID uint32) (*Task, error) {
	var err error

	if err = global.DBEngine.First(t, "id = ? and is_del != ?", ID, IsDelete).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (t *Task) ActiveList(pageOffset, pageSize int) ([]*Task, error) {
	var tasks []*Task
	var err error
	query := global.DBEngine.Where("status = ? and is_del != ?", TaskStatusEnable, IsDelete)
	if pageOffset >= 0 && pageSize > 0 {
		query.Offset(pageOffset).Limit(pageSize)
	}

	if err = query.Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t *Task) Enable() error {
	values := CommonMap{
		"status": TaskStatusEnable,
	}
	_, err := t.Update(values)
	if err != nil {
		return err
	}
	return nil
}

func (t *Task) Disable() error {
	values := CommonMap{
		"status": TaskStatusDisable,
	}
	_, err := t.Update(values)
	if err != nil {
		return err
	}
	return nil
}
