package model

// HomeInput is a struct for input data
type HomeInput struct {
	Short      string `v:"required|min-length:2|max-length:256#请输入短链接|最小长度为{min}|最大长度为{max}" json:"short" dc:"shortUrl 短链接"`
	RawQuery   string `json:"rawQuery" dc:"rawQuery 原始查询"`
	ClientIP   string `json:"clientIp"`
	UserAgent  string `json:"userAgent"`
	Referer    string `json:"referer"`
	ShortAll   string `json:"shortAll"`
	TraceID    string `json:"traceId"`
	VisitState uint   `json:"visitState" dc:"visitState 访问状态 0默认 100正常 200失效"`
	Host       string `json:"host" dc:"host 域名 包括端口"`
}

// HomeOutput is a struct for output data
type HomeOutput string
