package tools

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"os"
)

var logger *zap.Logger
var sugarlogger *zap.SugaredLogger


func WriteLogGin(c *gin.Context){
	InitLogger()
	defer logger.Sync()
	testlog1("https://www.sogo.com")
	testlog1("https://www.sogo.d")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello zaplog!",
	})
}

func InitLogger(){
	//logger, _ = zap.NewProduction()
	//sugarlogger = logger.Sugar()

	core := zapcore.NewCore(getEncoder(),getLogWrite(),zapcore.DebugLevel)
	logger = zap.New(core)
	sugarlogger = logger.Sugar()

}

func testlog1(url string){
	resp, err := http.Get(url)
	if err != nil {
		sugarlogger.Error("Error fetching url...",
			zap.String("url",url),
			zap.Error(err))
	}else{
		sugarlogger.Info("Success..",
			zap.String("statusCode",resp.Status),
			zap.String("url",url))
		resp.Body.Close()
	}
}
func testlog2(url string){
	resp, err := http.Get(url)
	if err != nil {
		sugarlogger.Error("Error fetching url...",
			zap.String("url",url),
			zap.Error(err))
	}else{
		sugarlogger.Info("Success..",
			zap.String("statusCode",resp.Status),
			zap.String("url",url))
		resp.Body.Close()
	}
}
func getLogWrite() zapcore.WriteSyncer {
	file, _ := os.OpenFile("src/logs/test.log",os.O_APPEND|os.O_CREATE|os.O_RDWR,0744)
	return zapcore.AddSync(file)
}
func getEncoder() zapcore.Encoder{
	return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
}