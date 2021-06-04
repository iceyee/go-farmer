package fsql

//
// import (
//     "database/sql"
//     _ "github.com/go-sql-driver/mysql"
//     "github.com/iceyee/go-farmer/v5/fassert"
//     "testing"
//     //
// )
//
// type t150 struct {
//     Id   int64  `sql_column:"id" sql_key:"1"`
//     Name string `sql_column:"name"`
// }
//
// func (t150) TableName() string {
//     return "fsql_test"
// }
//
// func f150(a []interface{}) []Entity {
//     var result []Entity
//     result = make([]Entity, len(a))
//     for index, value := range a {
//         result[index] = value.(Entity)
//     }
//     return result
// }
//
// func TestRepository(t *testing.T) {
//     var e error
//     dsn := "root:Shi-+123@tcp(db.farmer.ink:3306)/go_farmer"
//     db, e := sql.Open("mysql", dsn)
//     fassert.CheckError(e)
//     r := NewRepository(db, t150{})
//     a001, e := r.FindAll("")
//     _, e = r.DeleteAll(f150(a001)...)
//     fassert.CheckError(e)
//     a002, e := r.Count("")
//     fassert.CheckError(e)
//     fassert.Assert(0 == a002, "删除所有记录")
//     a003 := []Entity{
//         t150{
//             Id:   1,
//             Name: "one",
//         },
//         t150{
//             Id:   2,
//             Name: "two",
//         },
//         t150{
//             Id:   3,
//             Name: "three",
//         },
//         t150{
//             Id:   4,
//             Name: "four",
//         },
//         t150{
//             Id:   5,
//             Name: "five",
//         },
//     }
//     a004, e := r.SaveAll(a003...)
//     fassert.CheckError(e)
//     fassert.Assert(5 == a004, "插入5条记录")
//     a005, e := r.Save(t150{Id: 4, Name: "four, four"})
//     fassert.CheckError(e)
//     fassert.Assert(1 == a005, "修改1条记录")
//     a006, e := r.Exists(t150{Id: 1})
//     fassert.CheckError(e)
//     fassert.Assert(a006, "Id:1 存在")
//     a007, e := r.Exists(t150{Id: 10})
//     fassert.CheckError(e)
//     fassert.Assert(!a007, "Id:10 不存在")
//     return
// }
