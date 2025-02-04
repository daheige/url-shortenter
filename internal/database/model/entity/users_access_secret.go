// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UsersAccessSecret is the golang structure for table users_access_secret.
type UsersAccessSecret struct {
	Id         uint64      `json:"id"         description:"ID"`
	AccountNo  uint64      `json:"accountNo"  description:"企业管理员 ID 一致"`
	SecretId   string      `json:"secretId"   description:"授权应用 ID"`
	SecretKey  string      `json:"secretKey"  description:"授权应用 key"`
	Salt       string      `json:"salt"       description:"盐值"`
	SaltKey    string      `json:"saltKey"    description:"盐值 key"`
	GrantType  string      `json:"grantType"  description:"授权类型：默认空，API token 授权:client_credentials"`
	State      uint        `json:"state"      description:"状态 0 默认 100 正常 200 禁用"`
	CreateTime *gtime.Time `json:"createTime" description:"创建时间"`
	ModifyTime *gtime.Time `json:"modifyTime" description:"修改时间"`
}
