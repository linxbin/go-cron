package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/linxbin/corn-service/pkg/app"
	"github.com/linxbin/corn-service/pkg/errcode"
	"strings"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			authorization string
			token         string
			errCode       = errcode.Success
		)
		if s, exist := c.GetQuery("Authorization"); exist {
			authorization = s
		} else {
			authorization = c.GetHeader("Authorization")
		}
		token = strings.Split(authorization, " ")[1]
		if token == "" {
			errCode = errcode.InvalidParams
		} else {
			claims, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					errCode = errcode.UnauthorizedTokenTimeout
				default:
					errCode = errcode.UnauthorizedTokenError
				}
			} else {
				c.Set("userId", claims.UserId)
				c.Set("username", claims.Username)
			}
		}

		if errCode != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(errCode)
			c.Abort()
			return
		}
		c.Next()
	}
}
