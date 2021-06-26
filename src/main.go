package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"mygin/src/dao/mysql"
	"mygin/src/dao/redis"
	"mygin/src/routers"
	"mygin/src/settings"
	"mygin/src/tools"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	settings.ReturnSetting()
	//开启10s刷新配置协程 加载读取锁
	go settings.FreashSetting()
	//加载log
	tools.InitLogger()
	zap.L().Debug("logger init success...in main ")
	//加载redis初始化检查
	zap.L().Debug(redis.ReidsInitConnectParamInMain())
	//加载mysql初始化检查
	zap.L().Debug(mysql.MysqlInitConnectParamInMain())
	zap.L().Debug(mysql.MysqlGoroseInitConnectParamInMain())

	//载入路由
	r := routers.SetupRouter()

	//协程开机监听端口
	//优雅重启  和 supervisor不可兼得 supervisor会自动拉起监控中的关机进程
	srv := &http.Server{
		Addr:    settings.GetSetting().App.Runhost + ":" + settings.GetSetting().App.Runport,
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Debug(fmt.Sprint("server listen on..."+settings.GetSetting().App.Runport, err))
		}
	}()
	zap.L().Debug(fmt.Sprint("upppppp...", settings.GetSetting().App.Runport))

	//平滑优雅关机
	quit := make(chan os.Signal, 1) //创建一个接收信号的通道
	//kill 默认会发送 syscall.SIGTERM 信号  常用的ctrl+c就是触发这种信号
	//kill -2 发送 syscall.SIGINT 信号  添加后才能捕获
	//kill -9 发送 syscall.SIGKILL 信号，不能被捕获 不需添加
	//signal.Notify把接收到的 syscall.SIGINT 或 syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit
	zap.L().Debug("Shutdown server...")
	//创建一个5s超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//5秒内优雅关闭服务器
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Debug(fmt.Sprint("Shutdown server error...", err))
	}
	zap.L().Debug("server exiting...bye")
}
