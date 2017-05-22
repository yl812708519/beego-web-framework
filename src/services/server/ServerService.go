package server

import (
	"models"
	"common"
)

type ServerService struct {
	serverDao models.ServerDao
}


func (this ServerService) FindById(id int64) ServerDTO{

	server := this.serverDao.FindOne(id)
	serverDTO := &ServerDTO{}
	common.Convert(server, serverDTO)
	return *serverDTO
}

func (this ServerService) Create(request ServerDTO) {
	server := &models.Server{}
	common.Convert(request, server)
	this.serverDao.Insert(*server)
}


