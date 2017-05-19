package main

import (
	_ "routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"fmt"
	"github.com/astaxie/beego/logs"
	"runtime"
	"common"
	"strconv"

	"models"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// 修改了默认的panic处理方法， 原方法在 config.go
	// 保证原有处理逻辑不变， 仅增加了处理serviceException的方法
	beego.BConfig.RecoverFunc = func(ctx *context.Context) {
		if err := recover(); err != nil {

			if err == beego.ErrAbort {
				return
			}
			if !beego.BConfig.RecoverPanic {
				panic(err)
			}
			// 新加逻辑， 如果可以转换为serviceException
			se, cok := err.(common.ServiceException)
			if cok {
				ctx.Output.SetStatus(400)
				ctx.Output.Body(se.Json())
				return
			}
			serviceError, cok := err.(common.ServiceError)
			if cok {
				ctx.Output.SetStatus(500)
				ctx.Output.Body(serviceError.Json())
				return
			}
			// 修改 私有方法调用为 公共方法   转换数据
			if beego.BConfig.EnableErrorsShow {
				if _, ok := beego.ErrorMaps[fmt.Sprint(err)]; ok {
					ecode, e := strconv.ParseInt(fmt.Sprint(err), 10, 64)
					if e != nil {
						panic("error code parse error" + e.Error())
					}
					beego.Exception(uint64(ecode), ctx)
					return
				}
			}

			var stack string
			logs.Critical("the request url is ", ctx.Input.URL())
			logs.Critical("Handler crashed with error", err)
			for i := 1; ; i++ {
				_, file, line, ok := runtime.Caller(i)
				if !ok {
					break
				}
				logs.Critical(fmt.Sprintf("%s:%d", file, line))
				stack = stack + fmt.Sprintln(fmt.Sprintf("%s:%d", file, line))
			}
			if beego.BConfig.EnableErrorsRender {
				//showErr(err, ctx, stack)

			}
		}
	}

	// 增加初始化 orm操作对象的钩子函数， 这个对象初始化之前要保证 model注册完
	beego.AddAPPStartHook(models.OrmInitHockFunc)

	beego.Run()
}
