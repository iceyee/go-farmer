package fsql

import (
	"database/sql"
	"github.com/iceyee/go-farmer/v5/ferror"
	"github.com/iceyee/go-farmer/v5/ftype"
	//
)

func (r *Repository) Ping() ftype.Error {
	var connection *sql.Conn
	var e error
	connection, e = r.db.Conn(r.ctx)
	if nil != e {
		return ferror.New(e)
	}
	defer connection.Close()
	e = connection.PingContext(r.ctx)
	if nil != e {
		return ferror.New(e)
	}
	return nil
}
