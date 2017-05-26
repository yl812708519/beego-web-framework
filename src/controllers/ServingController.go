package controllers

import (
	"services/devops"
	"common"
)

type ServingController struct {
	BaseController
	servingService devops.ServingService

}


// @Title 创建服务
// @Description
// @Param    tag            formData    string          true        "名称标签"
// @Param    serverIds      formData    []int64         false       "服务器ids 数组"
// @Param    application    formData    string          true        "应用"
// @Param    url            formData    string          false       "URL"
// @Param    version        formData    string          true        "版本"
// @Param    dependency     formData    string          false       "依赖"
// @Param    remark         formData    string          false       "备注"
// @Success 200 {status: success}
// @Failure 400 service exception
// @router /servings [post]
func (this *ServingController) Create() {
	s := devops.ServingCreateDTO{}
	this.parseJsonRequest(&s)
	this.valid(s)
	this.servingService.Create(s)
	this.renderSuccess()
}



// @Title 查找一个服务
// @Description
// @Param    id       query    int64          true       "service id"
// @Success 200 {object} serving.ServingDTO
// @Failure 400 service exception
// @router /serving/:id [get]
func (this *ServingController) FindOne() {
	id, err := this.GetInt64(":id")
	if err != nil{
		panic(common.NewServiceException(20003))
	}
	serverDTO := this.servingService.FindById(id)
	this.renderJSON(serverDTO)

}


