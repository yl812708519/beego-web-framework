package devops

import (
	"github.com/astaxie/beego/orm"
	"devops/daos"
)

func init() {
	orm.RegisterModel(new(Serving))
}


type Serving struct {
	Id          int64       `orm:"column(id)"`
	Tag         string      `orm:"column(tag)"`
	Application string      `orm:"column(application)"`
	Url         string      `orm:"column(url)"`
	Version     string      `orm:"column(version)"`
	Dependency  string      `orm:"column(dependency)"`
	Remark      string      `orm:"column(remark)"`
	CreatorId   int64       `orm:"column(creator_id)"`
	UpdaterId   int64       `orm:"column(updater_id)"`
	IsDeleted   bool        `orm:"column(is_deleted)"`
	CreatedAt   int64       `orm:"column(created_at)"`
	UpdatedAt   int64       `orm:"column(updated_at)"`

}

func (u Serving) TableName() string {
	return "servings"
}
type ServingDao struct {
	daos.BaseFunc
}

func (this ServingDao) FindByIds(ids []int64) []Serving {
	qs := this.InitFindByIdsQs(&Serving{}, ids)
	res := []Serving{}
	qs.All(&res)
	return res
}
func (this ServingDao) FindMapByIds(ids []int64) map[int64]Serving {
	servings := this.FindByIds(ids)
	result := map[int64]Serving{}
	for _, x := range servings {
		result[x.Id] = x
	}
	return result
}

func (this ServingDao) FindServingList(application string, page, pageSize int) ([]Serving, int64) {
	qs := this.InitQuerySetter(&Serving{})

	if len(application) > 0 {
		qs = qs.Filter("Application", application)
	}
	count,_ := qs.Count()
	qs = qs.Limit(pageSize, this.CalculateOffset(page, pageSize))
	c := []Serving{}
	qs.All(&c)
	return c, count
}






