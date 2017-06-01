package main

import (
	_ "devops/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"fmt"
	"github.com/astaxie/beego/logs"
	"runtime"
	"devops/common"
	"devops/daos"
	"os"
	"log"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		modeStr := args[1]
		if modeStr[:5] == "-mode" {
			beego.LoadAppConfig("ini", "conf/"+modeStr[6:]+".app.conf")
		}
	}
	log.Println("current run mode : "+ beego.BConfig.RunMode)
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// 修改了默认的panic处理方法， 原方法在 config.go
	// panic之后
	// 如果是serviceException： 返回json信息
	// 否则 输出 未知的系统错误, 并打印栈，
	beego.BConfig.RecoverFunc = func(ctx *context.Context) {
		if err := recover(); err != nil {

			if err == beego.ErrAbort {
				return
			}
			if !beego.BConfig.RecoverPanic {
				panic(err)
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


			// 新加逻辑， 如果可以转换为serviceException
			se, cok := err.(common.ServiceException)
			if cok {
				ctx.Output.SetStatus(400)
				ctx.Output.JSON(se, false, false)
			} else {
				ctx.Output.SetStatus(500)
				ctx.Output.JSON(common.NewServiceError(10000), false,false)
			}
			return

		}
	}

	// 增加初始化 orm操作对象的钩子函数， 这个对象初始化之前要保证 model注册完
	beego.AddAPPStartHook(daos.OrmInitHockFunc)

	beego.Run()
}
