package routers

import (
	"blog/global"
	"blog/internal/middleware"
	"blog/internal/routers/api"
	v1 "blog/internal/routers/api/v1"
	"blog/pkg/limiter"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	// 中间件的注册在注册顺序上注意, Recovery 这类应用中间件应当尽可能的早注册
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}

	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(global.ServerSetting.ContextTimeout * time.Second))

	// 翻译
	r.Use(middleware.Translations())

	r.Static("/docs", "./docs")
	// r.Static("/static", "./storage/uploads")
	// TODO: StaticFS 的实现, 源码
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	url := ginSwagger.URL("http://127.0.0.1:8080/docs/swagger.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.POST("/auth", api.GetAuth)

	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)

	article := v1.NewArticle()
	tag := v1.NewTag()

	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		// tag
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		// article
		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}

	return r
}

// TODO: 限流器的设置可以放到配置中
var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
	Key:          "/auth",
	FillInterval: time.Second, // 每隔1秒放一次令牌
	Capacity:     10,          // 桶容量为10
	Quantum:      10,          // 每次放10个令牌
})
