package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"reflect"
	"common"
	"conf"
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
	mysqlUser := beego.AppConfig.String(conf.DATABASE_MYSQL_USER_NAME)
	mysqlPass := beego.AppConfig.String(conf.DATABASE_MYSQL_PASSWORD)
	mysqlUrl:= beego.AppConfig.String(conf.DATABASE_MYSQL_URL)
	mysqlDb := beego.AppConfig.String(conf.DATABASE_MYSQL_DB)
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

func (this BaseFunc) initQuerySetter(tableName string) orm.QuerySeter {
	return ormer.QueryTable(tableName)
}

func (b *BaseFunc) calculateOffset(page, pageSize int) int {
	return (page - 1) * pageSize
}

func (b *BaseFunc) Insert(m interface{}) int64{
	mType := reflect.TypeOf(m)
	if _, ok := mType.Elem().FieldByName(createdAtField); ok {
		setCreatedAt(m)
	}
	if _, ok := mType.Elem().FieldByName(createdAtField); ok {
		setUpdatedAt(m)
	}
	id , err := ormer.Insert(m)
	if err != nil {
		log.Panicln(err)
	}
	return id
}

func (b *BaseFunc) InsertAll(m interface{}){
	mv := reflect.ValueOf(m)
	if reflect.ValueOf(m).Kind() != reflect.Slice {
		panic(common.NewServiceException(20002))
	}
	if mv.Len() <= 0 {
		panic(common.NewServiceException(20005))
	}

	m0 := mv.Index(0)
	mType := m0.Type()
	fmt.Println(mType)
	_, iscf := mType.FieldByName(createdAtField)
	_, isuf := mType.FieldByName(updatedAtField)

	for i:=0; i<mv.Len(); i++ {
		v := mv.Index(i)
		if iscf {
			v.FieldByName(createdAtField).SetInt(common.GetTimeStamp())
		}
		if isuf {
			v.FieldByName(updatedAtField).SetInt(common.GetTimeStamp())
		}
	}
	_ , err := ormer.InsertMulti(5, m)
	if err != nil {
		log.Panicln(err)
	}
}


func (b *BaseFunc) Update(m interface{}) {
	mType := reflect.TypeOf(m)
	if _, ok:=mType.Elem().FieldByName(updatedAtField); ok {
		mValue := reflect.ValueOf(m).Elem().FieldByName(updatedAtField)
		mValue.SetInt(common.GetTimeStamp())
	}
	_, err := ormer.Update(m)
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func (b *BaseFunc) FindOne(id int64, m interface{}) interface{} {
	qs := ormer.QueryTable(m).Filter("Id", id)
	mType := reflect.TypeOf(m)
	if _, ok:=mType.Elem().FieldByName(isDeleteField); ok{
		qs = qs.Filter(isDeleteField, false)
	}
	err := qs.One(m)
	if err == orm.ErrNoRows {
		// 没有找到记录
		panic(common.NewServiceException(20001))
	}
	return m
}

func (b *BaseFunc) Remove(id int64, m interface{}){
	mType := reflect.TypeOf(m)
	if _, ok:=mType.Elem().FieldByName(isDeleteField); ok{
		// 有is_deleted字段
		b.FindOne(id, m)
		mValue := reflect.ValueOf(m).Elem().FieldByName(isDeleteField)
		mValue.SetBool(true)
		ormer.Update(m, isDeleteField)
	} else {
		ormer.QueryTable(m).Filter("Id", id).Delete()
	}

}













func setUpdatedAt(m interface{}) {
	mValue := reflect.ValueOf(m).Elem().FieldByName(updatedAtField)
	mValue.SetInt(common.GetTimeStamp())
}

func setCreatedAt(m interface{}) {
	mValue := reflect.ValueOf(m).Elem().FieldByName(createdAtField)
	mValue.SetInt(common.GetTimeStamp())
}

