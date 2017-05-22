package controllers

import (
	"services"
)


// 示例代码
type HelloController struct {
	BaseController
	userService services.UserService
}



// @Title getStaticBlock
// @Description test api
// @Param   id     path    int64  true        "The email for login"
// @Success 200 {object} services.UserDTO
// @Failure 400 Invalid email supplied
// @Failure 400 User not found
// @router /throw/:id [get]
func (h HelloController) Get() {
	id := h.getInt64("id")
	u :=  h.userService.FindById(id)
	h.Data["json"] = u
	h.ServeJSON()
}




