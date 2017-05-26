package devops

import (
	"github.com/astaxie/beego/orm"
	"daos"
)

func init() {
	orm.RegisterModel(new(Serving))
}


type Serving struct {
	Id          int64       `orm:"column(id)"`
	Tag         string      `orm:"column(tag)"`
	Application string      `orm:"column(application)"`
	ServerId    int64       `orm:"column(server_id)"`
	Url         string      `orm:"column(url)"`
	Version     string      `orm:"column(version)"`
	Dependency  string      `orm:"column(dependency)"`
	Remark      string      `orm:"column(remark)"`
	CreatorId   int64       `orm:"column(creator_id)"`
	UpdaterId   int64       `orm:"column(updater_id)"`
	IsDeleted   bool        `orm:"column(is_deleted)"`
	CreatedAt   int64       `orm:"column(created_at)"`
	UpdatedAt   int64       `orm:"column(update_at)"`
}

func (u Serving) TableName() string {
	return "services"
}
type ServingDao struct {
	daos.BaseFunc
}







