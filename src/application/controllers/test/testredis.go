package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"mygin/src/tools"
	"net/http"
	"time"
)

func Sendredis(c *gin.Context){

	//测试redis
	rdb := tools.ReturnRedisDb()
	defer rdb.Close()
	//测试watch
	key := "watch_count"
	errw := rdb.Watch(func(tx *redis.Tx) error{
		n,err := tx.Get(key).Int()
		if err != nil && err != redis.Nil {
			fmt.Printf("try connecting fail,err:%v\n",err)
			return err
		}
		println(n)
		time.Sleep(time.Second*10)
		pipe := tx.Pipeline()
		pipe.Set(key,n+1,0)
		_, err = pipe.Exec()
		if err != nil {
			fmt.Printf("try connecting fail,err:%v\n",err)
			return err
		}

		println("over")
		return err
	},key)
	if errw != nil {
		fmt.Printf("try connecting fail,err:%v\n",errw)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello sendinfo!",
	})
}
