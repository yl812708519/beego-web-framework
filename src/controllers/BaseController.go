package controllers

import (
	"github.com/astaxie/beego"
	"common"
	"conf"
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

func (b BaseController) setCookie(key, value string){
	b.SetSecureCookie(beego.AppConfig.String(conf.COOKIE_SECRET_KEY), key, value)
}


func (b BaseController) getCookie(key string) string {
	v, ok := b.GetSecureCookie(beego.AppConfig.String(conf.COOKIE_SECRET_KEY), key)
	if !ok {
		panic(common.NewServiceError(10002))
	}
	return v
}


