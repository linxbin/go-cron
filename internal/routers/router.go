package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/linxbin/cron-service/pkg/app"
	"github.com/linxbin/cron-service/pkg/errcode"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(Cors()) // 默认跨域
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.NoRoute(func(context *gin.Context) {
		response := app.NewResponse(context)
		response.ToErrorResponse(errcode.NotFound)
	})

	group := r.Group("/api/v1")
	{
		InitWithoutAuthRouter(group)
		InitTaskRouter(group)    // 任务管理
		InitTaskLogRouter(group) // 任务日志管理
		InitUserRouter(group)    // 用户管理
	}

	return r
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("content-type", "application/json")
		}
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
