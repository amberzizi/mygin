package routers

import (
	"github.com/gin-gonic/gin"
	agentaction2 "mygin/src/application/controllers/agentaction"
	user2 "mygin/src/application/controllers/users"
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
	r.GET("/hellos", user2.Sendinfo)
	r.GET("/helloa", agentaction2.Sendinfo)
	r.GET("/hellosqlx", agentaction2.Sendsqlx)
	return r
}