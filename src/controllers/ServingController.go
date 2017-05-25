package controllers




type ServingController struct {
	BaseController

}


// @Title create server
// @Description
// @Param    serverId       formData    int64           false       "外网ip"
// @Param    tag            formData    string          true        "名称标签"
// @Param    application    formData    string          true        "应用"
// @Param    url            formData    string          false       "URL"
// @Param    version        formData    string          true        "版本"
// @Param    dependency     formData    string          false       "依赖"
// @Param    env            formData    string          false       "环境"
// @Param    remark         formData    string          false       "备注"
// @Success 200 {status: success}
// @Failure 400 service exception
// @router /servers [post]
func (this ServingController) Create() {



	this.renderSuccess()
}





