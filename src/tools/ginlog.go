// @Title  ginlog.go
// @Description  zap日志创建，tools.LogerProducter()获取 logger 和 sugarlogger日志对象
// @Author  amberhu  20210625
// @Update
package tools

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"mygin/src/settings"
	"time"
)

var logger *zap.Logger
var sugarlogger *zap.SugaredLogger

//type logsetting struct {
//	Log struct{
//		Level  string
//		Maxsize int
//		Maxage int
//		Maxbackups int
//	}
//}

//外部获取日志使用参数入口
func LogerProducter() (*zap.Logger, *zap.SugaredLogger) {
	InitLogger()
	return logger, sugarlogger
}

//初始化日志
func InitLogger() {
	//用已有日志格式
	//logger, _ = zap.NewProduction()
	//sugarlogger = logger.Sugar()
	//自定义日志格式core

	//core := zapcore.NewCore(getEncoder(), getLogWrite(),zapcore.InfoLevel)
	logsetting := settings.GetSetting() //配置文件
	var l = new(zapcore.Level)
	err := l.UnmarshalText([]byte(logsetting.Log.Level))
	if err != nil {
		return
	}
	core := zapcore.NewCore(getEncoder(), getLogWrite(settings.GetSetting().Log.Maxsize,
		settings.GetSetting().Log.Maxbackups, settings.GetSetting().Log.Maxage, false), l)
	//zap.AddCaller()增加函数调用信息
	logger = zap.New(core, zap.AddCaller())
	sugarlogger = logger.Sugar()
	//替换全局logger
	zap.ReplaceGlobals(logger)
	fmt.Printf("Log try init success,use zap.L().debug(...)\n")
}

//logo写入文件  日志每日转储 支持切割
func getLogWrite(maxsize int, maxbackups int, maxage int, ifcompress bool) zapcore.WriteSyncer {
	timeObj := time.Now()
	timestr := timeObj.Format("2006-01-02")
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "src/logs/" + timestr + ".log",
		MaxSize:    maxsize,    //M
		MaxBackups: maxbackups, //备份数量
		MaxAge:     maxage,     //最多保存多少天
		Compress:   ifcompress, //是否压缩

	}
	//file, _ := os.OpenFile("src/logs/"+timestr+".log",os.O_APPEND|os.O_CREATE|os.O_RDWR,0744)
	return zapcore.AddSync(lumberJackLogger)
}

//log 格式 json
func getEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

func Printetest() {
	println(settings.GetSetting().Redis.Host)
}
