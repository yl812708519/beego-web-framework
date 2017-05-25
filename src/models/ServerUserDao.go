package models

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(new(ServerUser))
}

type ServerUser struct{
	Id int64                `orm:"column(id)"`
	ServerId int64          `orm:"column(server_id)"`
	UserName string         `orm:"column(user_name)"`
	Password string         `orm:"column(password)"`
	IsDeleted bool          `orm:"column(is_deleted)"`
	CreatedAt int64         `orm:"column(created_at)"`

}

func (u ServerUser) TableName() string {
	return "server_users"
}


type ServerUserDao struct {
	BaseFunc
}
//
//func (this ServerUserDao) InsertAll(users []ServerUser) {
//	this.insertAll(users)
//}


func (this ServerUserDao) FindByServerIds(ids []int64) []ServerUser {
	qs := this.initQuerySetter("server_users").Filter("ServerId__in", ids)
	var users []ServerUser
	qs.All(&users)
	return users
}

func (this ServerUserDao) DeleteByServerId(id int64) {
	qs := this.initQuerySetter("server_users").Filter("ServerId", id)
	qs.Update(orm.Params{"IsDeleted": true})
}

func (this ServerUserDao) DeleteByServerIds(ids []int64) {
	qs := this.initQuerySetter("server_users").Filter("ServerId__in", ids)
	qs.Update(orm.Params{"IsDeleted": true})
}

