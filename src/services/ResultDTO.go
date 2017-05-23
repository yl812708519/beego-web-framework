package services





// 是否使用 之后待定

type ResultVO struct {
	result interface{}
}

type ResultPageVO struct{
	Results []interface{}   `json:"results"`
	Count   int64           `json:"count"`
}



func SetResult(r *ResultPageVO, values ...interface{}) {
	r.Results = values
}






