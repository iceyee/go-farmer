package web

import (
	"sort"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	//
)

const apiTemplate string = `
    <table class="table table-dark mb-5">
        <thead></thead>
        <tbody>
            <tr>
                <th scope="row">Description</th>
                <td>$$Description</td>
            </tr>
            <tr>
                <th scope="row">Url</th>
                <td>$$Url</td>
            </tr>
            <tr>
                <th scope="row">Method</th>
                <td>$$Method</td>
            </tr>
            <tr>
                <th scope="row">Parameters</th>
                <td>$$Parameters</td>
            </tr>
            <tr>
                <th scope="row">Response</th>
                <td>$$Response</td>
            </tr>
            <tr>
                <th scope="row">Remarks</th>
                <td>$$Remarks</td>
            </tr>
        </tbody>
    </table>
    <br class="mb-5">
`

const apiTemplate2 string = `
<!doctype html>
<html>
<head>
    <title>Api Document</title>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, user-scalable=no, max-scale=1">
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta2/dist/css/bootstrap.min.css"
      rel="stylesheet" integrity="sha384-BmbxuPwQa2lc/FVzBcNJ7UAyJxM6wuqIj61tLrc4wSX0szH/Ev+nYRRuWlolflfl"
      crossorigin="anonymous">
    <style></style>
</head>
<body class="container py-5">
    $$
    <script src="https://code.jquery.com/jquery-3.6.0.min.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta2/dist/js/bootstrap.bundle.min.js"
          integrity="sha384-b5kHyXgcpbZJO/tY9Ul7kGkf1S0CWuKcCD38l8YkeH8z8QjE0GmW1gYU5S9FOnJ0"
          crossorigin="anonymous"></script>
</body>
</html>
`

type Controller interface {
	GetApi() []ApiDocument
}

type ControllerRegistry struct {
	document string
	router   map[string]ApiDocument
}

// 用于注册控制器
var ControllerRegistryA *ControllerRegistry

func init() {
	ControllerRegistryA = new(ControllerRegistry)
	ControllerRegistryA.router = make(map[string]ApiDocument, 0xff)
	type A struct{}
	ControllerRegistryA.router["/0/api"] = ApiDocument{
		ArgumentType: reflect.TypeOf(*new(A)),
		Key:          "api&api",
		Method:       "GET, POST",
		Url:          "",
		a1:           make(map[string]map[string]interface{}, 0xf),
		document:     "",
		processor:    reflect.ValueOf(api),
	}
	ControllerRegistryA.router["/0/status"] = ApiDocument{
		ArgumentType: reflect.TypeOf(*new(A)),
		Key:          "api&status",
		Method:       "GET, POST",
		Url:          "",
		a1:           make(map[string]map[string]interface{}, 0xf),
		document:     "",
		processor:    reflect.ValueOf(status),
	}

	return
}

