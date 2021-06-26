package tools

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

//获取DB对象
func ReturnMsqlDb() *sql.DB{
	mysql := GetSetting()//配置文件
	// init mysql db
	if err := initMySQL(mysql);err!=nil{
		fmt.Printf("try connecting fail,err:%v\n",err)
	}
	return db
}

// @title    initMySQL
// @description   初始化数据库连接函数
// @auth      amberhu         20210624 15:35
// @param     mysql           mysqlsetting     mysql设置参数
// @return    none-db            sql.DB          为全局参数赋值
// @return    err               error           报错
func initMySQL(mysql *Setting)(err error){
	dsn := mysql.Mysql.Username+":"+mysql.Mysql.Password+"@tcp("+mysql.Mysql.Host+":"+mysql.Mysql.Port+")/"+mysql.Mysql.Dbname
	db, err = sql.Open("mysql",dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	//db.SetConnMaxLifetime(time.Second * 10)
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(10)
	return err
}

//main里面用的初始化参数文件
func MysqlInitConnectParamInMain() string{
	err := initMySQL(GetSetting())
	if err != nil {
		fmt.Printf("mysql try connecting fail,err:%v\n",err)
		return "mysql try connecting fail"
	}else{
		fmt.Printf("mysql try connecting success\n")
		return "mysql try connecting success"
	}
}