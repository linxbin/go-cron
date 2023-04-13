package cron

import (
	"errors"
	"fmt"
	"github.com/linxbin/cron-service/internal/utils"
	"golang.org/x/net/context"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/jakecoffman/cron"
	"github.com/linxbin/cron-service/global"
	"github.com/linxbin/cron-service/internal/model"
)

type Cron struct{}

// TaskCount 任务计数
type TaskCount struct {
	wg   sync.WaitGroup
	exit chan struct{}
}

func (tc *TaskCount) Add() {
	tc.wg.Add(1)
}

func (tc *TaskCount) Done() {
	tc.wg.Done()
}

func (tc *TaskCount) Exit() {
	tc.wg.Done()
	<-tc.exit
}

func (tc *TaskCount) Wait() {
	tc.Add()
	tc.wg.Wait()
	close(tc.exit)
}

func NewCron() *Cron {
	return &Cron{}
}

var (
	serviceCron *cron.Cron
	taskCount   TaskCount // 任务计数-正在运行的任务
)

func (c *Cron) Initialize() error {
	serviceCron = cron.New()
	serviceCron.Start()
	page := 1
	pageSize := 1000
	var t model.Task
	for {
		taskList, err := t.ActiveList(page, pageSize)
		if err != nil {
			return err
		}
		if len(taskList) == 0 {
			return nil
		}
		for _, item := range taskList {
			if err = c.AddTask(item); err != nil {
				return err
			}
		}
		page++
	}
}

func (c *Cron) AddTask(task *model.Task) error {
	taskFunc := createJob(task)
	if taskFunc == nil {
		return errors.New("创建任务处理Job失败")
	}

	cronName := strconv.Itoa(int(task.ID))
	err := panicToError(func() {
		serviceCron.AddFunc(task.Spec, taskFunc, cronName)
	})
	if err != nil {
		global.Logger.Errorf("添加任务到调度器失败: %s", err)
		return err
	}

	fmt.Printf(task.Name)
	return nil
}

func (c *Cron) RemoveTask(id int) {
	serviceCron.RemoveJob(strconv.Itoa(id))
}

// Run 直接运行任务
func (c *Cron) Run(task *model.Task) {
	go createJob(task)()
}

// RemoveAndAddTask 删除任务后添加
func (c *Cron) RemoveAndAddTask(task *model.Task) error {
	c.RemoveTask(int(task.ID))
	err := c.AddTask(task)
	if err != nil {
		return err
	}
	return nil
}

func createJob(task *model.Task) cron.FuncJob {
	taskFunc := func() {
		taskCount.Add()
		defer taskCount.Done()

		taskLogId, err := beforeExecJob(task)
		if err != nil || taskLogId <= 0 {
			return
		}
		global.Logger.Infof("开始执行任务#%s#命令-%s", task.Name, task.Command)
		result := execJob(task)
		global.Logger.Infof("任务完成#%s#命令-%s", task.Name, task.Command)
		if err = afterExecJob(taskLogId, result); err != nil {
			return
		}
	}

	return taskFunc
}

func beforeExecJob(task *model.Task) (taskLogId uint32, err error) {
	taskLogId, err = createTaskLog(task)
	if err != nil {
		return 0, err
	}

	return taskLogId, nil
}

func createTaskLog(task *model.Task) (uint32, error) {
	tl := model.TaskLog{
		TaskId:     task.ID,
		Name:       task.Name,
		Spec:       task.Spec,
		Command:    task.Command,
		Timeout:    task.Timeout,
		RetryTimes: task.RetryTimes,
		StartTime:  time.Now(),
		EndTime:    time.Now(),
		Status:     model.TaskLogStatusRunning,
	}
	err := tl.Create()
	if err != nil {
		return 0, err
	}
	return tl.ID, nil
}

func updateTaskLog(taskLogId uint32, result TaskResult) error {
	var status int
	if result.Err != nil {
		status = model.TaskLogStatusFailure
	} else {
		status = model.TaskLogStatusComplete
	}
	values := model.CommonMap{
		"status":      status,
		"end_time":    time.Now(),
		"retry_times": result.RetryTimes,
		"result":      result.Result,
	}
	tl := model.TaskLog{
		Model: &model.Model{ID: taskLogId},
	}
	return tl.Update(values)
}

// 任务执行后置操作
func afterExecJob(taskLogId uint32, result TaskResult) error {
	return updateTaskLog(taskLogId, result)
}

type TaskResult struct {
	Result     string
	Err        error
	RetryTimes uint8
}

// 执行具体任务
func execJob(task *model.Task) TaskResult {
	defer func() {
		if err := recover(); err != nil {
			global.Logger.Errorf("panic#service/task.go:execJob#%s", err)
		}
	}()
	// 默认只运行任务一次
	var execTimes uint8 = 1
	if task.RetryTimes > 0 {
		execTimes = task.RetryTimes
	}
	var i uint8 = 0
	var err error
	var output string
	timeout := time.Duration(task.Timeout) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	for i < execTimes {
		output, err = utils.ExecShell(ctx, task.Command)
		if err == nil {
			return TaskResult{Result: output, Err: err, RetryTimes: i}
		}
		i++
		if i < execTimes {
			global.Logger.Infof("任务执行失败#任务id-%d#重试第%d次#输出-%s#错误-%s", task.ID, i, output, err.Error())
			if task.RetryInterval > 0 {
				time.Sleep(time.Duration(task.RetryInterval) * time.Second)
			} else {
				// 默认重试间隔时间，每次递增1分钟
				time.Sleep(time.Duration(i) * time.Minute)
			}
		}
	}
	return TaskResult{Result: output, Err: err, RetryTimes: i}
}

// PanicToError Panic转换为error
func panicToError(f func()) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf(panicTrace(e))
		}
	}()
	f()
	return
}

// PanicTrace panic调用链跟踪
func panicTrace(err interface{}) string {
	stackBuf := make([]byte, 4096)
	n := runtime.Stack(stackBuf, false)

	return fmt.Sprintf("panic: %v %s", err, stackBuf[:n])
}
