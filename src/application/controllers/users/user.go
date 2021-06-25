package user

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"mygin/src/tools"
	"net/http"
	"strconv"
	"time"
)
type configuration struct {
	Enabled bool
	Path    string
	Username  string
	Passwd  string
}

type configuration2 struct {
	Section struct{
		Enabled bool
		Path    string
	}
}


func Sendinfo(c *gin.Context){
	//json config
	//file, _ := os.Open("src/conf/systeminfo.json")
	//defer file.Close()
	//decoder := json.NewDecoder(file)
	//conf := configuration{}
	//err := decoder.Decode(&conf)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//}
	//fmt.Println(conf)

	//config := configuration2{}
	//err := gcfg.ReadFileInto(&config, "src/conf/systeminfo.ini")
	//if err != nil {
	//	fmt.Println("Failed to parse config file: %s", err)
	//}
	//fmt.Println(config.Section.Path)

	test := false
	if test  {
		//测试zaplog
		logger,_ := tools.LogerProducter()
		logger.Warn("watch user...")
		//测试二维码生成
		randfinal := rand.New(rand.NewSource(time.Now().UnixNano()))
		randname := randfinal.Intn(1000)
		var url = tools.CreateQrcode(200,200,"testinfo",strconv.Itoa(randname))
		println(url)
	}

	//测试redis
	rdb := tools.ReturnRedisDb()
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


	c.JSON(http.StatusOK, gin.H{
		"message": "Hello sendinfo!",
	})
}
