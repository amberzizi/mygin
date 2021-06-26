package mysql

import (
	"database/sql"
	"fmt"
	"mygin/src/settings"
)

var Db *sql.DB

//获取DB对象
//直接可以用mysq.Db获取
func ReturnMsqlDb() *sql.DB {
	mysql := settings.GetSetting() //配置文件
	// init mysql db
	if err := initMySQL(mysql); err != nil {
		fmt.Printf("try connecting fail,err:%v\n", err)
	}
	return Db
}

// @title    initMySQL
// @description   初始化数据库连接函数
// @auth      amberhu         20210624 15:35
// @param     mysql           mysqlsetting     mysql设置参数
// @return    none-db            sql.DB          为全局参数赋值
// @return    err               error           报错
func initMySQL(mysql *settings.Setting) (err error) {
	dsn := mysql.Mysql.Username + ":" + mysql.Mysql.Password + "@tcp(" + mysql.Mysql.Host + ":" + mysql.Mysql.Port + ")/" + mysql.Mysql.Dbname
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = Db.Ping()
	//db.SetConnMaxLifetime(time.Second * 10)
	Db.SetMaxOpenConns(200)
	Db.SetMaxIdleConns(10)
	return err
}

//main里面用的初始化参数文件
func MysqlInitConnectParamInMain() string {
	err := initMySQL(settings.GetSetting())
	if err != nil {
		fmt.Printf("mysql try connecting fail,err:%v\n", err)
		return "mysql try connecting fail"
	} else {
		fmt.Printf("mysql try connecting success\n")
		return "mysql try connecting success"
	}
}
