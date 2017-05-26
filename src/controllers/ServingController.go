package controllers

import "services/serving"

type ServingController struct {
	BaseController
	service serving.ServingService

}


// @Title 创建服务
// @Description
// @Param    serverId       formData    int64           false       "外网ip"
// @Param    tag            formData    string          true        "名称标签"
// @Param    serverIds      formData    []int64         false       "服务器ids 数组"
// @Param    application    formData    string          true        "应用"
// @Param    url            formData    string          false       "URL"
// @Param    version        formData    string          true        "版本"
// @Param    dependency     formData    string          false       "依赖"
// @Param    env            formData    string          false       "环境"
// @Param    remark         formData    string          false       "备注"
// @Success 200 {status: success}
// @Failure 400 service exception
// @router /servings [post]
func (this *ServingController) Create() {
	s := serving.ServingCreateDTO{}
	this.parseJsonRequest(&s)
	this.valid(s)
	this.service.Create(s)
	this.renderSuccess()
}



// @Title 查找一个服务
// @Description
// @Param    id       query    int64          true       "service id"
// @Success 200 {object} serving.ServingDTO
// @Failure 400 service exception
// @router /serving/:id [get]
func (this *ServingController) FindOne() {



	this.renderSuccess()
}


