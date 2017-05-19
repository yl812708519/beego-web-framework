package models

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(new(Server))
}

type Server struct{
	Id int64                `orm:"column(id)"`
	ServerId string         `orm:"column(server_id)"`
	Tag string              `orm:"column(tag)"`
	Application string      `orm:"column(application)"`
	Env string              `orm:"column(env)"`
	EngineRoom string       `orm:"column(engine_room)"`
	Core int                `orm:"column(core)"`
	Memory int              `orm:"column(memory)"`
	IntranetIp string       `orm:"column(intranet_ip)"`
	ExtranetIp string       `orm:"column(extranet_ip)"`
	Remark string           `orm:"column(remark)"`
	IsDeleted bool          `orm:"column(is_deleted)"`
	UpdaterId int64         `orm:"column(updater_id)"`
	CreatedAt int64         `orm:"column(created_at)"`
	UpdatedAt int64         `orm:"column(updated_at)"`

}

func (u Server) TableName() string {
	return "servers"
}


type ServerDao struct {

}

func (d ServerDao) FindOne(id int64) Server {
	server := Server{}
	findOne(id, &server)
	return server
}



