package crud

import (
//
)

// 对应数据库的表, TableName()返回表名,
// 字段的映射关系靠属性的标签来表示, `column:"column_name"`,
// 主键靠属性的标签来表示, `key:"true"`, key的值任意, 不为空就行.
type Entity interface {
	TableName() string
}
