package controllers

import (
	"services/server"
	"common"
	"github.com/astaxie/beego"
	"services"
	"fmt"
	"conf"
)

type ServerController struct {
	BaseController
	serverService server.ServerService
}


//静态数据接口

// @Title get applications
// @Description   获取静态数据列表 key : applications/envs/engineRom
// @Success 200 like this:  { result: ["aa", "bb", "cc"] }
// @Failure 400 service exception
// @router /constant/:key [get]
func (this *ServerController) GetApplications() {
	key := this.GetString(":key")

	v := services.ConstantMap[key]
	fmt.Println(v)
	if v == nil {
		panic(common.NewServiceException(20001))
	}
	r := services.ResultVO{Result:v}
	this.renderJSON(r)
}


// @Title create server
// @Description
// @Param    tag            formData    string          true        "名称标签"
// @Param    env            formData    string          false        "环境"
// @Param    engineRoom     formData    string          true        "机房"
// @Param    core           formData    int             true        "核心数"
// @Param    memory         formData    int             true        "内存数"
// @Param    intranetIp     formData    string          false        "内网ip"
// @Param    extranetIp     formData    string          false        "外网ip"
// @Param    remark         formData    string          false        "备注"
// @Param    users          formData    []server.User   true        "用户"
// @Param    disks          formData    []server.Disk   true        "磁盘"
// @Success 200 {status: success}
// @Failure 400 service exception
// @router /servers [post]
func (this *ServerController) Create() {
	serverRequest := &server.ServerDTO{}
	this.parseJsonRequest(serverRequest)
	this.valid(*serverRequest)
	this.serverService.Create(*serverRequest)

	this.renderSuccess()
}

// @Title create server
// @Description
// @Param    tag            formData    string          true        "名称标签"
// @Param    env            formData    string          false        "环境"
// @Param    engineRoom     formData    string          true        "机房"
// @Param    core           formData    int             true        "核心数"
// @Param    memory         formData    int             true        "内存数"
// @Param    intranetIp     formData    string          false        "内网ip"
// @Param    extranetIp     formData    string          false        "外网ip"
// @Param    remark         formData    string          false        "备注"
// @Param    users          formData    []server.User   true         "用户"
// @Param    disks          formData    []server.Disk   true         "磁盘"
// @Success 200 {status: success}
// @Failure 400 service exception
// @router /servers/:id [put]
func (this *ServerController) Update() {
	serverRequest := &server.ServerDTO{}
	this.parseJsonRequest(serverRequest)
	serverRequest.Id, _ = this.GetInt64(":id")
	this.serverService.Update(*serverRequest)

	this.renderSuccess()
}



// @Title find server
// @Description 查询server信息
// @Param   id     path    int64  true        "server id"
// @Success 200 {object} services.ServerDTO
// @Failure 400 400 service exception
// @router /servers/:id [get]
func (this *ServerController) FindById() {
	id, err := this.GetInt64(":id")
	if err != nil{
		panic(common.NewServiceException(20003))
	}
	serverDTO := this.serverService.FindById(id)
	this.renderJSON(serverDTO)
}



// @Title create server
// @Description  查询列表， result 结构为 serverDTO 数组
// @Param    page           query    int     false        "页码, 缺省：1 "
// @Param    pageSize       query    int     false        "步长, 缺省:15"
// @Param    application    query    string  false        "应用"
// @Param    engineRoom     query    string  false        "机房"
// @Param    env            query    int     false        "环境"
// @Param    ip             query    string  false        "ip"
// @Success 200 resultsVO:  { "count": 0, "results": [ serverDTO ] }
// @Failure 400 service exception
// @router /servers [get]
func (this *ServerController) FindList() {
	r := server.ListRequest{}

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


	resultVO := this.serverService.FindList(r)
	this.renderJSON(resultVO)
}


// @Title delete server
// @Description 删除server信息
// @Param   id     path    int64  true        "server id"
// @Success 200 {status: "success"}
// @Failure 400 400 service exception
// @router /servers/:id [delete]
func (this *ServerController) Remove() {
	id, err := this.GetInt64(":id")
	if err != nil{
		panic(common.NewServiceException(20003))
	}
	this.serverService.Remove(id)
	this.renderSuccess()
}


// @Title delete server
// @Description 删除server信息
// @Param   ids    formData    []int64    true        "json like this: {"ids": [2,3,4,5,6]}"
// @Success 200 {status: "success"}
// @Failure 400 400 service exception
// @router /servers [delete]
func (this *ServerController) RemoveByIds() {
	m := map[string][]int64{}
	this.parseJsonRequest(&m)
	this.serverService.RemoveByIds(m["ids"])
	this.renderSuccess()
}



