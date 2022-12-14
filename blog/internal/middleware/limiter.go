package middleware

import (
	"blog/pkg/app"
	"blog/pkg/errcode"
	"blog/pkg/limiter"

	"github.com/gin-gonic/gin"
)

func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			/*
				TakeAvailable 方法, 会占用存储桶中立即可用的令牌的数量,
				返回值为删除的令牌数, 如果没有可用的令牌, 将会返回 0,
				也就是已经超出配额了, 因此这时将返回 errcode.TooManyRequest
				状态告诉客户端需要减缓并控制请求速度;
			*/
			if count == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
