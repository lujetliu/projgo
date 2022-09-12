package middleware

import (
	"blog/pkg/app"
	"blog/pkg/errcode"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)

		// token存储在Query Param中
		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else {
			// token存储在Header中
			token = c.GetHeader("token")
		}

		response := app.NewResponse(c)

		if token == "" {
			ecode = errcode.InvalidParams
			response.ToErrorResponse(ecode)
			c.Abort()
			return
		}

		_, err := app.ParseToken(token)
		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				ecode = errcode.UnauthorizedTokenTimeout
			default:
				ecode = errcode.UnauthorizedTokenError
			}
		}

		if ecode != errcode.Success {
			response.ToErrorResponse(ecode)
			c.Abort()
			return
		}
		c.Next()
	}
}
