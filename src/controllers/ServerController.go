package controllers

import (
	"services/server"
	"common"
	"encoding/json"
	"github.com/astaxie/beego"
)

type ServerController struct {
	BaseController
	serverService server.ServerService
}



// @Title create server
// @Description
// @Param    tag    formData    string  true        "名称标签"
// @Param    env    formData    string  true        "环境"
// @Param    engineRoom    formData    string  true        "机房"
// @Param    core    formData    int  true        "核心数"
// @Param    memory    formData    int  true        "内存数"
// @Param    intranetIp    formData    string  true        "内网ip"
// @Param    extranetIp    formData    string  true        "外网ip"
// @Param    remark    formData    string  true        "备注"
// @Param    users    formData    []server.User  true        "用户"
// @Param    disks    formData    []server.Disk  true        "磁盘"
// @Success 200 {status: success}
// @Failure 400 service exception
// @router /servers [post]
func (this *ServerController) Create() {
	serverRequest := &server.ServerDTO{}
	if err := this.ParseForm(serverRequest); err != nil {
		panic(err)
	}
	// todo 对接口时测试form-data, 是否可以绑定
	//users := make([]string, 0, 2)
	//this.Ctx.Input.Bind(&users, "users")

	json.Unmarshal(this.Ctx.Input.RequestBody, serverRequest)
	this.serverService.Create(*serverRequest)

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
// @Success 200 {object} services.ResultPageVO
// @Failure 400 service exception
// @router /servers [get]
func (this *ServerController) FindList() {
	r := server.ListRequest{}

	d, err := beego.AppConfig.Int("defaultPageSize")
	if err!=nil {
		panic(common.NewServiceError(10001))
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




