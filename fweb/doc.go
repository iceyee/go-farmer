// web框架.
//
// 接口说明支持以下标签.
/*
   @Url

   @MapTo

   @Method [GET,POST等等]

   @Description

   @Response

   @Remarks

   @Parameter | [name] | [type] | [required] | [default] | [description]
   @Constraints | [name] | [min] | [max] | [not] | [regexp]
*/
//
// @Url和@MapTo是必须的, 不然会报错.
//
// @Parameter和@Constraints是用' | '分隔.
//
// [type]: string, int64, float64.
//
// [regexp]: 应该尽量避免出现'||', 因为这个会被特殊处理.
//
// 如果url参数为空, 则对应的入参指针也是nil.
//
// 示例代码
//
/*
   type A struct {
   }

   func (*A) Filter(path string) bool {
       return true
   }

   func (*A) Process(
       session *Session,
       w http.ResponseWriter,
       r *http.Request) bool {

       if "DELETE" == r.Method {
           r.Method = "GET"
       }
       return true
   }

   type B struct {
   }

   func (*B) Test__() string {
       return `
          @Url /test

          @MapTo Test

          @Method GET

          @Description 这是描述

          @Response 响应说明

          @Remarks 备注

          @Parameter | A | int64 | 1 |   | 这是A
          @Constraints | A | 0 | 100 | 50 |

          @Parameter | B | float64 |   | 1 | 这是B
          @Constraints | B | 0 | 100 | 50 |

          @Parameter | C | string | 1 |   | 这是C
          @Constraints | C |   |   | 50 | hello

          @Parameter | D | int64 |   |   | 这是D
          @Constraints | D |   |   |   |
          `
   }

   type t struct {
       A int64
       B float64
       C string
       D int64
   }

   func (*B) Test(
       session *Session,
       w http.ResponseWriter,
       r *http.Request,
       A *int64,
       B *float64,
       C *string,
       D *int64) {

       a := t{
           A: *A,
           B: *B,
           C: *C,
       }
       if nil == D {
           a.D = -1
       } else {
           a.D = *D
       }
       WriteJson(w, a)
       return
   }

   func Test(t *testing.T) {
       RegistryInterceptor(new(A))
       RegistryController(new(B))
       RegistryFileServer("/", "/tmp/")
       panic(Listen(":9999"))
   }
*/
package fweb

import (
//
)
