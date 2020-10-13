package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	historyFieldNames          = builderx.FieldNames(&History{})
	historyRows                = strings.Join(historyFieldNames, ",")
	historyRowsExpectAutoSet   = strings.Join(stringx.Remove(historyFieldNames, "create_time", "update_time"), ",")
	historyRowsWithPlaceHolder = strings.Join(stringx.Remove(historyFieldNames, "msg", "create_time", "update_time"), "=?,") + "=?"

	cacheHistoryMsgPrefix = "cache#History#msg#"
)

type (
	HistoryModel struct {
		sqlc.CachedConn
		table string
	}

	History struct {
		Msg      string    `db:"msg"`
		LastEcho time.Time `db:"last_echo"`
		Times    uint64    `db:"times"`
	}
)

func NewHistoryModel(conn sqlx.SqlConn, c cache.CacheConf, table string) *HistoryModel {
	return &HistoryModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      table,
	}
}

func (m *HistoryModel) Insert(data History) (sql.Result, error) {
	query := `insert into ` + m.table + ` (` + historyRowsExpectAutoSet + `) values (?, ?, ?) on duplicate key update times=times+1 `
	return m.ExecNoCache(query, data.Msg, data.LastEcho, data.Times)
}

func (m *HistoryModel) FindOne(msg string) (*History, error) {
	historyMsgKey := fmt.Sprintf("%s%v", cacheHistoryMsgPrefix, msg)
	var resp History
	err := m.QueryRow(&resp, historyMsgKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := `select ` + historyRows + ` from ` + m.table + ` where msg = ? limit 1`
		return conn.QueryRow(v, query, msg)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *HistoryModel) Update(data History) error {
	historyMsgKey := fmt.Sprintf("%s%v", cacheHistoryMsgPrefix, data.Msg)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := `update ` + m.table + ` set ` + historyRowsWithPlaceHolder + ` where msg = ?`
		return conn.Exec(query, data.LastEcho, data.Times, data.Msg)
	}, historyMsgKey)
	return err
}

func (m *HistoryModel) Delete(msg string) error {
	historyMsgKey := fmt.Sprintf("%s%v", cacheHistoryMsgPrefix, msg)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := `delete from ` + m.table + ` where msg = ?`
		return conn.Exec(query, msg)
	}, historyMsgKey)
	return err
}

func (m *HistoryModel) Last() (*History, error) {
	var resp History

	err := m.QueryRowNoCache(&resp, `SELECT * FROM echo_history ORDER BY last_echo DESC LIMIT 1 `)

	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
