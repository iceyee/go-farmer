package fsql

import (
	"database/sql"
	"github.com/iceyee/go-farmer/v3/ferror"
	"github.com/iceyee/go-farmer/v3/flog"
	"github.com/iceyee/go-farmer/v3/ftype"
	"reflect"
	//
)

func (r *Repository) Find(entity Entity) (
	interface{}, ftype.Error) {

	var a001 []interface{}
	a001 = make([]interface{}, 0, 0xff)
	var entityValue reflect.Value
	entityValue = reflect.ValueOf(entity)
	for _, x := range r.columns {
		if x.IsPrimaryKey {
			a001 = append(a001, entityValue.FieldByName(x.Name).Interface())
		}
	}
	return r.FindOne(r.sqlWhere, a001...)
}

func (r *Repository) FindOne(where string, args ...interface{}) (
	interface{}, ftype.Error) {

	var result []interface{}
	var e error
	result, e = r.FindAll(where, args...)
	if nil != e {
		return nil, e
	} else if 0 == len(result) {
		return nil, nil
	} else {
		return result[0], nil
	}
}

func (r *Repository) FindAll(where string, args ...interface{}) (
	[]interface{}, ftype.Error) {

	var a001 string
	a001 = r.sqlSelect + " " + where
	flog.Debug(a001)
	var connection *sql.Conn
	var e error
	connection, e = r.db.Conn(r.ctx)
	if nil != e {
		return nil, ferror.New(e)
	}
	defer connection.Close()
	var rows *sql.Rows
	rows, e = connection.QueryContext(r.ctx, a001, args...)
	if nil != e {
		return nil, ferror.New(e)
	}
	defer rows.Close()
	var result []interface{}
	result = make([]interface{}, 0, 0xff)
	for rows.Next() {
		var a002 []interface{}
		a002 = make([]interface{}, 0, 0xff)
		var entityValue reflect.Value
		entityValue = reflect.New(r.entityType).Elem()
		for _, x := range r.columns {
			a002 = append(a002, entityValue.FieldByName(x.Name).Addr().Interface())
		}
		e = rows.Scan(a002...)
		if nil != e {
			return nil, ferror.New(e)
		}
		result = append(result, entityValue.Interface())
	}
	return result, nil
}
