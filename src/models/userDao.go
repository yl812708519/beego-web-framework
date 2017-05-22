package models

import (
	"github.com/astaxie/beego/orm"
)

// 示例代码
func init() {
	orm.RegisterModel(new(User))
}

type User struct{
	Id int64            `orm:"column(id)"`
	Name string         `orm:"column(name)"`
	IsGay bool          `orm:"column(is_gay)"`
}

func (u User) TableName() string {
	return "users"
}



type UserDao struct {
	BaseFunc
}

func (this UserDao) GetNewModel() *User{
	return new(User)
}

func (this UserDao) FindOne(id int64) User{
	u := User{}
	this.findOne(id, &u)
	return u
}





