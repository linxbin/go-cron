package model

import (
	"github.com/jinzhu/gorm"
	"github.com/linxbin/cron-service/global"
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

func (tg *TaskLog) Create() error {
	return global.DBEngine.Create(&tg).Error
}

func (tg *TaskLog) Update(values interface{}) error {
	if err := global.DBEngine.Model(tg).Where("id = ? AND is_del != ?", tg.ID, IsDelete).Updates(values).Error; err != nil {
		return err
	}
	return nil
}

func (tg *TaskLog) Count() (int, error) {
	var count int
	if err := global.DBEngine.Model(&tg).Where("task_id = ? and is_del != ?", tg.TaskId, IsDelete).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (tg *TaskLog) List(pageOffset, pageSize int) ([]*TaskLog, error) {
	var taskLogs []*TaskLog
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		global.DBEngine.Offset(pageOffset).Limit(pageSize)
	}
	if err = global.DBEngine.Where("task_id = ? and is_del != ?", tg.TaskId, IsDelete).Order("id desc").Find(&taskLogs).Error; err != nil {
		return nil, err
	}

	return taskLogs, nil
}

func (tg *TaskLog) Clear() error {
	return global.DBEngine.Where("task_id = ? and is_del != ?", tg.TaskId, IsDelete).Delete(&TaskLog{}).Error
}

func (tg *TaskLog) Detail() (*TaskLog, error) {
	var err error

	if err = global.DBEngine.First(tg, "id = ? and is_del != ?", tg.ID, IsDelete).Error; err != nil {
		return tg, err
	}
	return tg, nil
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
