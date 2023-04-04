// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ShortAuditLogDao is the data access object for table short_audit_log.
type ShortAuditLogDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns ShortAuditLogColumns // columns contains all the column names of Table for convenient usage.
}

// ShortAuditLogColumns defines and stores column names for table short_audit_log.
type ShortAuditLogColumns struct {
	Id                 string // ID
	ShortNo            string // 短链标识
	TrxId              string // 记录ID
	FullScreenshot     string // 整屏镜像
	Content            string // 网页内容
	HashContent        string // hash值 sha256
	ContentPath        string // 内容文件地址
	SafetyAuditAlibaba string // 阿里内容安全审核
	SafetyAuditTencent string // 腾讯内容审核
	AuditState         string // 审核状态 0 默认 10000阿里异常，20000 腾讯异常，100000俩者都异常
	HashState          string // hash状态 0 默认 100正常 200失效
	RedirectState      string // 跳转状态 0 默认 100正常 200异常
	ModifyTime         string // 修改时间
	CreateTime         string // 创建时间
}

// shortAuditLogColumns holds the columns for table short_audit_log.
var shortAuditLogColumns = ShortAuditLogColumns{
	Id:                 "id",
	ShortNo:            "short_no",
	TrxId:              "trx_id",
	FullScreenshot:     "full_screenshot",
	Content:            "content",
	HashContent:        "hash_content",
	ContentPath:        "content_path",
	SafetyAuditAlibaba: "safety_audit_alibaba",
	SafetyAuditTencent: "safety_audit_tencent",
	AuditState:         "audit_state",
	HashState:          "hash_state",
	RedirectState:      "redirect_state",
	ModifyTime:         "modify_time",
	CreateTime:         "create_time",
}

// NewShortAuditLogDao creates and returns a new DAO object for table data access.
func NewShortAuditLogDao() *ShortAuditLogDao {
	return &ShortAuditLogDao{
		group:   "default",
		table:   "short_audit_log",
		columns: shortAuditLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ShortAuditLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ShortAuditLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ShortAuditLogDao) Columns() ShortAuditLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ShortAuditLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ShortAuditLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ShortAuditLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
