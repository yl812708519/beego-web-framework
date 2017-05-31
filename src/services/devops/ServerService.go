package devops

import (
	"common"
	"services"
	"log"
	"daos/devops"
)


type ServerService struct {
	serverDao     devops.ServerDao
	serverUserDao devops.ServerUserDao
	serverDiskDao devops.ServerDiskDao
}


func (this ServerService) FindById(id int64) ServerDTO{
	server := &devops.Server{}
	this.serverDao.FindOne(id, server)
	serverDTO := &ServerDTO{}
	common.Convert(*server, serverDTO)
	users := this.serverUserDao.FindByServerIds([]int64{id})
	disks := this.serverDiskDao.FindByServerIds([]int64{id})
	for _, u := range users {
		serverDTO.Users = append(serverDTO.Users, UserDTO{Id:u.Id, UserName:u.UserName, Password:u.Password, ServerId:u.ServerId})
	}
	for _, d := range disks {
		serverDTO.Disks = append(serverDTO.Disks, DiskDTO{Id:d.Id, RootPath:d.RootPath, Size:d.Size, ServerId:d.ServerId})
	}

	return *serverDTO
}


func (this ServerService) Create(request ServerDTO) {
	server := &devops.Server{}
	common.Convert(request, server)
	server.ServerId = common.GenUUID()
	id := this.serverDao.Insert(server)

	users, disks := this.getUserAndDiskByServerDTO(request, id)
	this.serverUserDao.InsertAll(users)
	this.serverDiskDao.InsertAll(disks)

}

func (this ServerService) Update(request ServerDTO) {
	// 判断是否有数据
	s := &devops.Server{}
	this.serverDao.FindOne(request.Id, s)
	// 更新
	server := &devops.Server{}
	common.Convert(request, server)
	server.ServerId = s.ServerId
	this.serverDao.Update(server)

	// 删除原来的
	this.serverDiskDao.DeleteByServerId(request.Id)
	this.serverUserDao.DeleteByServerId(request.Id)

	users, disks := this.getUserAndDiskByServerDTO(request, request.Id)
	// 根据请求创建所有的
	this.serverDiskDao.InsertAll(disks)
	this.serverUserDao.InsertAll(users)
}


func (this ServerService) getUserAndDiskByServerDTO(request ServerDTO, serverId ...int64) ([]devops.ServerUser, []devops.ServerDisk) {
	var users []devops.ServerUser
	var disks []devops.ServerDisk
	id := request.Id
	if id == 0 {
		if len(serverId) == 0 {
			log.Println("转换缺少serverId")
			panic(common.NewServiceError(20003))
		}
		id = serverId[0]
	}
	for _, u := range request.Users {
		serverUser := devops.ServerUser{ServerId: id, UserName: u.UserName, Password: u.Password}
		users = append(users, serverUser)
	}
	for _, d := range request.Disks {
		serverDisk := devops.ServerDisk{ServerId: id, RootPath: d.RootPath, Size: d.Size}
		disks = append(disks, serverDisk)
	}
	return users, disks
}

func (this ServerService) FindList(request ListRequest) services.ResultPageVO {

	servers, count := this.serverDao.FindList(request.Application, request.EngineRoom,
		request.Env, request.Ip, request.Page, request.PageSize)
	if len(servers) <= 0 {
		return services.ResultPageVO{Results:[]ServerDTO{}, Count:count}
	}
	var serverIds []int64
	for _, s := range servers {
		serverIds = append(serverIds, s.Id)
	}
	users := this.serverUserDao.FindByServerIds(serverIds)
	disks := this.serverDiskDao.FindByServerIds(serverIds)
	//fmt.Println(common.Converts(UserDTO{}, users...))
	dtos := this.convertServer(servers, users, disks)

	return services.ResultPageVO{dtos, count}
}
func (this ServerService) convertServer(servers []devops.Server, users []devops.ServerUser,
		disks []devops.ServerDisk) []ServerDTO {
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
	return dtos
}

func (this ServerService) FindByIds(ids []int64) []ServerDTO {
	if len(ids) <= 0 {
		panic(common.NewServiceException(20005))
	}
	servers := this.serverDao.FindByIds(ids)
	users := this.serverUserDao.FindByServerIds(ids)
	disks := this.serverDiskDao.FindByServerIds(ids)
	return this.convertServer(servers, users, disks)

}

func (this ServerService) Remove(id int64){
	this.serverDao.Remove(id, &devops.Server{})
	this.serverDiskDao.DeleteByServerId(id)
	this.serverUserDao.DeleteByServerId(id)

}

func (this ServerService) RemoveByIds(ids []int64){
	this.serverDao.RemoveByIds(ids)
	this.serverDiskDao.DeleteByServerIds(ids)
	this.serverUserDao.DeleteByServerIds(ids)
}
