package serving

import (
	"models"
	"common"
	"services/server"
	"services"
)

type ServingService struct {
	dao models.ServerDao
	serverService server.ServerService
}


func (this ServingService) Create(dto ServingCreateDTO) {
	m := models.Serving{}
	common.Convert(dto, &m)
	this.dao.Insert(m)
}


func (this ServingService) FindById(id int64) services.ResultVO {
	m := models.Serving{}
	this.dao.FindOne(id, m)
	dto := ServingDTO{}
	common.Convert(m, &dto)
	return services.NewResultVO(dto)
}

