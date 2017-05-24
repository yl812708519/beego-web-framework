package services


// 应用
var applications = []string{
	"car-league-app",
	"car-league-web",
	"car-mall-mobile",
	"car-visit",
	"car-finance-api",
	"car-finance-admin",
}

// 机房
var engineRooms = []string{
	"IDC",
	"阿里云",
	"UCloud",
}

// 环境
var envs = []string{
	"dev",
	"test",
	"beta",
	"stg",
	"prod",
}

var ConstantMap = map[string][]string{
	"applications": applications,
	"engineRooms": engineRooms,
	"envs": envs,
}





