package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/gcfg.v1"
	"net/http"
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

	config := configuration2{}
	err := gcfg.ReadFileInto(&config, "src/conf/systeminfo.ini")
	if err != nil {
		fmt.Println("Failed to parse config file: %s", err)
	}
	fmt.Println(config.Section.Path)

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello sendinfo!",
	})
}
