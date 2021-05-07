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

func (r *Repository) Save(entity Entity) (int64, ftype.Error) {
	var a001 bool
	var e error
	a001, e = r.Exists(entity)
	if nil != e {
		return 0, e
	}
	if a001 {
		// 已经存在
		var a002 string
		a002 = r.sqlUpdate + " " + r.sqlWhere
		flog.Debug(a002)
		var a003 []interface{}
		a003 = make([]interface{}, 0, 0xff)
		var entityValue reflect.Value
		entityValue = reflect.ValueOf(entity)
		for _, x := range r.columns {
			if !x.IsPrimaryKey ||
				x.KeyGenerate != AUTO {

				a003 = append(a003, entityValue.FieldByName(x.Name).Interface())
			}
		}
		for _, x := range r.columns {
			if x.IsPrimaryKey {
				a003 = append(a003, entityValue.FieldByName(x.Name).Interface())
			}
		}
		var connection *sql.Conn
		connection, e = r.db.Conn(r.ctx)
		if nil != e {
			return 0, ferror.New(e)
		}
		defer connection.Close()
		var result sql.Result
		result, e = connection.ExecContext(r.ctx, a002, a003...)
		if nil != e {
			return 0, ferror.New(e)
		}
		var a004 int64
		a004, e = result.RowsAffected()
		if nil != e {
			return a004, ferror.New(e)
		}
		return a004, nil
	} else {
		// 不存在
		var a002 string
		a002 = r.sqlInsert
		flog.Debug(a002)
		var a003 []interface{}
		a003 = make([]interface{}, 0, 0xff)
		var entityValue reflect.Value
		entityValue = reflect.ValueOf(entity)
		for _, x := range r.columns {
			if !x.IsPrimaryKey ||
				x.KeyGenerate != AUTO {

				a003 = append(a003, entityValue.FieldByName(x.Name).Interface())
			}
		}
		var connection *sql.Conn
		connection, e = r.db.Conn(r.ctx)
		if nil != e {
			return 0, ferror.New(e)
		}
		defer connection.Close()
		var result sql.Result
		result, e = connection.ExecContext(r.ctx, a002, a003...)
		if nil != e {
			return 0, ferror.New(e)
		}
		var a004 int64
		a004, e = result.RowsAffected()
		if nil != e {
			return a004, ferror.New(e)
		}
		return a004, nil
	}
}

func (r *Repository) SaveAll(entities ...Entity) (int64, ftype.Error) {
	var result int64
	var sb001 *fstrings.StringBuffer
	sb001 = fstrings.NewStringBuffer()
	for _, x := range entities {
		var a001 int64
		var e error
		a001, e = r.Save(x)
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
