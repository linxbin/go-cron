package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/linxbin/cron-service/internal/routers/api/v1"
)

// InitWithoutAuthRouter 索引路由
func InitWithoutAuthRouter(Router *gin.RouterGroup) {
	user := v1.NewUser()

	router := Router.Group("")
	router.Use()
	{
		router.POST("/login", user.Login) // 登录
	}
}
