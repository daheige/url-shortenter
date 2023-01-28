// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AccessLogsDao is the data access object for table access_logs.
type AccessLogsDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns AccessLogsColumns // columns contains all the column names of Table for convenient usage.
}

// AccessLogsColumns defines and stores column names for table access_logs.
type AccessLogsColumns struct {
	Id         string // ID
	AccountNo  string // 账号标识
	ShortNo    string // 短链标识
	ShortUrl   string // 短链内容
	ShortAll   string // 带参数URL
	YearTime   string // 年份
	MonthTime  string // 月份
	DayTime    string // 日期
	AccessDate string // 访问日期
	AccessTime string // 访问时间
	UserAgent  string // 访问user_agent
	Ip         string // 访问IP
	TraceId    string // 链路追踪标识
	VisitState string // 访问状态 0默认，100正常 200失效
	ServerIp   string // 服务器IP
	CreateTime string // 创建时间
	ModifyTime string // 修改时间
}

//  accessLogsColumns holds the columns for table access_logs.
var accessLogsColumns = AccessLogsColumns{
	Id:         "id",
	AccountNo:  "account_no",
	ShortNo:    "short_no",
	ShortUrl:   "short_url",
	ShortAll:   "short_all",
	YearTime:   "year_time",
	MonthTime:  "month_time",
	DayTime:    "day_time",
	AccessDate: "access_date",
	AccessTime: "access_time",
	UserAgent:  "user_agent",
	Ip:         "ip",
	TraceId:    "trace_id",
	VisitState: "visit_state",
	ServerIp:   "server_ip",
	CreateTime: "create_time",
	ModifyTime: "modify_time",
}

// NewAccessLogsDao creates and returns a new DAO object for table data access.
func NewAccessLogsDao() *AccessLogsDao {
	return &AccessLogsDao{
		group:   "default",
		table:   "access_logs",
		columns: accessLogsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AccessLogsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AccessLogsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AccessLogsDao) Columns() AccessLogsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AccessLogsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AccessLogsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AccessLogsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
