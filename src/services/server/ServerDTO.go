package server



type ServerDTO struct {
	Id int64                `form:"_"`
	ServerId string         `form:"_"`
	Tag string              `form:"tag"`
	Application string      `form:"_"`
	Env string              `form:"env"`
	EngineRoom string       `form:"engineRoom"`
	Core int                `form:"core"`
	Memory int              `form:"memory"`
	IntranetIp string       `form:"intranetIp"`
	ExtranetIp string       `form:"extranetIp"`
	Remark string           `form:"remark"`

	Users []UserDTO            `form:"users"`
	Disks []DiskDTO            `form:"disks"`

	IsDeleted bool          `form:"_"`
	CreatedDate string      `form:"_"`
	UpdatedDate string      `form:"_"`

}

type DiskDTO struct {
	Id int64                `form:"_"`
	ServerId int64          `form:"_"`
	RootPath string         `form:"rootPath"`
	Size int                `form:"size"`
}


type UserDTO struct {
	Id int64                `form:"_"`
	ServerId int64          `form:"_"`
	UserName string         `form:"userName"`
	Password string         `form:"password"`
}


