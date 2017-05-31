package devops

import (
	"github.com/astaxie/beego/orm"
	"log"
	"daos"
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
	daos.BaseFunc
}

func (this ServerDao) RemoveByIds(ids []int64) {
	if len(ids) <= 0 {
		return
	}
	qs := daos.Ormer.QueryTable(&Server{}).Filter("Id__in", ids)
	qs.Update(orm.Params{
		daos.IsDeleteField : true,
	})
}

func (this ServerDao) FindList(application, engineRoom, env, ip string, page, pageSize int) ([]Server, int64) {
	server := Server{}
	qs := daos.Ormer.QueryTable(&server).Filter(daos.IsDeleteField, false)

	if len(application) > 0 {
		qs = qs.Filter("Application", application)
	}
	if len(engineRoom) > 0 {
		qs = qs.Filter("EngineRoom", engineRoom)
	}
	if len(env) > 0 {
		qs = qs.Filter("Env", env)
	}
	if len(ip) > 0 {
		qs = qs.Filter("ExtranetIp", ip)
	}
	count, err := qs.Count()
	if err != nil {
		log.Println(err)
	}

	qs = qs.Limit(pageSize).Offset(this.CalculateOffset(page, pageSize))
	var servers []Server
	qs.All(&servers)
	return servers, count

}


func (this ServerDao) FindByIds(ids []int64) []Server {
	qs := this.InitFindByIdsQs(&Server{}, ids)
	res := []Server{}
	qs.All(&res)
	return res
}

func (this ServerDao) FindMapByIds(ids []int64) map[int64]Server {
	servers := this.FindByIds(ids)
	result := map[int64]Server{}
	for _, x := range servers {
		result[x.Id] = x
	}
	return result
}

