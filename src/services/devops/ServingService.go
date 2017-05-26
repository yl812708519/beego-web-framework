package devops

import (
	"common"
	"services"
	"daos/devops"
)

type ServingService struct {
	servingDao devops.ServingDao
	serverService ServerService
}


func (this ServingService) Create(dto ServingCreateDTO) {
	m := &devops.Serving{}
	common.Convert(dto, m)
	this.servingDao.Insert(m)
}


func (this ServingService) FindById(id int64) services.ResultVO {
	m := &devops.Serving{}
	this.servingDao.FindOne(id, m)
	dto := ServingDTO{}
	common.Convert(m, &dto)
	return services.NewResultVO(dto)
}

