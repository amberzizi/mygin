/*
* 测试协程
*
*/
package test

import (
	"github.com/gin-gonic/gin"
	"mygin/src/tools"
	"net/http"
)


func Sendgo(c *gin.Context){
	go say()
	println("hello world again ")
	println(tools.SettingGlb.Redis.Host)

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello sendinfo!",
	})
}

func say(){
	println("hello world via channel")
}
