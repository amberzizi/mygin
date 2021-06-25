// @Title  ginlog.go
// @Description  zap日志创建，tools.LogerProducter()获取 logger 和 sugarlogger日志对象
// @Author  amberhu  20210625
// @Update
package tools

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var logger *zap.Logger
var sugarlogger *zap.SugaredLogger

//使用日志测试
func WriteLogGin(c *gin.Context){
	//InitLogger()
	//defer sugarlogger.Sync()
	//testlog1("https://www.sogo.com")
	//testlog2("https://www.sogo.d")
	//c.JSON(http.StatusOK, gin.H{
	//	"message": "Hello zaplog!",
	//})
}

//外部获取日志使用参数入口
func LogerProducter() (*zap.Logger,*zap.SugaredLogger){
	InitLogger()
	return logger, sugarlogger
}

//初始化日志
func InitLogger(){
	//用已有日志格式
	//logger, _ = zap.NewProduction()
	//sugarlogger = logger.Sugar()
	//自定义日志格式core
	core := zapcore.NewCore(getEncoder(), getLogWrite(),zapcore.InfoLevel)
	logger = zap.New(core)
	sugarlogger = logger.Sugar()

}

//logo写入文件  日志每日转储
func getLogWrite() zapcore.WriteSyncer {
	timeObj := time.Now()
	timestr := timeObj.Format("2006-01-02")
	file, _ := os.OpenFile("src/logs/"+timestr+".log",os.O_APPEND|os.O_CREATE|os.O_RDWR,0744)
	return zapcore.AddSync(file)
}
//log 格式 json
func getEncoder() zapcore.Encoder{
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}


//测试

//func testlog1(url string){
//	resp, err := http.Get(url)
//	if err != nil {
//		sugarlogger.Error("Error fetching url...",
//			zap.String("url",url),
//			zap.Error(err))
//	}else{
//		sugarlogger.Info("Success..",
//			zap.String("statusCode",resp.Status),
//			zap.String("url",url))
//		resp.Body.Close()
//	}
//}
//func testlog2(url string){
//	resp, err := http.Get(url)
//	if err != nil {
//		sugarlogger.Warn("Error fetching url...",
//			zap.String("url",url),
//			zap.Error(err))
//	}else{
//		sugarlogger.Info("Success..",
//			zap.String("statusCode",resp.Status),
//			zap.String("url",url))
//		resp.Body.Close()
//	}
//}