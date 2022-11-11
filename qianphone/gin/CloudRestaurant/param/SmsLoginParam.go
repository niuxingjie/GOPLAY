package param

// 用于解析前端传来的参数
type SmsLoginParam struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}
