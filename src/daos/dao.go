package daos

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
	Ormer orm.Ormer
	regModels map[string] reflect.Type
)

const (
	Id = "id"
	IsDeleteField = "IsDeleted"
	CreatedAtField = "CreatedAt"
	UpdatedAtField = "UpdatedAt"
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
	Ormer = orm.NewOrm()
	return nil
}

// 基础方法类， 减少一些代码量
type BaseFunc struct {

}

// param: ptr value
// 如果要查被删掉的， 请自己初始化query setter
func (this BaseFunc) InitQuerySetter(m interface{}) orm.QuerySeter {
	// 默认按照创建时间倒叙排列
	qs := Ormer.QueryTable(m).OrderBy("-"+CreatedAtField)
	mType := reflect.TypeOf(m)
	if _, ok:=mType.Elem().FieldByName(IsDeleteField); ok{
		qs = qs.Filter(IsDeleteField, false)
	}
	return qs
}

// 虽然是公共的和调用方不在一个包里。 原则上只有dao层使用
func (this BaseFunc) InitFindByIdsQs(m interface{}, ids []int64) orm.QuerySeter {
	if len(ids) <= 0 {
		panic(common.NewServiceException(20005))
	}
	qs := Ormer.QueryTable(m).Filter(Id+"__in", ids)
	mType := reflect.TypeOf(m)
	if _, ok:=mType.Elem().FieldByName(IsDeleteField); ok{
		qs = qs.Filter(IsDeleteField, false)
	}
	return qs
}

func (b *BaseFunc) CalculateOffset(page, pageSize int) int {
	return (page - 1) * pageSize
}

// 接收指针
func (b *BaseFunc) Insert(m interface{}) int64{
	mType := reflect.TypeOf(m)
	if _, ok := mType.Elem().FieldByName(CreatedAtField); ok {
		setCreatedAt(m)
	}
	if _, ok := mType.Elem().FieldByName(CreatedAtField); ok {
		setUpdatedAt(m)
	}
	id , err := Ormer.Insert(m)
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
	_, iscf := mType.FieldByName(CreatedAtField)
	_, isuf := mType.FieldByName(UpdatedAtField)

	for i:=0; i<mv.Len(); i++ {
		v := mv.Index(i)
		if iscf {
			v.FieldByName(CreatedAtField).SetInt(common.GetTimeStamp())
		}
		if isuf {
			v.FieldByName(UpdatedAtField).SetInt(common.GetTimeStamp())
		}
	}
	_ , err := Ormer.InsertMulti(5, m)
	if err != nil {
		log.Panicln(err)
	}
}


func (b *BaseFunc) Update(m interface{}) {
	mType := reflect.TypeOf(m)
	if _, ok:=mType.Elem().FieldByName(UpdatedAtField); ok {
		mValue := reflect.ValueOf(m).Elem().FieldByName(UpdatedAtField)
		mValue.SetInt(common.GetTimeStamp())
	}
	_, err := Ormer.Update(m)
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func (b *BaseFunc) FindOne(id int64, m interface{}) interface{} {
	qs := Ormer.QueryTable(m).Filter("Id", id)
	mType := reflect.TypeOf(m)
	if _, ok:=mType.Elem().FieldByName(IsDeleteField); ok{
		qs = qs.Filter(IsDeleteField, false)
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
	if _, ok:=mType.Elem().FieldByName(IsDeleteField); ok{
		// 有is_deleted字段
		b.FindOne(id, m)
		mValue := reflect.ValueOf(m).Elem().FieldByName(IsDeleteField)
		mValue.SetBool(true)
		Ormer.Update(m, IsDeleteField)
	} else {
		Ormer.QueryTable(m).Filter("Id", id).Delete()
	}
}

func (b *BaseFunc) RemoveByIds(ids []int64, m interface{}){
	mType := reflect.TypeOf(m)
	if _, ok:=mType.Elem().FieldByName(IsDeleteField); ok{
		// 有is_deleted字段
		qs := b.InitQuerySetter(m)
		qs.Filter("Id__in", ids).Update(orm.Params{IsDeleteField: true})
	} else {
		Ormer.QueryTable(m).Filter("Id_in", ids).Delete()
	}
}











func setUpdatedAt(m interface{}) {
	mValue := reflect.ValueOf(m).Elem().FieldByName(UpdatedAtField)
	mValue.SetInt(common.GetTimeStamp())
}

func setCreatedAt(m interface{}) {
	mValue := reflect.ValueOf(m).Elem().FieldByName(CreatedAtField)
	mValue.SetInt(common.GetTimeStamp())
}

