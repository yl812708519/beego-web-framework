package models

import "github.com/astaxie/beego/orm"

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
	BaseFunc
}

//func (this ServerDiskDao) InsertAll(disks []ServerDisk) {
//	this.insertAll(disks)
//}

func (this ServerDiskDao) FindByServerIds(ids []int64) []ServerDisk {
	qs := ormer.QueryTable("server_disks")
	qs.Filter("ServerId__in", ids)
	var disks []ServerDisk
	qs.All(&disks)
	return disks
}

