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
	"gopkg.in/gcfg.v1"
)

var rdb *redis.Client

type redissetting struct {
	//Section struct{
	//	Enabled bool
	//	Path    string
	//}
	Redis struct{
		Host  string
		Password string
		Db int
		Poolsize int
	}
}

//初始化redis 连接
func initRedisClient(redisset *redissetting) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: redisset.Redis.Host,
		Password: redisset.Redis.Password,
		DB:	redisset.Redis.Db,
		PoolSize: redisset.Redis.Poolsize,
	})

	_, err = rdb.Ping().Result()
	return err
}

//对外返回redis连接对象
func ReturnRedisDb() *redis.Client{
	redisset := returnRedisSetting()
	// init mysql db
	if err := initRedisClient(redisset);err!=nil{
		fmt.Printf("try connecting fail,err:%v\n",err)
	}
	return rdb
}

//获取配置文件参数
func returnRedisSetting() *redissetting {
	redisset := redissetting{}
	err := gcfg.ReadFileInto(&redisset, "src/conf/systeminfo.ini")
	if err != nil {
		fmt.Println("Failed to parse config file: %s", err)
	}
	return &redisset
}

//使用样例
func testdb(){
	rdb := ReturnRedisDb()
	defer rdb.Close()
	err := rdb.Set("score",100,0).Err()
	if err != nil {
		println("faild set")
		return
	}
	val, err := rdb.Get("score").Result()
	if err != nil {
		println("faild get")
		return
	}
	println(val)
}