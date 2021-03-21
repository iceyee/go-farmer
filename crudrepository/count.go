package crudrepository

import (
	// TODO
	//
	"database/sql"
	"github.com/iceyee/go-farmer/farmer"
)

// 条件查询符合条件的记录数
func Count(db *sql.DB, table Table, where string, args ...interface{}) (int64, error) {
	// sql1 - sql语句
	// statement1 - sql语句
	// rows1 - 查询返回的东西
	var sql1 string = "select count(*) from " + table.Name + " " + where
	//
	println(sql1)
	//
	statement1, e := db.Prepare(sql1)
	if nil != e {
		return 0, farmer.NewFarmerError(e)
	}
	defer statement1.Close()

	rows1, e := statement1.Query(args...)
	if nil != e {
		return 0, farmer.NewFarmerError(e)
	}
	if !rows1.Next() {
		return 0, farmer.NewFarmerError("1")
	}

	// result1 - 输出
	var result1 int64
	e = rows1.Scan(&result1)
	if nil != e {
		return 0, farmer.NewFarmerError(e)
	}
	return result1, nil
}
