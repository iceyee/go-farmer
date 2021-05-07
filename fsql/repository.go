package fsql

import (
	"context"
	"database/sql"
	"github.com/iceyee/go-farmer/v3/flog"
	"github.com/iceyee/go-farmer/v3/fstrings"
	"reflect"
	"strings"
	//
)

type KeyMethod int64

const (
	AUTO KeyMethod = iota
	IDENTITY
)

// 字段
type t307 struct {
	IsPrimaryKey bool         // 是否为主键
	KeyGenerate  KeyMethod    // 主键生成方式
	MapTo        string       // 字段名
	Name         string       // 属性名
	Type         reflect.Type // 类型
}

type Repository struct {
	columns    []t307          // 字段
	ctx        context.Context // 上下文
	db         *sql.DB         // 数据库连接
	entityType reflect.Type    //
	sqlDelete  string          // delete from TABLE
	sqlInsert  string          // insert into TABLE set xxx = ?, xxx = ?
	sqlSelect  string          // select xxx, xxx, xxx from TABLE
	sqlUpdate  string          // update TABLE set xxx = ?, xxx = ?
	sqlWhere   string          // where xxx = ? and xxx = ?
	tableName  string          // 表名
}

func NewRepository(db *sql.DB, entity Entity) *Repository {
	if nil == db {
		panic("nil == db")
	}
	if nil == entity {
		panic("nil == entity")
	}
	var r *Repository
	r = new(Repository)
	r.db = db
	r.ctx = context.Background()
	r.tableName = entity.TableName()
	r.columns = make([]t307, 0, 0xff)
	var entityType reflect.Type
	entityType = reflect.TypeOf(entity)
	r.entityType = entityType
	if "" == entityType.Name() {
		panic("entity不能是指针")
	}
	for x := 0; x < entityType.NumField(); x++ {
		var field reflect.StructField
		field = entityType.Field(x)
		if field.Anonymous {
			continue
		}
		var tag reflect.StructTag
		tag = field.Tag
		if "" == tag.Get("sql_column") {
			panic(entityType.String() + "." + field.Name + " 没有定义标签'sql_column'")
		}
		var a001 t307
		a001.MapTo = tag.Get("sql_column")
		a001.Name = field.Name
		a001.Type = field.Type
		if "" == tag.Get("sql_key") {
			a001.IsPrimaryKey = false
		} else {
			a001.IsPrimaryKey = true
			var sqlKey string
			sqlKey = tag.Get("sql_key")
			if "auto" == sqlKey {
				a001.KeyGenerate = AUTO
			} else if "identity" == sqlKey {
				a001.KeyGenerate = IDENTITY
			} else {
				a001.KeyGenerate = IDENTITY
			}
		}
		r.columns = append(r.columns, a001)
	}
	if 0 == len(r.columns) {
		panic(entityType.String() + "定义错误, 请检查是否有定义属性, 属性名是否为大写")
	}
	var existsPrimaryKey bool
	existsPrimaryKey = false
	for _, x := range r.columns {
		if x.IsPrimaryKey {
			existsPrimaryKey = true
			break
		}
	}
	if !existsPrimaryKey {
		panic("缺少主键")
	}
	// 所有字段
	var a002 []string
	a002 = make([]string, 0, 0xff)
	// 主键字段
	var a003 []string
	a003 = make([]string, 0, 0xff)
	// 除了自增的字段
	var a004 []string
	a004 = make([]string, 0, 0xff)
	for _, x := range r.columns {
		if x.IsPrimaryKey {
			a002 = append(a002, x.MapTo)
			a003 = append(a003, x.MapTo)
		} else {
			a002 = append(a002, x.MapTo)
		}
		if !x.IsPrimaryKey ||
			x.KeyGenerate != AUTO {

			a004 = append(a004, x.MapTo)
		}
	}
	r.sqlDelete = "delete from " + r.tableName
	r.sqlInsert = "insert into " + r.tableName + " set " + strings.Join(a004, " = ?, ") + " = ?"
	r.sqlSelect = "select " + strings.Join(a002, ", ") + " from " + r.tableName
	r.sqlUpdate = "update " + r.tableName + " set " + strings.Join(a004, " = ?, ") + " = ?"
	r.sqlWhere = "where " + strings.Join(a003, " = ? and ") + " = ?"
	var sb001 *fstrings.StringBuffer
	sb001 = fstrings.NewStringBuffer()
	sb001.Append("fsql.NewRepository()")
	sb001.Append("\n")
	sb001.Append(entityType.String())
	sb001.Append("\n")
	sb001.Append(r.sqlDelete)
	sb001.Append("\n")
	sb001.Append(r.sqlInsert)
	sb001.Append("\n")
	sb001.Append(r.sqlSelect)
	sb001.Append("\n")
	sb001.Append(r.sqlUpdate)
	sb001.Append("\n")
	sb001.Append(r.sqlWhere)
	flog.Debug(sb001)
	return r
}
