package crud

import (
"reflect"
	"database/sql"
	"fmt"
	"github.com/iceyee/go-farmer/farmer"
	//
)

// 条件查询单个记录, 靠主键来识别,
// 返回Entity, error.
func Find(db *sql.DB, table Table, entity Entity) (interface{}, error) {
	// sql1 - sql语句
	// statement1
	// args - Query()的传参(type []interface{})
	// rows1 - Query()之后返回的东西
	var sql1 string = table.Select + " " + table.Where
	if verbose {
		println(sql1)
	}
	statement1, e := db.Prepare(sql1)
	if nil != e {
		return nil, farmer.NewFarmerError(e)
	}
	defer statement1.Close()
	var args = arguments2(table, entity)
	if verbose {
		fmt.Printf("%v\n", args)
	}
	rows1, e := statement1.Query(args...)
	if nil != e {
		return nil, farmer.NewFarmerError(e)
	}

	if !rows1.Next() {
		return nil, nil
	}
	result1, args := arguments3(table)
	e = rows1.Scan(args...)
	if nil != e {
		return nil, farmer.NewFarmerError(e)
	}
	return reflect.ValueOf(result1).Elem().Interface(), nil
}

// 条件查询多个记录,
// 返回[]Entity, error.
func FindAll(db *sql.DB, table Table, where string, args ...interface{}) ([]interface{}, error) {
	// sql1 - sql语句
	// statement1
	// args - Query(),Scan()的传参
	// rows1 - Query()之后返回的东西
	// result1 - 输出
	// result2 - 单个输出, 最后合并到result1
	var sql1 string = table.Select + " " + where
	if verbose {
		println(sql1)
		fmt.Printf("%v\n", args)
	}
	statement1, e := db.Prepare(sql1)
	if nil != e {
		return nil, farmer.NewFarmerError(e)
	}
	defer statement1.Close()
	rows1, e := statement1.Query(args...)
	if nil != e {
		return nil, farmer.NewFarmerError(e)
	}

	var result1 []interface{} = make([]interface{}, 0, 0xff)
	for rows1.Next() {
		result2, args := arguments3(table)
		e = rows1.Scan(args...)
		if nil != e {
			return nil, farmer.NewFarmerError(e)
		}
		result1 = append(result1, reflect.ValueOf(result2).Elem().Interface())
	}
	return result1, nil
}
