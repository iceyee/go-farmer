// web框架: 拦截器, 控制器, 文件服务器, ApiDocument.
//
/*
使用这个框架的几个步骤:
  1. 授权域名访问. AuthorizeDomain()
  2. 设置拦截器. InterceptorRegistryA.Registry()
  3. 设置控制器. ControllerRegistryA.Registry()
  4. 设置文件服务器. FileServerA.Registry()
  5. 启动. Listen()

有一些已经设置好的行为, /0/api是查看文档的, /0/status是查看服务器状态.

定义Url参数的标签, 用###间隔, 支持的关键字如下:
  default - 默认值
  desc - 对这个参数的说明, 在生成api文档时用到
  max - 最大值
  min - 最小值
  name - url参数的名字, 必须
  not - 不等于
  regexp - 正则表达式, 当type是pattern时才有效
  require - 是否必须
  type - 类型, 包括string, int, float, hex, bool, email, pattern, 默认string

  default和require必须二选一, 同时存在则default无效

下面举几个例子.

Example Of Interceptor

定义一个拦截器, 功能是将HEAD请求改为GET请求, 代码如下:
    type farmerInterceptor struct {
    }

    func (i *farmerInterceptor) Process(w http.ResponseWriter, r *http.Request) bool {
        if "HEAD" == r.Method {
            r.Method = "GET"
        }
        return true
    }

Example Of Controller

接下来定义一个Controller, 没有什么功能, 主要用来说明怎么定义Controller.
    type farmerController struct {
    }

    func (c *farmerController) GetApi() []ApiDocument {
        return []ApiDocument{
            c.a1(),
        }
    }

    type t4715 struct {
        A string `web:"name:A ###type:string ###require: ###desc:描述1 ###not:hello "`
        B int64  `web:"name:b ###type:hex ###default:0xf ###desc:这个要求是十六进制数 ###max:0xff ###min:0x2 ###not:50"`
        C bool   `web:"name:c ###type:bool ###default:false "`
    }

    func (c *farmerController) a1() ApiDocument {
        return ApiDocument{
            ArgumentType: reflect.TypeOf(t4715{}),
            Description:  "这是第一个测试接口",
            Key:          "test1",
            MapTo:        "A",
            Method:       "GET,POST",
            Remarks:      "测试1",
            Response:     "如果参数验证没问题, 就返回ok, 否则返回错误的参数, 以及状态码400",
            Url:          "/0/test1",
        }
    }

    func (c *farmerController) A(w http.ResponseWriter, r *http.Request, arg interface{}) {
        a1 := arg.(t4715)
        print("A=")
        println(a1.A)
        print("b=")
        println(a1.B)
        print("c=")
        println(a1.C)
        println()
        println()
        w.Write([]byte("ok"))
        return
    }

Example Of FileServer

然后是文件服务器, 这个不需要定义, 只需要指定好url映射的目录即可.

比如, FileServerA.Registry("/t/", "/tmp/"), 那么/t/1会映射到/tmp/1, /t/1/2会映射到/tmp/1/2, 以此类推.

Starting

最后一步, 授权域名访问, 注册拦截器, 控制器, 文件服务器, 并启动, 示例如下:
    AuthorizeDomain("farmer.ink")
    InterceptorRegistryA.Registry(new(farmerInterceptor))
    ControllerRegistryA.Registry(new(farmerController))
    FileServerA.Registry("/", "/tmp/")
    Listen(":8888")

这时, 访问/localhost:8888/0/status可以看到服务器状态, 访问/localhost:8888/0/api可以看到服务器所有的api
*/
package web
