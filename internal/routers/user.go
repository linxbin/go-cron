package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/linxbin/cron-service/internal/middleware"
	v1 "github.com/linxbin/cron-service/internal/routers/api/v1"
)

// InitUserRouter 索引路由
func InitUserRouter(Router *gin.RouterGroup) {
	user := v1.NewUser()

	router := Router.Group("user")
	router.Use(middleware.JWT())
	{
		router.GET("/info", user.Info) // 获取授权用户信息
	}
}
