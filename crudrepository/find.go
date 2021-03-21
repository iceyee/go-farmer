package crudrepository

import (
	// TODO
	//
	"database/sql"
	"github.com/iceyee/go-farmer/farmer"
)

// 条件查询单个记录, 靠主键来识别
func Find(db *sql.DB, table Table, entity Entity) (interface{}, error) {
	// sql1 - sql语句
	// statement1
	// args - Query()的传参
	// rows1 - Query()之后返回的东西
	var sql1 string = table.Select + " " + table.Where
	//
	println(sql1)
	//
	statement1, e := db.Prepare(sql1)
	if nil != e {
		return nil, farmer.NewFarmerError(e)
	}
	defer statement1.Close()
	var args []interface{} = Arguments2(table, entity)
	rows1, e := statement1.Query(args...)
	if nil != e {
		return nil, farmer.NewFarmerError(e)
	}

	if !rows1.Next() {
		return nil, nil
	}
	result1, args := Arguments3(table, entity)
	e = rows1.Scan(args...)
	if nil != e {
		return nil, farmer.NewFarmerError(e)
	}
	return result1, nil
}

// 条件查询多个记录
func FindAll(db *sql.DB, table Table, entity Entity, where string, args ...interface{}) ([]interface{}, error) {
	// sql1 - sql语句
	// statement1
	// args - Query(),Scan()的传参
	// rows1 - Query()之后返回的东西
	// result1 - 输出
	// result2 - 单个输出, 最后合并到result1
	var sql1 string = table.Select + " " + where
	//
	println(sql1)
	//
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
		result2, args := Arguments3(table, entity)
		e = rows1.Scan(args...)
		if nil != e {
			return nil, farmer.NewFarmerError(e)
		}
		result1 = append(result1, result2)
	}
	return result1, nil
}
