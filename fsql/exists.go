package fsql

import (
	"github.com/iceyee/go-farmer/v3/ftype"
	"reflect"
	//
)

func (r *Repository) Exists(entity Entity) (bool, ftype.Error) {
	var entityValue reflect.Value
	entityValue = reflect.ValueOf(entity)
	var a001 []interface{}
	a001 = make([]interface{}, 0, 0xff)
	for _, x := range r.columns {
		if x.IsPrimaryKey {
			a001 = append(a001, entityValue.FieldByName(x.Name).Interface())
		}
	}
	var result int64
	var e error
	result, e = r.Count(r.sqlWhere, a001...)
	if nil != e {
		return false, e
	} else {
		return 0 < result, nil
	}
}
