package crudrepository

import (
	// TODO
	//
	"database/sql"
	"github.com/iceyee/go-farmer/farmer"
)

// 保存或更新一个记录
func Save(db *sql.DB, table Table, entity Entity) (int64, error) {
	// 先查是否存在
	// sql1 - sql语句
	// statement1
	// rows1
	var sql1 string = table.Select + " " + table.Where
	statement1, e := db.Prepare(sql1)
	if nil != e {
		return 0, farmer.NewFarmerError(e)
	}
	defer statement1.Close()
	var args []interface{} = Arguments2(table, entity)
	rows1, e := statement1.Query(args...)
	if nil != e {
		return 0, farmer.NewFarmerError(e)
	}

	if rows1.Next() {
		// 已存在
		// sql2 - sql语句
		// statement2 - Exec()
		// args2 - update部分的参数
		// args3 - where部分的参数
		// result2 - 输出
		sql2 := table.Update + " " + table.Where
		//
		println(sql2)
		//
		statement2, e := db.Prepare(sql2)
		if nil != e {
			return 0, farmer.NewFarmerError(e)
		}
		args2 := Arguments1(table, entity)
		args3 := Arguments2(table, entity)
		for _, value := range args3 {
			args2 = append(args2, value)
		}
		result2, e := statement2.Exec(args2...)
		if nil != e {
			return 0, farmer.NewFarmerError(e)
		}
		return result2.RowsAffected()
	} else {
		// 不存在
		// sql2 - sql语句
		// statement2 - Exec()
		// args2 -
		// result2 - 输出
		sql2 := table.Insert
		//
		println(sql2)
		//
		statement2, e := db.Prepare(sql2)
		if nil != e {
			return 0, farmer.NewFarmerError(e)
		}
		args2 := Arguments1(table, entity)
		result2, e := statement2.Exec(args2...)
		if nil != e {
			return 0, farmer.NewFarmerError(e)
		}
		return result2.RowsAffected()
	}
}

// 保存多条记录
func SaveAll(db *sql.DB, table Table, entityA []Entity) (int64, error) {
	// affected - 输出
	// entity
	// a1 - 临时变量
	var affected int64
	for _, entity := range entityA {
		a1, e := Save(db, table, entity)
		if nil != e {
			return affected, e
		}
		affected += a1
	}
	return affected, nil
}
