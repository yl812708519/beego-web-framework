package devops

import (
	"daos"
	"log"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(ServerServing))
}

type ServerServing struct {
	Id              int64       `orm:"column(id)"`
	ServerId        int64       `orm:"column(server_id)"`
	ServingId       int64       `orm:"column(serving_id)"`
	Env             string      `orm:"column(env)"`
	Application     string      `orm:"column(application)"`
	EngineRoom      string      `orm:"column(engine_room)"`
	IntranetIp      string      `orm:"column(intranet_ip)"`
	ExtranetIp      string      `orm:"column(extranet_ip)"`
	IsDeleted       bool        `orm:"column(is_deleted)"`
	CreatorId       int64       `orm:"column(creator_id)"`
	UpdaterId       int64       `orm:"column(updater_id)"`
	CreatedAt       int64       `orm:"column(created_at)"`
	UpdatedAt       int64       `orm:"column(updated_at)"`

}

func (m ServerServing) TableName() string {
	return "server_servings"
}

// 关联表， 用于列表查询
type ServerServingDao struct {
	daos.BaseFunc
}


func (this ServerServingDao) FindByServingId(id int64) []ServerServing {
	qs := this.InitQuerySetter(&ServerServing{})
	qs = qs.Filter("ServingId", id)
	res := []ServerServing{}
	qs.All(&res)
	return res
}

// 根据条件查询serving 的列表, 多对多使用了分组
func (this ServerServingDao) FindServingList(application, engineRoom, env, ip string, page, pageSize int) ([]ServerServing, int64) {
	serverServing := ServerServing{}

	// 分组查询有问题。 会死在mysql only_full_group_by 上，设置多个字段不好用。。。
	qs := this.InitQuerySetter(&serverServing)

	if len(engineRoom) > 0 {
		qs = qs.Filter("EngineRoom", engineRoom)
	}
	if len(env) > 0 {
		qs = qs.Filter("Env", env)
	}

	if len(ip) > 0 {
		qs = qs.Filter("ExtranetIp", ip)
	}
	count, err := qs.GroupBy("serving_id").Count()
	if err != nil {
		log.Println(err)
	}

	var res []ServerServing
	qs.All(&res, "serving_id")
	return res, count

}





