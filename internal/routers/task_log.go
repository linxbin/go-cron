package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/linxbin/cron-service/internal/middleware"
	v1 "github.com/linxbin/cron-service/internal/routers/api/v1"
)

// InitTaskLogRouter 索引路由
func InitTaskLogRouter(Router *gin.RouterGroup) {
	taskLog := v1.NewTaskLog()
	router := Router.Group("task-log")
	router.Use(middleware.JWT())
	{
		router.GET("/list/:task_id", taskLog.List) // 任务日志列表
		router.GET("/detail/:id", taskLog.Detail)  // 任务日志详情
		router.POST("/clear", taskLog.Clear)       // 清空任务日志
	}
}
