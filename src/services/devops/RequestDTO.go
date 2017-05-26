package devops






// 列表查询请求结构
type ServerListRequest struct {
	Page int
	PageSize int
	Application string
	EngineRoom string
	Env string
	Ip string
}






