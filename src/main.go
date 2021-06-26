package main

import (
	"fmt"
	"go.uber.org/zap"
	"mygin/src/routers"
	"mygin/src/tools"
)

func main() {
	tools.ReturnSetting()
	//开启10s刷新配置协程 加载读取锁
	go tools.FreashSetting()
	//加载log
	tools.InitLogger()
	zap.L().Debug("logger init success...in main ")
	//加载redis初始化检查
	zap.L().Debug(tools.ReidsInitConnectParamInMain())
	//加载mysql初始化检查
	zap.L().Debug(tools.MysqlInitConnectParamInMain())

	r := routers.SetupRouter()
	if err := r.Run(tools.GetSetting().App.Runhost+":"+tools.GetSetting().App.Runport); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}



