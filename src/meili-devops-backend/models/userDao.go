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
}

func (u *User) TableName() string {
	return "users"
}



type UserDao struct {
}

func (d *UserDao) GetNewModel() *User{
	return new(User)
}





