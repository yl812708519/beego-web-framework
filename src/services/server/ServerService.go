package server

import (
	"models"
	"common"
)

type ServerService struct {
	serverDao models.ServerDao
}


func (s ServerService) FindById(id int64) ServerDTO{

	server := s.serverDao.FindOne(id)
	serverDTO := &ServerDTO{}
	common.Convert(server, serverDTO)
	return *serverDTO
}



