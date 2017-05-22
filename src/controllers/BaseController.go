package controllers

import (
	"github.com/astaxie/beego"
	"common"
	"strings"
)


// 常用的controller 方法放到这里
type BaseController struct {
	 beego.Controller
}

// 直接写key 就可以， 在这里加上:
func (b BaseController) getInt64(key string, def ...int64) int64 {
	if !strings.HasPrefix(key, ":"){
		key = ":" + key
	}
	v , err:= b.GetInt64(key, def...)
	if err!=nil  {
		panic(common.NewServiceException(20003))
	}
	return v
}



func (b BaseController) renderJSON(object interface{}){
	b.Data["json"] = object
	b.ServeJSON()
}

// post 接口返回成功结果
func (b BaseController) renderSuccess() {
	b.Data["json"] = map[string] string{"status": "success"}
}



