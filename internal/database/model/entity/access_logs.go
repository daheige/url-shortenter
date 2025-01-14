// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AccessLogs is the golang structure for table access_logs.
type AccessLogs struct {
	Id         uint64      `json:"id"         description:"ID"`
	AccountNo  uint64      `json:"accountNo"  description:"账号标识"`
	ShortNo    uint64      `json:"shortNo"    description:"短链标识"`
	ShortUrl   string      `json:"shortUrl"   description:"短链内容"`
	ShortAll   string      `json:"shortAll"   description:"带参数 URL"`
	YearTime   uint        `json:"yearTime"   description:"年份"`
	MonthTime  uint        `json:"monthTime"  description:"月份"`
	DayTime    uint        `json:"dayTime"    description:"日期"`
	AccessDate *gtime.Time `json:"accessDate" description:"访问日期"`
	AccessTime *gtime.Time `json:"accessTime" description:"访问时间"`
	UserAgent  string      `json:"userAgent"  description:"访问 user_agent"`
	Ip         string      `json:"ip"         description:"访问 IP"`
	TraceId    string      `json:"traceId"    description:"链路追踪标识"`
	VisitState uint        `json:"visitState" description:"访问状态 0 默认，100 正常 200 失效"`
	ServerIp   string      `json:"serverIp"   description:"服务器 IP"`
	CreateTime *gtime.Time `json:"createTime" description:"创建时间"`
	ModifyTime *gtime.Time `json:"modifyTime" description:"修改时间"`
}
