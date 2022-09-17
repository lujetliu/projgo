package middleware

import (
	"blog/global"
	"blog/pkg/logger"
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// 双写
func (w AccessLogWriter) Writer(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = bodyWriter // 由此在写入响应的时候获取到了响应体

		beginTime := time.Now().Unix()
		c.Next() // TODO: Next 后的代码在何时执行?

		endTime := time.Now().Unix()

		fields := logger.Fields{
			"request": c.Request.PostForm.Encode(),
			"reponse": bodyWriter.body.String(),
		}

		global.Logger.WithFields(fields).Info(`access log: method : %s, 
			status_code: %d, begin_time: %d, end_time: %d`,
			c.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime,
		)
	}
}
