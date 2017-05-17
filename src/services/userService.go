package services

import (
	"models"
)
// 示例代码 service 无状态 不保存数据，只做逻辑
type UserService struct {
	userDao models.UserDao
}

func (s *UserService) FindById(id int64) UserDTO {
	user := s.userDao.GetNewModel()
	user.Id = id
	models.FindOne(user)
	return UserDTO{user.Id, user.Name}
}


