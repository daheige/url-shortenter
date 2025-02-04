// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UrlIpCountStatsDao is the data access object for table url_ip_count_stats.
type UrlIpCountStatsDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns UrlIpCountStatsColumns // columns contains all the column names of Table for convenient usage.
}

// UrlIpCountStatsColumns defines and stores column names for table url_ip_count_stats.
type UrlIpCountStatsColumns struct {
	ShortUrl        string // 短链内容
	ShortNo         string // 短链的唯一 ID
	TodayCount      string //
	YesterdayCount  string //
	Last7DaysCount  string //
	MonthlyCount    string //
	TotalCount      string //
	DTodayCount     string //
	DYesterdayCount string //
	DLast7DaysCount string //
	DMonthlyCount   string //
	DTotalCount     string //
}

// urlIpCountStatsColumns holds the columns for table url_ip_count_stats.
var urlIpCountStatsColumns = UrlIpCountStatsColumns{
	ShortUrl:        "short_url",
	ShortNo:         "short_no",
	TodayCount:      "today_count",
	YesterdayCount:  "yesterday_count",
	Last7DaysCount:  "last_7_days_count",
	MonthlyCount:    "monthly_count",
	TotalCount:      "total_count",
	DTodayCount:     "d_today_count",
	DYesterdayCount: "d_yesterday_count",
	DLast7DaysCount: "d_last_7_days_count",
	DMonthlyCount:   "d_monthly_count",
	DTotalCount:     "d_total_count",
}

// NewUrlIpCountStatsDao creates and returns a new DAO object for table data access.
func NewUrlIpCountStatsDao() *UrlIpCountStatsDao {
	return &UrlIpCountStatsDao{
		group:   "default",
		table:   "url_ip_count_stats",
		columns: urlIpCountStatsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UrlIpCountStatsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UrlIpCountStatsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UrlIpCountStatsDao) Columns() UrlIpCountStatsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UrlIpCountStatsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UrlIpCountStatsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UrlIpCountStatsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
