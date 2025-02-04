// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ShortAuditLog is the golang structure for table short_audit_log.
type ShortAuditLog struct {
	Id                 uint64      `json:"id"                 description:"ID"`
	ShortNo            uint64      `json:"shortNo"            description:"短链标识"`
	TrxId              uint64      `json:"trxId"              description:"记录 ID"`
	FullScreenshot     string      `json:"fullScreenshot"     description:"整屏镜像"`
	Content            string      `json:"content"            description:"网页内容"`
	HashContent        string      `json:"hashContent"        description:"hash 值 sha256"`
	ContentPath        string      `json:"contentPath"        description:"内容文件地址"`
	SafetyAuditAlibaba string      `json:"safetyAuditAlibaba" description:"阿里内容安全审核"`
	SafetyAuditTencent string      `json:"safetyAuditTencent" description:"腾讯内容审核"`
	AuditState         uint        `json:"auditState"         description:"审核状态 0 默认 10000 阿里异常，20000 腾讯异常，100000 俩者都异常"`
	HashState          uint        `json:"hashState"          description:"hash 状态 0 默认 100 正常 200 失效"`
	RedirectState      uint        `json:"redirectState"      description:"跳转状态 0 默认 100 正常 200 异常"`
	ModifyTime         *gtime.Time `json:"modifyTime"         description:"修改时间"`
	CreateTime         *gtime.Time `json:"createTime"         description:"创建时间"`
}
