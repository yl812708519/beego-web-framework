package devops



type ServerDTO struct {
	Id int64                `form:"_"`
	ServerId string         `form:"_"`
	Tag string              `valid:"Required;MaxSize(50)" form:"tag"`
	Application string      `form:"_"`
	Env string              `form:"env"`
	EngineRoom string       `form:"engineRoom" valid:"Required"`
	Core int                `form:"core" valid:"Required"`
	Memory int              `form:"memory" valid:"Required"`
	IntranetIp string       `form:"intranetIp" valid:"IP"`
	ExtranetIp string       `form:"extranetIp" valid:"IP"`
	Remark string           `form:"remark" valid:"MaxSize(200)"`

	Users []UserDTO            `form:"users"`
	Disks []DiskDTO            `form:"disks"`

	IsDeleted bool          `form:"_"`
	CreatedDate string      `form:"_"`
	UpdatedDate string      `form:"_"`

}

type DiskDTO struct {
	Id int64                `form:"_"`
	ServerId int64          `form:"_"`
	RootPath string         `form:"rootPath" valid:"Required"`
	Size int                `form:"size" valid:"Required"`
}


type UserDTO struct {
	Id int64                `form:"_"`
	ServerId int64          `form:"_"`
	UserName string         `form:"userName" valid:"Required"`
	Password string         `form:"password" valid:"Required"`
}


