package controllers

import (
	"github.com/astaxie/beego"
	"services/devops"
	"common"
	"conf"
)

type ServingController struct {
	BaseController
	servingService devops.ServingService

}


// @Title 创建服务
// @Description
// @Param    tag            formData    string          true        "名称标签"
// @Param    serverIds      formData    []int64         false       "服务器ids 数组"
// @Param    application    formData    string          true        "应用"
// @Param    url            formData    string          false       "URL"
// @Param    version        formData    string          true        "版本"
// @Param    dependency     formData    string          false       "依赖"
// @Param    remark         formData    string          false       "备注"
// @Success 200 {status: success}
// @Failure 400 service exception
// @router /servings [post]
func (this *ServingController) Create() {
	s := devops.ServingCreateDTO{}
	this.parseJsonRequest(&s)
	this.valid(s)
	this.servingService.Create(s)
	this.renderSuccess()
}



// @Title 查找一个服务
// @Description
// @Param    id       query    int64          true       "service id"
// @Success 200 {object} serving.ServingDetailDTO
// @Failure 400 service exception
// @router /servings/:id [get]
func (this *ServingController) FindOne() {
	id, err := this.GetInt64(":id")
	if err != nil{
		panic(common.NewServiceException(20003))
	}
	serverDTO := this.servingService.FindById(id)
	this.renderJSON(serverDTO)

}


// @Title 查询服务列表
// @Description
// @Param    page           query    int     false        "页码, 缺省：1 "
// @Param    pageSize       query    int     false        "步长, 缺省:15"
// @Param    application    query    string  false        "应用"
// @Param    engineRoom     query    string  false        "机房"
// @Param    env            query    int     false        "环境"
// @Param    ip             query    string  false        "ip"
// @Success 200 {object} []serving.ServingDTO
// @Failure 400 service exception
// @router /servings [get]
func (this *ServingController) FindList() {
	r := devops.ListRequest{}

	d, err := beego.AppConfig.Int(conf.DEFAULT_GET_LIST_PAGE_SIZE)
	if err!=nil {
		panic(common.NewServiceError(10000))
	}
	var e error
	r.Page, e = this.GetInt("page", 1)
	r.PageSize, e = this.GetInt("pageSize", d)
	if e != nil {
		panic(common.NewServiceException(20003))
	}
	r.Application = this.GetString("application")
	r.EngineRoom= this.GetString("engineRoom")
	r.Env= this.GetString("env")
	r.Ip= this.GetString("ip")

	resultVO := this.servingService.FindServings(r)
	this.renderJSON(resultVO)
}