package fsql

import (
	"database/sql"
	"github.com/iceyee/go-farmer/v4/ferror"
	"github.com/iceyee/go-farmer/v4/flog"
	"github.com/iceyee/go-farmer/v4/ftype"
	//
)

func (r *Repository) Count(where string, args ...interface{}) (int64, ftype.Error) {
	var a001 string
	a001 = "select count(*) from " + r.tableName + " " + where
	flog.Debug(a001)
	var connection *sql.Conn
	var e error
	connection, e = r.db.Conn(r.ctx)
	if nil != e {
		return 0, ferror.New(e)
	}
	defer connection.Close()
	var rows *sql.Rows
	rows, e = connection.QueryContext(r.ctx, a001, args...)
	if nil != e {
		return 0, ferror.New(e)
	}
	defer rows.Close()
	rows.Next()
	var result int64
	e = rows.Scan(&result)
	if nil != e {
		return 0, ferror.New(e)
	}
	return result, nil
}
