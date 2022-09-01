package main

/*
	HTTP 301: 永久重定向
		表示被请求的资源已永久移动到新位置, 即我们常说的301跳转, 并且将来任何
		对此资源的引用都应该使用本响应返回的URI;
	HTTP 307: 临时重定向
		表示请求的资源现在临时从不同的URI响应请求; 由于这样的重定向是临时的,
		客户端应当继续向原有地址发送以后的请求;


*/

import (
	"blog/global"
	"blog/internal/model"
	"blog/internal/routers"
	"blog/pkg/logger"
	"blog/pkg/setting"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}

}

// @title 博客系统
// @version 1.0
// @description go 项目实战
// @termsOfService  git@github.com:lujetliu/projgo.git
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout * time.Second,
		WriteTimeout:   global.ServerSetting.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 10, // 20MB
	}
	s.ListenAndServe()
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		// TODO: 熟悉 lumberjack 库的使用
		// 结构体添加了 yaml 标签, 支持从配置文件解析以下参数
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600, // 日志文件允许的最大占用空间为 600 MB TODO: 如果超过了是如何处理的
		MaxAge:    10,  // 日志文件最大生存周期为10天 TODO: 如何监控文件过期的
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}
