package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/linxbin/cron-service/internal/middleware"
	v1 "github.com/linxbin/cron-service/internal/routers/api/v1"
)

// InitTaskRouter 索引路由
func InitTaskRouter(Router *gin.RouterGroup) {

	task := v1.NewTask()
	router := Router.Group("task")
	router.Use(middleware.JWT())
	{
		router.POST("/create", task.Create)    // 创建任务
		router.POST("/update", task.Update)    // 更新任务
		router.POST("/delete", task.Delete)    // 删除任务
		router.GET("/list", task.List)         // 任务列表
		router.GET("/detail/:id", task.Detail) // 任务详情
		router.POST("/enable", task.Enable)    // 开启任务
		router.POST("/disable", task.Disable)  // 关闭任务
	}
}
