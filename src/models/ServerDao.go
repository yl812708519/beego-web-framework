package models

import (
	"github.com/astaxie/beego/orm"
	"log"
)

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
	CreatorId int64         `orm:"column(creator_id)"`
	UpdaterId int64         `orm:"column(updater_id)"`
	CreatedAt int64         `orm:"column(created_at)"`
	UpdatedAt int64         `orm:"column(updated_at)"`

}

func (u Server) TableName() string {
	return "servers"
}


type ServerDao struct {
	BaseFunc
}

func (this ServerDao) RemoveByIds(ids []int64) {
	if len(ids) <= 0 {
		return
	}
	qs := ormer.QueryTable(&Server{}).Filter("Id__in", ids)
	qs.Update(orm.Params{
		isDeleteField : true,
	})
}

func (this ServerDao) FindList(application, engineRoom, env, ip string, page, pageSize int) ([]Server, int64) {
	server := Server{}

	cond := orm.NewCondition()

	if len(application) > 0 {
		cond.And("Application", application)
	}
	if len(engineRoom) > 0 {
		cond.And("EngineRoom", engineRoom)
	}
	if len(env) > 0 {
		cond.And("Env", env)
	}
	if len(ip) > 0 {
		ipCond := orm.NewCondition()
		ipCond.And("IntranetIp", ip)
		ipCond.Or("ExtranetIp", ip)
		cond.AndCond(ipCond)
	}
	qs := ormer.QueryTable(&server).SetCond(cond).Filter(isDeleteField, false)
	count, err := qs.Count()
	if err != nil {
		log.Println(err)
	}

	qs = qs.Limit(pageSize).Offset(this.calculateOffset(page, pageSize))
	var servers []Server
	qs.All(&servers)
	return servers, count

}



