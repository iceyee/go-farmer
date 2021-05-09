// web框架.
//
// 接口说明支持以下标签.
/*
   @Url

   @MapTo

   @SortKey

   @Method [GET,POST等等]

   @Description

   @Response

   @Remarks

   @Parameter | [name] | [type] | [required] | [default] | [description]
   @Constraints | [name] | [min] | [max] | [not] | [regexp]
*/
// @Parameter和@Constraints是用' | '分隔, 如果要某个留空, 那么就应该'|  |',
// 中间要留够位置, 不然会出现非预期的情况.
//
// 对于非string类型, [required]和[default]必须二选一.
//
// @Url和@MapTo是必须的, 不然会报错.
//
// 示例代码
//
/*
   type A struct {
   }

   func (*A) Filter(path string) bool {
       return true
   }

   func (*A) Process(session *Session, w http.ResponseWriter, r *http.Request) bool {
       if "DELETE" == r.Method {
           r.Method = "GET"
       }
       return true
   }

   type B struct {
   }

   func (*B) F123() string {
       return `
       @Description 这是描述

       @MapTo Test

       @Method GET

       @Remarks 备注

       @Response 响应说明

       @Url /test

       @Parameter | A | int64 | 1 |  | 这是A
       @Constraints | A | 0 | 100 | 50 |

       @Parameter | B | float64 |  | 1 | 这是B
       @Constraints | B | 0 | 100 | 50 |

       @Parameter | C | string | 1 |  | 这是C
       @Constraints | C |  |  | 50 | hello
       `
   }

   type t struct {
       A int64
       B float64
       C string
   }

   func (*B) Test(
       session *fweb.Session,
       w http.ResponseWriter,
       r *http.Request,
       A int64,
       B float64,
       C string) {

       a := t{
           A: A,
           B: B,
           C: C,
       }
       WriteJson(w, a)
       return
   }

   func Test(t *testing.T) {
       RegistryInterceptor(new(A))
       RegistryController(new(B))
       RegistryFileServer("/", "/tmp/")
       Listen(":9999")
   }
*/
package fweb

import (
//
)
