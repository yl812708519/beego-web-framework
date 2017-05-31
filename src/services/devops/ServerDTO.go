package devops



type ServerDTO struct {
	Id int64                `json:"id"`
	ServerId string         `json:"serverId"`
	Tag string              `valid:"Required;MaxSize(50)" json:"tag"`
	Application string      `json:"application"`
	Env string              `json:"env"`
	EngineRoom string       `json:"engineRoom" valid:"Required"`
	Core int                `json:"core" valid:"Required"`
	Memory int              `json:"memory" valid:"Required"`
	IntranetIp string       `json:"intranetIp" valid:"IP"`
	ExtranetIp string       `json:"extranetIp" valid:"IP"`
	Remark string           `json:"remark" valid:"MaxSize(200)"`

	Users []UserDTO            `json:"users"`
	Disks []DiskDTO            `json:"disks"`

	IsDeleted bool          `json:"-"`
	CreatedDate string      `json:"createdDate"`
	UpdatedDate string      `json:"updatedDate"`

}

type DiskDTO struct {
	Id int64                `json:"id"`
	ServerId int64          `json:"serverId"`
	RootPath string         `json:"rootPath" valid:"Required"`
	Size int                `json:"size" valid:"Required"`
}


type UserDTO struct {
	Id int64                `json:"id"`
	ServerId int64          `json:"serverId"`
	UserName string         `json:"userName" valid:"Required"`
	Password string         `json:"password" valid:"Required"`
}




type SimpleServerDTO struct {
	Id int64                `json:"id"`
	ServerId string         `json:"serverId"`
	Tag string              `valid:"Required;MaxSize(50)"`
	Env string              `json:"env"`
	EngineRoom string       `json:"engineRoom"`
	Core int                `json:"core"`
	Memory int              `json:"memory"`
	ExtranetIp string       `json:"extranetIp"`
}


