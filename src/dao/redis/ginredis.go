// @Title  ginredis.go
// @Description  zap日志创建，tools.ReturnRedisDb() redis对象  单点 未扩展哨兵 集群
// @Author  amberhu  20210625
// @Update

//测试redis
//rdb := tools.ReturnRedisDb()
//defer rdb.Close()   回收*****
package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"mygin/src/settings"
)

var Rdb *redis.Client

//对外返回redis连接对象
//可以直接用redis.Rdb
func ReturnRedisDb() *redis.Client {
	redisset := settings.GetSetting() //配置文件
	// init mysql db
	if err := initRedisClient(redisset); err != nil {
		fmt.Printf("try connecting fail,err:%v\n", err)
	}
	return Rdb
}

//初始化redis 连接
func initRedisClient(redisset *settings.Setting) (err error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     redisset.Redis.Host,
		Password: redisset.Redis.Password,
		DB:       redisset.Redis.Db,
		PoolSize: redisset.Redis.Poolsize,
	})

	_, err = Rdb.Ping().Result()
	return err
}

//main里面用的初始化参数文件 初始化连接
func ReidsInitConnectParamInMain() string {
	err := initRedisClient(settings.GetSetting())
	if err != nil {
		fmt.Printf("redis try connecting fail,err:%v\n", err)
		return "redis try connecting fail,err"
	} else {
		return "redis try connecting success"
	}
}
