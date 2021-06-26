package settings

import (
	"fmt"
	"gopkg.in/gcfg.v1"
	"sync"
	"time"
)

//全局配置变量
var SettingGlb *Setting

//同步锁
var configLock = new(sync.RWMutex)

//配置文件结构体
type Setting struct {
	App struct {
		Name    string
		Mode    string
		Runport string
		Runhost string
	}
	Mysql struct {
		Host              string
		Dbname            string
		Username          string
		Password          string
		Port              string
		Maxconnection     int
		Maxidleconnection int
		Prefix            string
	}
	Redis struct {
		Host     string
		Password string
		Db       int
		Poolsize int
	}
	Log struct {
		Level      string
		Maxsize    int
		Maxage     int
		Maxbackups int
	}
}

//载入配置文件
func ReturnSetting() {
	settinginner := Setting{}
	err := gcfg.ReadFileInto(&settinginner, "src/conf/systeminfo.ini")
	if err != nil {
		fmt.Println("Failed to parse config file: %s", err)
	}
	configLock.Lock()
	SettingGlb = &settinginner
	configLock.Unlock()
}

//获取配置文件
func GetSetting() *Setting {
	configLock.RLock()
	defer configLock.RUnlock()
	return SettingGlb
}

//协程启动定时调用载入配置文件
func FreashSetting() {
	for {
		time.Sleep(10 * time.Second)
		ReturnSetting()
	}

}
