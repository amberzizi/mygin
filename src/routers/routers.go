package routers

import (
	"github.com/gin-gonic/gin"
	"mygin/src/controllers/agentaction"
	"mygin/src/controllers/tools"
	user "mygin/src/controllers/users"
	"net/http"
)

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello q1mi!",
	})
}

// SetupRouter 配置路由信息
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/hello", helloHandler)
	r.GET("/hellos", user.Sendinfo)
	r.GET("/helloa",agentaction.Sendinfo)
	r.GET("/hellosqlx",agentaction.Sendsqlx)
	r.GET("/hellozap",tools.WriteLogGin)
	return r
}