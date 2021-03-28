package crud

import (
	"database/sql"
	"fmt"
	"github.com/iceyee/go-farmer/v1/farmer"
	//
)

// 删除单个记录, 靠主键来识别,
// 返回affected, error.
func Delete(db *sql.DB, table Table, entity Entity) (int64, error) {
	// sql1 - sql语句
	// statement1 - statement
	// args - 主键值(type []interface{})
	// result1 - 执行Exec()后返回的东西
	var sql1 string = table.Delete + " " + table.Where
	if verbose {
		println(sql1)
	}
	statement1, e := db.Prepare(sql1)
	if nil != e {
		return 0, farmer.NewFarmerError(e)
	}
	defer statement1.Close()
	var args = arguments2(table, entity)
	if verbose {
		fmt.Printf("%v\n", args)
	}
	result1, e := statement1.Exec(args...)
	if nil != e {
		return 0, farmer.NewFarmerError(e)
	}
	return result1.RowsAffected()
}

// 删除多个记录, 靠各自的主键来识别,
// 返回affected, error.
func DeleteAll(db *sql.DB, table Table, entityA []Entity) (int64, error) {
	// sql1 - sql语句
	// statement1 - statement
	// args - 主键值(type []interface{})
	// result1 - 执行Exec()后返回的东西
	// affected - 累计
	// a1 - 临时变量
	var sql1 string = table.Delete + " " + table.Where
	if verbose {
		println(sql1)
	}
	statement1, e := db.Prepare(sql1)
	if nil != e {
		return 0, farmer.NewFarmerError(e)
	}
	defer statement1.Close()
	var affected int64
	for _, entity := range entityA {
		var args = arguments2(table, entity)
		if verbose {
			fmt.Printf("%v\n", args)
		}
		result1, e := statement1.Exec(args...)
		if nil != e {
			return affected, farmer.NewFarmerError(e)
		}
		a1, e := result1.RowsAffected()
		if nil != e {
			return affected, farmer.NewFarmerError(e)
		}
		affected += a1
	}
	return affected, nil
}
