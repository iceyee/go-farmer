// 便捷的增删查改.
/*
使用步骤:
    1. 定义类. 实现Entity接口. 在属性名的标签加上column表示映射字段, 加key表示这是主键.
    2. Compile(). 调用这个接口之后得到type Table, 这里面包含有数据表映射的信息.
    3. 根据需要调用Count(), Find(), Delete(), Save()等等.
    提示: Entity类不要用指针

Example 

以下为测试代码:
    import (
        "database/sql"
        _ "github.com/go-sql-driver/mysql"
        "github.com/iceyee/go-farmer/farmer"
        "math/rand"
        "testing"
        "time"
        //
    )

    type T1 struct {
        Id   int64  `column:"id" key:"1"`
        Name string `column:"name"`
    }

    func (T1) TableName() string {
        return "table1"
    }

    func Test(t *testing.T) {
        // 连接DB
        var dsn = "root:123456@tcp(localhost:3306)/test"
        db, e := sql.Open("mysql", dsn)
        farmer.CheckError(e)
        farmer.CheckError(db.Ping())
        defer db.Close()

        // 注册table映射
        Debug()
        rand.Seed(time.Now().Unix())
        var table = Compile(T1{})
        // 删除旧记录
        _, e = db.Exec("delete from table1 where id < 5")
        farmer.CheckError(e)
        // Count()
        // a1 - (int64)
        a1, e := Count(db, table, "")
        farmer.CheckError(e)
        farmer.Assert(0 == a1 || 1 == a1)

        // a2 - (type T1)
        // a3 - (type T1)
        // a4 - (type T1)
        // a5 - (type T1)
        var b1 = []string{
            "farmer",
            "iceyee",
            "ubuntu",
            "linux",
            "debian",
            "redhat",
        }
        var a2 = T1{
            Id:   1,
            Name: b1[rand.Int()%len(b1)],
        }
        var a3 = T1{
            Id:   2,
            Name: b1[rand.Int()%len(b1)],
        }
        var a4 = T1{
            Id:   3,
            Name: b1[rand.Int()%len(b1)],
        }
        var a5 = T1{
            Id:   4,
            Name: b1[rand.Int()%len(b1)],
        }
        var a6 = T1{
            Id:   5,
            Name: b1[rand.Int()%len(b1)],
        }

        // Save() SaveAll()
        a1, e = Save(db, table, a2)
        farmer.CheckError(e)
        farmer.Assert(1 == a1)
        a1, e = SaveAll(db, table, []Entity{a3, a4, a5, a6})
        farmer.CheckError(e)
        farmer.Assert(3 == a1 || 4 == a1)

        // Count()
        a1, e = Count(db, table, "")
        farmer.CheckError(e)
        farmer.Assert(5 == a1)

        // Find() FindAll()
        // a7 - (type interface{})
        // a8 - (type []interface{})
        a7, e := Find(db, table, T1{Id: 5})
        farmer.CheckError(e)
        farmer.Assert(a6.Name == a7.(T1).Name)
        a7, e = Find(db, table, T1{Id: 6})
        farmer.CheckError(e)
        farmer.Assert(nil == a7)

        // [0, 3), {1, 2}
        a8, e := FindAll(db, table, "where 0 <= id and id < 3")
        farmer.CheckError(e)
        farmer.Assert(2 == len(a8))
        farmer.Assert(a2.Name == a8[0].(T1).Name)
        farmer.Assert(a3.Name == a8[1].(T1).Name)

        return
    }
*/
package crudrepository
