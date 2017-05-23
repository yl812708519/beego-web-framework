package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"reflect"
	"common"
)

var (
	ormer orm.Ormer
	regModels map[string] reflect.Type
)

const (
	isDeleteField = "IsDeleted"
	createdAtField = "CreatedAt"
	updatedAtField = "CreatedAt"
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


// 基础方法类， 减少一些代码量
type BaseFunc struct {

}

func (b *BaseFunc) calculateOffset(page, pageSize int) int {
	return (page - 1) * pageSize
}

func (b *BaseFunc) Insert(m interface{}) int64{
	id , err := ormer.Insert(m)
	if err != nil {
		log.Panicln(err)
	}
	return id

}

func (b *BaseFunc) InsertAll(m interface{}){
	_ , err := ormer.InsertMulti(5, m)
	if err != nil {
		log.Panicln(err)
	}
}


func (b *BaseFunc) FindOne(id int64, m interface{}) interface{} {
	qs := ormer.QueryTable(m).Filter("Id", id)
	err := qs.One(m)
	if err == orm.ErrNoRows {
		// 没有找到记录
		panic(common.NewServiceException(20001))
	}
	return m
}

func (b *BaseFunc) Remove(id int64, m interface{}){
	mType := reflect.TypeOf(m)
	if _, ok:=mType.FieldByName(isDeleteField); ok{
		// 有is_deleted字段
		b.FindOne(id, m)
		mValue := reflect.ValueOf(m).FieldByName(isDeleteField)
		mValue.SetBool(true)
		ormer.Update(m, isDeleteField)
	} else {
		ormer.QueryTable(m).Filter("Id", id).Delete()
	}

}





