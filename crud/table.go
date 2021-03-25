package crud

import (
	"reflect"
	"strings"
	//
)

var verbose = false

// 开启调试模式, 在这个模式下, 会输出一些额外消息.
func Debug() {
	verbose = true
	return
}

// 将有映射关系的类编译, 即可得到这个表.
type Table struct {
	Name       string                         // 表名
	Delete     string                         // delete from TABLE
	Insert     string                         // insert into TABLE set xxx = ?, xxx = ?
	Select     string                         // select xxx, xxx, xxx from TABLE
	Update     string                         // update TABLE set xxx = ?, xxx = ?
	Where      string                         // where xxx = ? and xxx = ?
	columns    []string                       // 字段集
	entityType reflect.Type                   // Entity的类型Type
	keys       []string                       // 主键集
	mapTo      map[string]reflect.StructField // 字段和属性名的映射, key是字段, value是属性类型
}

// 由Entity得到Table
func Compile(entity Entity) Table {
	// 定义Table, 初始化属性
	var table Table
	table.Name = entity.TableName()
	table.columns = make([]string, 0, 0xff)
	table.entityType = reflect.TypeOf(entity)
	table.keys = make([]string, 0, 0xff)
	table.mapTo = make(map[string]reflect.StructField, 0xff)

	// field1 - entity的属性, 各种类型(type reflect.StructField)
	// column1 - 字段名(type string)
	for x := 0; x < table.entityType.NumField(); x++ {
		var field1 = table.entityType.Field(x)
		var column1 string = field1.Tag.Get("column")
		if "" == column1 {
			panic("数据表映射必须要在标签中写column")
		}
		table.columns = append(table.columns, column1)
		table.mapTo[column1] = field1

		if "" != field1.Tag.Get("key") {
			table.keys = append(table.keys, column1)
		}
	}
	if 0 == len(table.keys) {
		panic("糟糕的定义, 缺少主键")
	}

	// 合成Delete, Insert, Select, Update
	// index - 索引(type int)
	// value - 字段(type string)
	// a1 - xxx = ? (type []string)
	// a2 - xxx (type []string)
	// a3 - (type string)
	var a1 []string = make([]string, 0, 0xff)
	var a2 []string = make([]string, 0, 0xff)
	for index, value := range table.columns {
		a1 = append(a1, value+" = ?")
		a2 = append(a2, value)
		_ = index
	}
	table.Delete = "delete from " + table.Name
	table.Insert = "insert into " + table.Name + " set " + strings.Join(a1, ", ")
	table.Select = "select " + strings.Join(a2, ", ") + " from " + table.Name
	table.Update = "update " + table.Name + " set " + strings.Join(a1, ", ")

	// 合成Where
	// index - 索引(type int)
	// value - 字段(type string)
	// value - 属性名(type string)
	// a4 - xxx = ? (type []string)
	var a4 []string = make([]string, 0, 0xff)
	for index, value := range table.keys {
		a4 = append(a4, value+" = ?")
		_ = index
	}
	table.Where = " where " + strings.Join(a4, " and ")

	return table
}

// 根据entity的属性生成, 可用于DB.Prepare().Exec(...),
// 这个是生成全部字段的, 和Arguments2()有区别
func arguments1(table Table, entity Entity) []interface{} {
	// index - 索引(type int)
	// value - 字段(type string)
	// result - 输出结果(type []interface{})
	// value1 - entity的Value (type reflect.Value)
	var result = make([]interface{}, 0, len(table.columns))
	var value1 = reflect.ValueOf(entity)
	for index, value := range table.columns {
		result = append(result, value1.FieldByName(table.mapTo[value].Name).Interface())
		_ = index
	}
	return result
}

// 根据entity的属性生成, 可用于DB.Prepare().Exec(...),
// 这个是生成主键字段的, 和Arguments1()有区别
func arguments2(table Table, entity Entity) []interface{} {
	// index - 索引(type int)
	// value - 字段(type string)
	// result - 输出结果(type []interface{})
	// value1 - entity的Value (type reflect.Value)
	// counter - 计数, 索引(type int)
	var result = make([]interface{}, 0, len(table.keys))
	var value1 = reflect.ValueOf(entity)
	for index, value := range table.keys {
		result = append(result, value1.FieldByName(table.mapTo[value].Name).Interface())
		_ = index
	}
	return result
}

// 生成两个参数, 返回第一个是类, 第二个可用于sql.Rows.Scan(...)
func arguments3(table Table) (interface{}, []interface{}) {
	// index - 索引(type int)
	// value - 字段(type string)
	// data - 输出结果1 (type Value(interface{} - *Entity))
	// result - 输出结果2 (type []interface{} - 指针集)
	var data = reflect.New(table.entityType)
	var result = make([]interface{}, 0, len(table.columns))
	for index, value := range table.columns {
		result = append(result, data.Elem().FieldByName(table.mapTo[value].Name).Addr().Interface())
		_ = index
	}
	return data.Interface(), result
}