// 注册控制器
func (c *ControllerRegistry) Registry(controller Controller) {
	// api - ([]ApiDocument)
	// type1 - Controller (reflect.Type)
	// value1 - Controller (reflect.Value)
	// name1 - Controller's Name (string)
	// method1 - Controller.Method (reflect.Method)
	api := controller.GetApi()
	type1 := reflect.TypeOf(controller)
	value1 := reflect.ValueOf(controller)
	name1 := type1.String()
	for x := 0; x < len(api); x++ {
		if "" == api[x].Url {
			panic(name1 + ", Url不能为空")
		}
		if "" == api[x].Method {
			panic(name1 + ", Method不能为空")
		}
		method1, ok := type1.MethodByName(api[x].MapTo)
		if !ok {
			panic(name1 + "." + api[x].MapTo + "(), 不存在这个方法")
		}
		expectedDefinition1 := "func(" + name1 + ", http.ResponseWriter, *http.Request, interface {})"
		if method1.Type.String() != expectedDefinition1 {
			panic(name1 + "." + method1.Name +
				"(), 方法参数定义错误\n  预期 " + expectedDefinition1 +
				"\n  实际 " + method1.Type.String())
		}
		api[x].processor = value1.MethodByName(api[x].MapTo)

		// field1 - api.ArgumentType.属性(reflect.StructField)
		// tag1 - api.ArgumentType.属性.标签.web (string)
		// a2 - 用于保存解析标签出来的子标签(map[string]interface{})
		// a3 - 切割后的子标签集合([]string)
		// a4 - 单个子标签(string)
		// a5 - 切割子标签得集合, 两个元素([]string)
		// a6 - 临时变量
		api[x].a1 = make(map[string]map[string]interface{}, 0xf)
		for y := 0; y < api[x].ArgumentType.NumField(); y++ {
			field1 := api[x].ArgumentType.Field(y)
			a2 := make(map[string]interface{}, 0xf)
			api[x].a1[field1.Name] = a2
			tag1 := field1.Tag.Get("web")
			if "" == tag1 {
				continue
			}
			a3 := strings.Split(tag1, "###")
			for _, a4 := range a3 {
				a4 = strings.Trim(a4, " ")
				a5 := strings.Split(a4, ":")
				if len(a5) < 2 {
					panic(name1 + "." + field1.Name + ", 错误的标签")
				}
				a2[a5[0]] = strings.Join(a5[1:], ":")
			}

			if _, ok := a2["name"]; !ok {
				panic(name1 + "." + field1.Name + ", 错误的标签")
			}
			if _, _1 := a2["require"]; !_1 {
				if _, _2 := a2["default"]; !_2 {
					panic(name1 + "." + field1.Name + ", 错误的标签, 不能同时没有require和default")
				}
			}
			if _, ok := a2["type"]; !ok {
				a2["type"] = "string"
			} else {
				switch a2["type"].(string) {

				case "string":
					a2["regexp"] = ""

				case "int":
					a2["regexp"] = `^[\+\-]?[0-9]+$`

				case "hex":
					a2["regexp"] = `^[0-9a-fA-F]+$|^0x[0-9a-fA-F]+$`

				case "float":
					a2["regexp"] = `^[\+\-]?[0-9]+$|^[\+\-]?[0-9]*\.[0-9]*$|^[\+\-]?[0-9]*\.[0-9]*[eE][\+\-][0-9]+$`

				case "bool":
					a2["regexp"] = `^true$|^false$`

				case "email":
					a2["regexp"] = `^[0-9a-zA-Z_\-\.]+@[0-9a-zA-Z]+\.[a-zA-Z]+$`

				case "pattern":
					if _, ok := a2["regexp"]; !ok {
						panic(name1 + "." + field1.Name + ", 错误的标签")
					}
					if _, e := regexp.CompilePOSIX(a2["regexp"].(string)); nil != e {
						panic(name1 + "." + field1.Name + ", 错误的标签\n" + e.Error())
					}

				default:
					panic(name1 + "." + field1.Name + ", 错误的标签")
				}
			}

			// 生成文档
			if true {
				api[x].parameters = api[x].parameters + a2["name"].(string) + " - "
				if _, ok := a2["require"]; ok {
					api[x].parameters = api[x].parameters + "必须"
				} else {
					api[x].parameters = api[x].parameters + "可选"
				}
				if a6, ok := a2["desc"]; ok {
					api[x].parameters = api[x].parameters + ", " + a6.(string)
				}
				api[x].parameters = api[x].parameters + ", 类型是" + a2["type"].(string)
				if "pattern" == a2["type"].(string) {
					api[x].parameters = api[x].parameters + ", 正则表达式" + a2["regexp"].(string)
				}
				if a6, ok := a2["max"]; ok {
					api[x].parameters = api[x].parameters + ", 最大" + a6.(string)
				}
				if a6, ok := a2["min"]; ok {
					api[x].parameters = api[x].parameters + ", 最小" + a6.(string)
				}
				if a6, ok := a2["not"]; ok {
					api[x].parameters = api[x].parameters + ", 不能是" + a6.(string)
				}
				if a6, ok := a2["default"]; ok {
					api[x].parameters = api[x].parameters + ", 默认" + a6.(string)
				}
				api[x].parameters = api[x].parameters + "\n"
			}

			if a6, ok := a2["max"]; ok {
				if "hex" == a2["type"].(string) {
					if strings.HasPrefix(a6.(string), "0x") &&
						2 <= len(a6.(string)) {
						a6 = a6.(string)[2:]
					}
					a7, e := strconv.ParseInt(a6.(string), 16, 64)
					if nil != e {
						panic(name1 + "." + field1.Name + ", 错误的标签\n" + e.Error())
					}
					a2["max"] = float64(a7)
				} else {
					a7, e := strconv.ParseFloat(a6.(string), 64)
					if nil != e {
						panic(name1 + "." + field1.Name + ", 错误的标签\n" + e.Error())
					}
					a2["max"] = a7
				}
			}
			if a6, ok := a2["min"]; ok {
				if "hex" == a2["type"].(string) {
					if strings.HasPrefix(a6.(string), "0x") &&
						2 <= len(a6.(string)) {
						a6 = a6.(string)[2:]
					}
					a7, e := strconv.ParseInt(a6.(string), 16, 64)
					if nil != e {
						panic(name1 + "." + field1.Name + ", 错误的标签\n" + e.Error())
					}
					a2["min"] = float64(a7)
				} else {
					a7, e := strconv.ParseFloat(a6.(string), 64)
					if nil != e {
						panic(name1 + "." + field1.Name + ", 错误的标签\n" + e.Error())
					}
					a2["min"] = a7
				}
			}
		}
		api[x].parameters = strings.Replace(api[x].parameters, "&", "&amp;", -1)
		api[x].parameters = strings.Replace(api[x].parameters, "<", "&lt;", -1)
		api[x].parameters = strings.Replace(api[x].parameters, ">", "&gt;", -1)
		api[x].parameters = strings.Replace(api[x].parameters, "\n", "<br>", -1)
		api[x].document = apiTemplate
		api[x].document = strings.Replace(api[x].document, "$$Description", api[x].Description, -1)
		api[x].document = strings.Replace(api[x].document, "$$Method", api[x].Method, -1)
		api[x].document = strings.Replace(api[x].document, "$$Parameters", api[x].parameters, -1)
		api[x].document = strings.Replace(api[x].document, "$$Remarks", api[x].Remarks, -1)
		api[x].document = strings.Replace(api[x].document, "$$Response", api[x].Response, -1)
		api[x].document = strings.Replace(api[x].document, "$$Url", api[x].Url, -1)
		c.router[api[x].Url] = api[x]
	}
	return
}

