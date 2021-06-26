// @Title  ginredis.go
// @Description  zap日志创建，tools.ReturnRedisDb() redis对象  单点 未扩展哨兵 集群
// @Author  amberhu  20210625
// @Update

//测试redis
//rdb := tools.ReturnRedisDb()
//defer rdb.Close()   回收*****
package tools

import (
	"fmt"
	"github.com/go-redis/redis"
)

var rdb *redis.Client

//对外返回redis连接对象
func ReturnRedisDb() *redis.Client{
	redisset := GetSetting()//配置文件
	// init mysql db
	if err := initRedisClient(redisset);err!=nil{
		fmt.Printf("try connecting fail,err:%v\n",err)
	}
	return rdb
}


//初始化redis 连接
func initRedisClient(redisset *Setting) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: redisset.Redis.Host,
		Password: redisset.Redis.Password,
		DB:	redisset.Redis.Db,
		PoolSize: redisset.Redis.Poolsize,
	})

	_, err = rdb.Ping().Result()
	return err
}

//main里面用的初始化参数文件 初始化连接
func ReidsInitConnectParamInMain() string{
	err := initRedisClient(GetSetting())
	if err != nil{
		fmt.Printf("redis try connecting fail,err:%v\n",err)
		return "redis try connecting fail,err"
	}else{
		return "redis try connecting success"
	}
}
