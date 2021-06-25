package user

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	tools2 "mygin/src/application/controllers/tools"
	"net/http"
	"strconv"
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

	//测试zaplog
	logger,_ := tools2.LogerProducter()
	logger.Warn("watch user...")
	//测试二维码生成
	randname := rand.Intn(1000)
	var url = tools2.CreateQrcode(200,200,"testinfo",strconv.Itoa(randname))

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello sendinfo!"+url,
	})
}