// 处理路由, 返回true表示继续路由
func (c *ControllerRegistry) process(w http.ResponseWriter, r *http.Request) bool {
	apiDocument, ok := c.router[r.URL.Path]
	if ok {
		// 匹配成功
		if !strings.Contains(apiDocument.Method, r.Method) {
			http.Error(w, "禁止的请求方法", 405)
			return false
		}
		arg, ok, e := c.validate(w, r, apiDocument)
		if nil != e {
			http.Error(w, e.Error(), 400)
			println(e.Error())
		} else if !ok {
		} else {
			apiDocument.processor.Call([]reflect.Value{
				reflect.ValueOf(w),
				reflect.ValueOf(r),
				reflect.ValueOf(arg)})
		}
		return false
	} else {
		// 匹配不成功
		return true
	}
}

func api(w http.ResponseWriter, r *http.Request, arg interface{}) {
	if "" == ControllerRegistryA.document {
		// 没文档
		if len(ControllerRegistryA.router) <= 1 {
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(""))
			return
		} else {
			// 有api
			// 开始生成
			var a1 = make(map[string]string, 4*len(ControllerRegistryA.router))
			var a2 = make([]string, 0, len(ControllerRegistryA.router))
			var a3 = ""
			for key, value := range ControllerRegistryA.router {
				a1[key] = value.document
				a2 = append(a2, key)
			}
			sort.Strings(a2)
			for index, value := range a2 {
				a3 = a3 + a1[value]
				_ = index
			}
			ControllerRegistryA.document = strings.Replace(apiTemplate2, "$$", a3, -1)
		}
	}
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(ControllerRegistryA.document))
	return
}

func status(w http.ResponseWriter, r *http.Request, arg interface{}) {
	w.Write([]byte("ok"))
	return
}
