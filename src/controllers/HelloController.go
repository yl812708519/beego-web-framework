package controllers

import (
	"services"
	"common"
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
	id, err := h.GetInt64(":id")
	if err != nil{
		panic(common.NewServiceException(20003))
	}
	u :=  h.userService.FindById(id)
	h.Data["json"] = u
	h.ServeJSON()
}


