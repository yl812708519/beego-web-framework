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
// @Description get all the staticblock by key
// @Param   key     path    string  true        "The email for login"
// @Success 200 {object} services.UserDTO
// @Failure 400 Invalid email supplied
// @Failure 404 User not found
// @router /throw/:key [get]
func (h *HelloController) Get() {
	var u *services.UserDTO
	u =  h.userService.FindById(1)
	u.Param = h.GetString("key")
	h.Data["json"] = *u
	h.ServeJSON()
}




