package devops


type ServingCreateDTO struct {

	Id          int64       `json:"id"`
	Tag         string      `json:"tag" valid:"Required"`
	Application string      `json:"application" valid:"Required"`
	ServerIds   []int64     `json:"serverIds" `
	Url         string      `json:"url"`
	Version     string      `json:"version"`
	Dependency  string      `json:"dependency"`
	Remark      string      `json:"remark"`
	IsDeleted   bool        `json:"isDeleted"`
	CreatedDate   string    `json:"createdDate"`
	UpdatedDate   string    `json:"updatedDate"`
}

type ServingDTO struct {
	Id          int64               `json:"id"`
	Tag         string              `json:"tag"`
	Application string              `json:"application"`
	Url         string              `json:"url"`
	Version     string              `json:"version"`
	Dependency  string              `json:"dependency"`
	Remark      string              `json:"remark"`
	IsDeleted   bool                `json:"isDeleted"`
	CreatedDate   string            `json:"createdDate"`
	UpdatedDate   string            `json:"updatedDate"`
	Servers     []ServerDTO  `json:"servers"`
}



