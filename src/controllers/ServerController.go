package controllers

import (
	"services/server"
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
// @Success 200 {status: success}
// @Failure 400 service exception
// @router /server [post]
func (this *ServerController) Create() {
	serverRequest := &server.ServerDTO{}
	if err := this.ParseForm(serverRequest); err != nil {
		panic(err)
	}
	this.serverService.Create(*serverRequest)
	this.renderSuccess()
}



// @Title find server
// @Description 查询server信息
// @Param   id     path    int64  true        "server id"
// @Success 200 {object} services.ServerDTO
// @Failure 400 400 service exception
// @router /server/:id [get]
func (this *ServerController) FindById() {
	id := this.getInt64("id")
	serverDTO := this.serverService.FindById(id)
	this.renderJSON(serverDTO)
}
