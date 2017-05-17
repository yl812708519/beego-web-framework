package controllers

import (
	"meili-devops-backend/services"
)
// 示例代码
type HelloController struct {
	BaseController
	userService services.UserService
}

func (h *HelloController) Get() {
	h.userService.FindById(1)
}




