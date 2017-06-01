package devops

import (
	"github.com/astaxie/beego/orm"
	"devops/daos"
)

func init() {
	orm.RegisterModel(new(ServerDisk))
}

type ServerDisk struct{
	Id int64                `orm:"column(id)"`
	ServerId int64          `orm:"column(server_id)"`
	RootPath string         `orm:"column(root_path)"`
	Size int                `orm:"column(size)"`
	IsDeleted bool          `orm:"column(is_deleted)"`
	CreatedAt int64         `orm:"column(created_at)"`

}

func (this ServerDisk) TableName() string {
	return "server_disks"
}


type ServerDiskDao struct {
	daos.BaseFunc
}

//func (this ServerDiskDao) InsertAll(disks []ServerDisk) {
//	this.insertAll(disks)
//}

func (this ServerDiskDao) FindByServerIds(ids []int64) []ServerDisk {
	qs := this.InitQuerySetter(&ServerDisk{}).Filter(daos.IsDeleteField, false)
	qs = qs.Filter("ServerId__in", ids)
	var disks []ServerDisk
	qs.All(&disks)
	return disks
}


func (this ServerDiskDao) DeleteByServerId(id int64) {
	qs := this.InitQuerySetter(&ServerDisk{}).Filter(daos.IsDeleteField, false).Filter("ServerId", id)
	qs.Update(orm.Params{"IsDeleted": true})
}

func (this ServerDiskDao) DeleteByServerIds(ids []int64) {
	qs := this.InitQuerySetter(&ServerDisk{}).Filter("ServerId__in", ids)
	qs.Update(orm.Params{"IsDeleted": true})
}