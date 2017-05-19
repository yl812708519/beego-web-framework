package controllers

import "services/server"

type ServerController struct {
	BaseController
	serverService server.ServerService
}



// @Title find server
// @Description 查询server信息
// @Param   id     path    int64  true        "server id"
// @Success 200 {object} services.ServerDTO
// @Failure 400 Invalid email supplied
// @Failure 404 User not found
// @router /server/:id [get]
func (c *ServerController) FindById() {
	id := c.getInt64("id")
	serverDTO := c.serverService.FindById(id)
	c.renderJSON(serverDTO)
}

