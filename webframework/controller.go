package webframework

import (
	// TODO
	//
	"net/http"
	"reflect"
	"strings"
)

const apiTemplate string = `
    <table class="table mb-5">
        <thead></thead>
        <tbody>
            <tr>
                <th scope="row">Description</th>
                <td>$Description</td>
            </tr>
            <tr>
                <th scope="row">Url</th>
                <td>$Url</td>
            </tr>
            <tr>
                <th scope="row">Method</th>
                <td>$Method</td>
            </tr>
            <tr>
                <th scope="row">Parameters</th>
                <td>$Parameters</td>
            </tr>
            <tr>
                <th scope="row">Response</th>
                <td>$Response</td>
            </tr>
            <tr>
                <th scope="row">Remarks</th>
                <td>$Remarks</td>
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

// 路由的规则是, 路径由方法名转换得到, 转小写, 把'_'换成'-', 即可得到路径,
// 方法的参数必须是(http.ResponseWriter, *http.Request),
// 需要继承这个类
type Controller struct{}

type ControllerRegistry struct {
	api         map[string]string
	controllers []interface{}
	document    string
	router      map[string]reflect.Value
}

// 用于注册控制器
var ControllerRegistryA *ControllerRegistry

func init() {
	ControllerRegistryA = new(ControllerRegistry)
	ControllerRegistryA.api = make(map[string]string, 0xff)
	ControllerRegistryA.controllers = make([]interface{}, 0, 0xf)
	ControllerRegistryA.router = make(map[string]reflect.Value, 0xff)
	ControllerRegistryA.router["/api"] = reflect.ValueOf(api)
	ControllerRegistryA.router["/status"] = reflect.ValueOf(status)
	return
}

func api(w http.ResponseWriter, r *http.Request) {
	if "" == ControllerRegistryA.document &&
		0 < len(ControllerRegistryA.api) {
		ControllerRegistryA.GenerateApiDocument()
	}
	w.Header().Set("Content-Type", "text/html")
	_, e := w.Write([]byte(ControllerRegistryA.document))
	if nil != e {
		panic(e)
	}
	return
}

func status(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
	return
}

// 注册控制器
func (c *ControllerRegistry) Registry(controller interface{}) {
	c.controllers = append(c.controllers, controller)
	// 路由
	for _, controller := range c.controllers {
		type1 := reflect.TypeOf(controller)
		value1 := reflect.ValueOf(controller)
		for x := 0; x < type1.NumMethod(); x++ {
			method1 := type1.Method(x)
			methodDefinition1 := method1.Type.String()
			if !strings.Contains(methodDefinition1, ", http.ResponseWriter, *http.Request)") {
				continue
			}
			road1 := strings.ToLower(method1.Name)
			road1 = strings.ReplaceAll(road1, "_", "-")
			c.router["/"+road1] = value1.Method(x)
		}
	}
	// 生成文档
	for _, controller := range c.controllers {
		type1 := reflect.TypeOf(controller)
		value1 := reflect.ValueOf(controller)
		for x := 0; x < type1.NumMethod(); x++ {
			method1 := type1.Method(x)
			methodDefinition1 := method1.Type.String()
			if !strings.Contains(methodDefinition1, "ApiDocument") {
				continue
			}
			apiDocument1 := value1.Method(x).Call([]reflect.Value{})[0].Interface().(ApiDocument)
			var apiDocument2 string = apiTemplate
			apiDocument2 = strings.Replace(apiDocument2, "$Description", apiDocument1.Description, -1)
			apiDocument2 = strings.Replace(apiDocument2, "$Method", apiDocument1.Method, -1)
			apiDocument2 = strings.Replace(apiDocument2, "$Parameters", apiDocument1.Parameters, -1)
			apiDocument2 = strings.Replace(apiDocument2, "$Remarks", apiDocument1.Remarks, -1)
			apiDocument2 = strings.Replace(apiDocument2, "$Response", apiDocument1.Response, -1)
			apiDocument2 = strings.Replace(apiDocument2, "$Url", apiDocument1.Url, -1)
			c.api[apiDocument1.Key] = apiDocument2
		}
	}
	return
}

func (c *ControllerRegistry) GenerateApiDocument() {
	var apiDocument3 string
	for _, value := range c.api {
		apiDocument3 = apiDocument3 + value
	}
	c.document = strings.Replace(apiTemplate2, "$$", apiDocument3, -1)
	return
}

// 处理路由, 返回true表示继续路由
func (c *ControllerRegistry) Process(w http.ResponseWriter, r *http.Request) bool {
	method1, ok := c.router[r.URL.Path]
	if ok {
		// 匹配成功
		method1.Call([]reflect.Value{reflect.ValueOf(w), reflect.ValueOf(r)})
		return false
	} else {
		// 匹配不成功
		return true
	}
}
