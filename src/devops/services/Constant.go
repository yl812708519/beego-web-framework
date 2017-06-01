package services


type Option struct {
	Label string    `json:"label"`
	Value string    `json:"value"`
}



// 应用
var applications =[]Option{
	Option{"car-league-app", "car-league-app"},
	Option{"car-league-web", "car-league-web"},
	Option{"car-mall-mobile", "car-mall-mobile"},
	Option{"car-visit", "car-visit"},
	Option{"car-finance-api", "car-finance-api"},
	Option{"car-finance-admin", "car-finance-admin"},
}



// 机房
var engineRooms =  []Option{
	Option{"IDC", "IDC"},
	Option{"阿里云", "阿里云"},
	Option{"UCLoud", "UCLoud"},
}



// 环境
var envs = []Option{
	Option{"dev", "dev"},
	Option{"test", "test"},
	Option{"beta", "beta"},
	Option{"stg", "stg"},
	Option{"prod", "prod"},
}


var ConstantMap = map[string][]Option{
	"applications": applications,
	"engineRooms": engineRooms,
	"envs": envs,
}





