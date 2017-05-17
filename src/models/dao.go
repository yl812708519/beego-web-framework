package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"reflect"
)

var (
	ormer orm.Ormer
	regModels map[string] reflect.Type
)

func init() {
	// 读取配置
	mysqlUser := beego.AppConfig.String("mysqlUser")
	mysqlPass := beego.AppConfig.String("mysqlPass")
	mysqlUrl:= beego.AppConfig.String("mysqlUrl")
	mysqlDb := beego.AppConfig.String("mysqlDb")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4", mysqlUser, mysqlPass, mysqlUrl, mysqlDb)
	log.Println("data source url: " + dataSource)
	//注册mysql
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dataSource)

}


func OrmInitHockFunc() error{
	ormer = orm.NewOrm()
	return nil
}


func Insert(m interface{}) interface{}{
	id , err := ormer.Insert(m)
	if err != nil {
		log.Panicln(err)
	}
	return id

}

func FindOne(m interface{}) {

	err := ormer.Read(m)
	if err != nil {
		log.Panicln(err)
	}
}



