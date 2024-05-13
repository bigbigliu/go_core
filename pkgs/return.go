package pkgs

// ResultInfo http响应消息结构体
type ResultInfo struct {
	Code  string `json:"code" `
	Msg   string `json:"msg"`
	Total int64  `json:"total,omitempty"`
	Data  any    `json:"data"`
}

// ReqQuery 列表查询参数
type ReqQuery struct {
	Key    string `json:"key"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
}
