// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html

package routers

import (
	"devops/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//ns := beego.NewNamespace("/v1",
	//	beego.NSNamespace("/object",
	//		beego.NSInclude(
	//			&controllers.ObjectController{},
	//		),
	//	),
	//	beego.NSNamespace("/user",
	//		beego.NSInclude(
	//			&controllers.UserController{},
	//		),
	//	),
	//)
	//beego.AddNamespace(ns)
	// 使用自动化文档一定要使用 namespace 的方式



	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/devops",
			beego.NSInclude(&controllers.ServerController{}),
			beego.NSInclude(&controllers.ServingController{}),
		),

	)


	beego.AddNamespace(ns)




}
