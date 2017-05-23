package controllers

import (
	"github.com/astaxie/beego"
)


// 常用的controller 方法放到这里
type BaseController struct {
	 beego.Controller
}

func (b BaseController) renderJSON(object interface{}){
	b.Data["json"] = object
	b.ServeJSON()
}

// post 接口返回成功结果
func (b BaseController) renderSuccess() {
	b.Data["json"] = map[string] string{"status": "success"}
	b.ServeJSON()
}


