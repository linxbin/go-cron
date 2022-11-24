package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/linxbin/corn-service/internal/middleware"
	v1 "github.com/linxbin/corn-service/internal/routers/api/v1"
)

// InitTaskLogRouter 索引路由
func InitTaskLogRouter(Router *gin.RouterGroup) {
	taskLog := v1.NewTaskLog()
	router := Router.Group("task-log")
	router.Use(middleware.JWT())
	{
		router.GET("/list/:task_id", taskLog.List) // 任务列表
		router.GET("/detail/:id", taskLog.Detail)  // 任务详情
	}
}
