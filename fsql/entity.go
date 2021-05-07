package fsql

import (
//
)

// 对应数据库的表, TableName()返回表名,
// 字段的映射关系靠属性的标签来表示, `sql_column:"column_name"`,
// 主键靠属性的标签来表示, `sql_key:"auto"`, sql_key的值可以是"auto"或"identity",
// "auto"表示自增主键, 不为空就行.
type Entity interface {
	TableName() string
}
