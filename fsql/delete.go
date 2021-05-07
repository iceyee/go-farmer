package fsql

import (
	"database/sql"
	"errors"
	"github.com/iceyee/go-farmer/v3/ferror"
	"github.com/iceyee/go-farmer/v3/flog"
	"github.com/iceyee/go-farmer/v3/fstrings"
	"github.com/iceyee/go-farmer/v3/ftype"
	"reflect"
	//
)

func (r *Repository) Delete(entity Entity) (int64, ftype.Error) {
	var a001 string
	a001 = r.sqlDelete + " " + r.sqlWhere
	flog.Debug(a001)
	var a002 []interface{}
	a002 = make([]interface{}, 0, 0xff)
	var entityValue reflect.Value
	entityValue = reflect.ValueOf(entity)
	for _, x := range r.columns {
		if x.IsPrimaryKey {
			a002 = append(a002, entityValue.FieldByName(x.Name).Interface())
		}
	}
	var connection *sql.Conn
	var e error
	connection, e = r.db.Conn(r.ctx)
	if nil != e {
		return 0, ferror.New(e)
	}
	defer connection.Close()
	var result sql.Result
	result, e = connection.ExecContext(r.ctx, a001, a002...)
	if nil != e {
		return 0, ferror.New(e)
	}
	var a003 int64
	a003, e = result.RowsAffected()
	if nil != e {
		return a003, ferror.New(e)
	}
	return a003, nil
}

func (r *Repository) DeleteAll(entities ...Entity) (int64, ftype.Error) {
	var result int64
	var sb001 *fstrings.StringBuffer
	sb001 = fstrings.NewStringBuffer()
	for _, x := range entities {
		var a001 int64
		var e error
		a001, e = r.Delete(x)
		if nil != e {
			sb001.Append("\n")
			sb001.Append(e)
		} else {
			result += a001
		}
	}
	var a002 string
	a002 = sb001.String()
	if "" == a002 {
		return result, nil
	} else {
		return result, errors.New(a002)
	}
}
