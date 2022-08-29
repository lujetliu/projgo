package main

/*
	HTTP 301: 永久重定向
		表示被请求的资源已永久移动到新位置, 即我们常说的301跳转, 并且将来任何
		对此资源的引用都应该使用本响应返回的URI;
	HTTP 307: 临时重定向
		表示请求的资源现在临时从不同的URI响应请求; 由于这样的重定向是临时的,
		客户端应当继续向原有地址发送以后的请求;


*/

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.Run()
}
