package server

import (
	"models"
	"common"
	"services"
)


type ServerService struct {
	serverDao models.ServerDao
	serverUserDao models.ServerUserDao
	serverDiskDao models.ServerDiskDao
}


func (this ServerService) FindById(id int64) ServerDTO{
	server := &models.Server{}
	this.serverDao.FindOne(id, server)
	serverDTO := &ServerDTO{}
	common.Convert(*server, serverDTO)
	return *serverDTO
}


func (this ServerService) Create(request ServerDTO) {
	var users []models.ServerUser
	var disks []models.ServerDisk
	server := &models.Server{}
	common.Convert(request, server)
	server.ServerId = common.GenUUID()
	this.serverDao.Insert(server)

	for _, u := range request.Users {
		serverUser := models.ServerUser{ServerId:server.Id, UserName:u.UserName, Password:u.Password}
		users = append(users, serverUser)
	}
	this.serverUserDao.InsertAll(users)
	for _, d := range request.Disks {
		serverDisk := models.ServerDisk{ServerId:server.Id, RootPath:d.RootPath, Size: d.Size}
		disks = append(disks, serverDisk)
	}
	this.serverDiskDao.InsertAll(disks)

}

func (this ServerService) FindList(request ListRequest) services.ResultPageVO {

	servers, count := this.serverDao.FindList(request.Application, request.EngineRoom,
		request.Env, request.Ip, request.Page, request.PageSize)
	var serverIds []int64
	for _, s := range servers {
		serverIds = append(serverIds, s.Id)
	}
	users := this.serverUserDao.FindByServerIds(serverIds)
	disks := this.serverDiskDao.FindByServerIds(serverIds)

	var dtos []ServerDTO

	for _, s := range servers {
		sd := ServerDTO{}
		common.Convert(s, &sd)
		for _, u := range users {
			if u.ServerId == s.Id {
				sd.Users = append(sd.Users, UserDTO{Id:u.Id, UserName:u.UserName, Password:u.Password, ServerId:u.ServerId})
			}
		}
		for _, d := range disks {
			if d.ServerId == s.Id {
				sd.Disks = append(sd.Disks, DiskDTO{Id:d.Id, RootPath:d.RootPath, Size:d.Size, ServerId:d.ServerId})
			}
		}
		dtos = append(dtos, sd)
	}
	var result services.ResultPageVO
	services.SetResult(&result, dtos)
	result.Count = count
	return result
}


func (this ServerService) remove(id int64){



}


