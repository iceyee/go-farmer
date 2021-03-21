package crudrepository

import (
	// TODO
	//
	"reflect"
	"strings"
)

// 将有映射关系的类编译, 即可得到这个表.
type Table struct {
	Name    string                  // 表名
	columns map[string]int          // 字段到Entity类的映射
	keys    map[string]int          // 主键到Entity类的映射
	types   map[string]reflect.Type // Entity的属性Type合集, 与columns对应
	Delete  string                  // delete from TABLE
	Insert  string                  // insert into TABLE set xxx = ?, xxx = ?
	Select  string                  // select xxx, xxx, xxx from TABLE
	Update  string                  // update TABLE set xxx = ?, xxx = ?
	Where   string                  // where xxx = ? and xxx = ?
}

// 由Entity得到Table
func Compile(entity Entity) Table {
	// 定义Table, 初始化属性
	var table Table
	table.Name = entity.TableName()
	table.columns = make(map[string]int, 0xff)
	table.keys = make(map[string]int, 0xff)
	table.types = make(map[string]reflect.Type, 0xff)

	// entityType1 - entity的Type, struct
	// field1 - entity的属性, 各种类型
	// column1 - 对应的数据库字段名
	var entityType1 reflect.Type = reflect.TypeOf(entity)
	for x := 0; x < entityType1.NumField(); x++ {
		var field1 reflect.StructField = entityType1.Field(x)
		var column1 string = field1.Tag.Get("column")
		if "" == column1 {
			panic("数据表映射必须要在标签中写column")
		}
		table.columns[column1] = x

		if "" != field1.Tag.Get("key") {
			table.keys[column1] = x
		}

		table.types[column1] = field1.Type
	}
	if 0 == len(table.keys) {
		panic("糟糕的定义")
	}

	// 合成Delete, Insert, Select, Update
	// key - 字段
	// value - 对应在类属性中的索引
	// a1 - xxx = ?
	// a2 - xxx
	// a3 - 临时变量
	var a1 []string = make([]string, 0, 0xf)
	var a2 []string = make([]string, 0, 0xf)
	for key, value := range table.columns {
		a1 = append(a1, key+" = ?")
		a2 = append(a2, key)
		value = value
	}
	table.Delete = "delete from " + table.Name
	table.Insert = "insert into " + table.Name + " set " + strings.Join(a1, ", ")
	table.Select = "select " + strings.Join(a2, ", ") + " from " + table.Name
	table.Update = "update " + table.Name + " set " + strings.Join(a1, ", ")

	// 合成Where
	// key - 字段, 也是主键
	// value - 对应在类属性中的索引
	// a4 - xxx = ?
	var a4 []string = make([]string, 0, 0xf)
	for key, value := range table.keys {
		a4 = append(a4, key+" = ?")
		value = value
	}
	table.Where = " where " + strings.Join(a4, " and ")

	return table
}

// 根据entity的属性生成, 可用于DB.Prepare().Exec(...),
// 这个是生成全部字段的, 和Arguments2()有区别
func Arguments1(table Table, entity Entity) []interface{} {
	// key - 字段
	// value - 对应在类属性中的索引
	// result - 输出结果
	// value1 - Entity类型的Value
	// value2 - 指向field的指针
	var result []interface{} = make([]interface{}, 0, len(table.columns))
	var value1 reflect.Value = reflect.ValueOf(entity)
	for key, value := range table.columns {
		var value2 reflect.Value = reflect.New(table.types[key])
		value2.Elem().Set(value1.Field(value))
		result = append(result, value2.Pointer())
	}
	return result
}

// 根据entity的属性生成, 可用于DB.Prepare().Exec(...),
// 这个是生成主键字段的, 和Arguments1()有区别
func Arguments2(table Table, entity Entity) []interface{} {
	// key - 主键字段
	// value - 对应在类属性中的索引
	// result - 输出结果
	// value1 - Entity类型的Value
	// value2 - 指向field的指针
	var result []interface{} = make([]interface{}, 0, len(table.columns))
	var value1 reflect.Value = reflect.ValueOf(entity)
	for key, value := range table.keys {
		var value2 reflect.Value = reflect.New(table.types[key])
		value2.Elem().Set(value1.Field(value))
		result = append(result, value2.Pointer())
	}
	return result
}

// 生成两个参数, 返回第一个是类, 第二个可用于sql.Rows.Scan(...)
func Arguments3(table Table, entity Entity) (interface{}, []interface{}) {
	// key - 字段
	// value - 对应在类属性中的索引
	// data - 输出结果1
	// result - 输出结果2
	// value2 - 指向field的指针
	var data reflect.Value = reflect.New(reflect.TypeOf(entity)).Elem()
	var result []interface{} = make([]interface{}, 0, len(table.columns))
	for key, value := range table.columns {
		result = append(result, data.Field(value).Addr().Pointer())
		key = key
	}
	return data, result
}
