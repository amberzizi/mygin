package mysql

import (
	"database/sql"
	"fmt"
	"github.com/gohouse/gorose"
	"go.uber.org/zap"
	"mygin/src/settings"
)

var Db *sql.DB
var Gdb *gorose.Connection

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
		zap.L().Error("mysql init faild", zap.Error(err))
	}

	err = Db.Ping()
	if err != nil {
		zap.L().Error("mysql init ping faild", zap.Error(err))
	}
	//db.SetConnMaxLifetime(time.Second * 10)
	Db.SetMaxOpenConns(mysql.Mysql.Maxconnection)
	Db.SetMaxIdleConns(mysql.Mysql.Maxidleconnection)
	return err
}

func initGoroseMySQL(mysql *settings.Setting) (err error) {
	var dbConfig1 = &gorose.DbConfigSingle{
		Driver:          "mysql",                                                                                                                             // 驱动: mysql/sqlite/oracle/mssql/postgres
		EnableQueryLog:  true,                                                                                                                                // 是否开启sql日志
		SetMaxOpenConns: mysql.Mysql.Maxconnection,                                                                                                           // (连接池)最大打开的连接数，默认值为0表示不限制
		SetMaxIdleConns: mysql.Mysql.Maxidleconnection,                                                                                                       // (连接池)闲置的连接数
		Prefix:          "",                                                                                                                                  // 表前缀
		Dsn:             mysql.Mysql.Username + ":" + mysql.Mysql.Password + "@tcp(" + mysql.Mysql.Host + ":" + mysql.Mysql.Port + ")/" + mysql.Mysql.Dbname, // 数据库链接
	}
	Gdb, err = gorose.Open(dbConfig1)
	if err != nil {
		zap.L().Error("mysql gorose init faild", zap.Error(err))
		return
	}
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

//main里面用的初始化参数文件
func MysqlGoroseInitConnectParamInMain() string {
	err := initGoroseMySQL(settings.GetSetting())
	if err != nil {
		fmt.Printf("mysql Gorose try connecting fail,err:%v\n", err)
		return "mysql Gorose try connecting fail"
	} else {
		fmt.Printf("mysql Gorose try connecting success\n")
		return "mysql Gorose try connecting success"
	}
}
