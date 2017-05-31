package devops

import (
	"common"
	"services"
	"daos/devops"
	"fmt"
)

type ServingService struct {
	servingDao devops.ServingDao
	serverDao devops.ServerDao
	serverServingDao devops.ServerServingDao
}


func (this ServingService) Create(dto ServingCreateDTO) {
	m := &devops.Serving{}
	common.Convert(dto, m)
	this.servingDao.Insert(m)
	// 创建关联关系
	this.createServerServing(dto, m)
}

func (this ServingService) createServerServing(dto ServingCreateDTO, serving *devops.Serving) {
	if len(dto.ServerIds) > 0 {
		servers := this.serverDao.FindByIds(dto.ServerIds)
		if len(servers) > 0 {
			models := []devops.ServerServing{}
			for _, server := range servers {
				models = append(models,
					devops.ServerServing{ServerId:server.Id,
						ServingId:serving.Id, Application:serving.Application,
						Env:server.Env, EngineRoom:server.EngineRoom,
						IntranetIp:server.IntranetIp, ExtranetIp: server.ExtranetIp,
					})
			}
			this.serverServingDao.InsertAll(models)
		}
	}
}

func (this ServingService) FindById(id int64) services.ResultVO {
	m := &devops.Serving{}
	this.servingDao.FindOne(id, m)
	dto := ServingDetailDTO{}
	common.Convert(*m, &dto)
	serverServings := this.serverServingDao.FindByServingId(id)
	if len(serverServings) > 0 {
		serverids := []int64{}
		for _, s := range serverServings {
			serverids = append(serverids, s.ServerId)
		}
		servers := this.serverDao.FindByIds(serverids)
		simpleDTOs := []SimpleServerDTO{}
		for _, server := range servers {
			sdto := SimpleServerDTO{}
			common.Convert(server, &sdto)
			simpleDTOs = append(simpleDTOs, sdto)
		}
		dto.Servers = simpleDTOs
	} else {dto.Servers = []SimpleServerDTO{}}
	return services.NewResultVO(dto)
}





func (this ServingService) FindServings(request ListRequest) services.ResultPageVO {
	// 需要关联 server 信息
	if len(request.EngineRoom)>0 || len(request.Env)>0 || len(request.Ip)> 0 {
		return this.findServingList(request)
	}
	servings, count := this.servingDao.FindServingList(request.Application, request.Page, request.PageSize)
	res := []ServingDTO{}
	for _, s:= range servings {
		d := ServingDTO{}
		common.Convert(s, &d)
		res = append(res, d)
	}
	return services.ResultPageVO{Results:res, Count:count}
}
func (this ServingService) findServingList(request ListRequest) services.ResultPageVO{
	models, count := this.serverServingDao.FindServingIdList(request.Application, request.EngineRoom, request.Env,
		request.Ip, request.Page, request.PageSize)
	result := services.ResultPageVO{Count:count}
	servingIds := []int64{}
	for _, model := range models {
		servingIds = append(servingIds, model.ServingId)
	}
	fmt.Println(models)
	servings := this.FindByIds(servingIds)

	result.Results = servings
	return result
}


func (this ServingService) FindByIds(ids []int64) []ServingDTO {
	servings := this.servingDao.FindByIds(ids)
	res := []ServingDTO{}
	for _, s:= range servings {
		d := ServingDTO{}
		common.Convert(s, &d)
		res = append(res, d)
	}
	return res

}



func (this ServingService) Update(dto ServingCreateDTO) {
	serving := &devops.Serving{}
	this.servingDao.FindOne(dto.Id, serving)
	common.Convert(dto, serving)
	this.servingDao.Update(serving)

	this.serverServingDao.DeleteByServingIds([]int64{serving.Id})
	this.createServerServing(dto, serving)

}


func (this ServingService) Remove(ids []int64) {
	m := devops.Serving{}
	this.servingDao.RemoveByIds(ids, &m)
	this.serverServingDao.DeleteByServingIds(ids)
}

