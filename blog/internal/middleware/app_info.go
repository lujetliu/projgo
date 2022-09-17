package middleware

/*
	服务信息存储
	经常会需要在进程内上下文设置一些内部信息, 例如是应用名称和应用版本号
	这类基本信息, 也可以是业务属性的信息存储, 例如是根据不同的租户号获取
	不同的数据库实例对象, 这时候需要有一个统一的地方处理;

	 gin.Context 所提供的 setter 和 getter 主要用于以上场景, 在 gin 中
	 称为元数据管理(Metadata Management)

	func (c *Context) Set(key string, value interface{}) {
		if c.Keys == nil {
			c.Keys = make(map[string]interface{})
		}
		c.Keys[key] = value
	}

	func (c *Context) Get(key string) (value interface{}, exists bool) {
		value, exists = c.Keys[key]
		return
	}

*/

import "github.com/gin-gonic/gin"

func AppInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app_name", "blog-service")
		c.Set("app_version", "1.0.0")
		c.Next()
	}
}
