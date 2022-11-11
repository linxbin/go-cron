package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

const (
	TaskLogStatusRunning  = 0 // 执行中
	TaskLogStatusComplete = 1 // 完成
	TaskLogStatusFailure  = 2 // 失败
)

type TaskLog struct {
	*Model
	TaskId     uint32    `json:"task_id"`
	Name       string    `json:"name"`
	Spec       string    `json:"spec"`
	Command    string    `json:"command"`
	Timeout    uint16    `json:"timeout"`
	RetryTimes uint8     `json:"retry_times"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	Result     string    `json:"result"`
	Status     uint8     `json:"status"`
}

func (tg *TaskLog) TableName() string {
	return "task_log"
}

func (tg *TaskLog) Create(db *gorm.DB) error {
	return db.Create(&tg).Error
}

func (tg *TaskLog) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(tg).Where("id = ? AND is_del != ?", tg.ID, IsDelete).Updates(values).Error; err != nil {
		return err
	}
	return nil
}

func (tg *TaskLog) Count(db *gorm.DB, taskId uint32) (int, error) {
	var count int
	if err := db.Model(&tg).Where("task_id = ? and is_del != ?", taskId, IsDelete).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (tg *TaskLog) List(db *gorm.DB, taskId uint32, pageOffset, pageSize int) ([]*TaskLog, error) {
	var taskLogs []*TaskLog
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if err = db.Where("task_id = ? and is_del != ?", taskId, IsDelete).Order("id desc").Find(&taskLogs).Error; err != nil {
		return nil, err
	}

	return taskLogs, nil
}

func (tg *TaskLog) Detail(db *gorm.DB, ID uint32) (TaskLog, error) {
	taskLog := TaskLog{}
	var err error

	if err = db.First(&taskLog, "id = ? and is_del != ?", ID, IsDelete).Error; err != nil {
		return taskLog, err
	}
	return taskLog, nil
}

func (tg *TaskLog) BatchDelete(db *gorm.DB, taskId uint32) error {
	var values map[string]int
	values = make(map[string]int)
	values["is_del"] = IsDelete
	if err := db.Model(tg).Where("task_id = ? AND is_del != ?", taskId, IsDelete).Updates(values).Error; err != nil {
		return err
	}
	return nil
}
