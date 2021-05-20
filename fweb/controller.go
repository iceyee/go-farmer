package fweb

import (
	"github.com/iceyee/go-farmer/v4/flog"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	//
)

var x251 map[string]t184 = make(map[string]t184, 0xfff)

func RegistryController(controller Controller) {
	var type001 reflect.Type
	type001 = reflect.TypeOf(controller)
	var value001 reflect.Value
	value001 = reflect.ValueOf(controller)
	for x := 0; x < type001.NumMethod(); x++ {
		var method reflect.Method
		method = type001.Method(x)
		var a001 string
		a001 = strings.Replace(method.Type.String(), " ", "", -1)
		if ok, _ := regexp.MatchString(`^func\([^\(\)]+\)string$`, a001); !ok {
			continue
		}
		var a002 string
		a002 = method.Func.Call([]reflect.Value{value001})[0].String()
		var b001 t184
		b001.Executor = value001
		b001.Parameters = make([]t377, 0, 0xff)
		var a003 []string
		a003 = strings.Split(a002, "\n")
		var a004 map[string]t377
		a004 = make(map[string]t377, 0xff)
		for _, y := range a003 {
			y = strings.TrimLeft(y, " \t\r")
			if "" == y ||
				!strings.HasPrefix(y, "@") {

				continue
			}
			if strings.HasPrefix(y, "@Constraints ") {
				var a005 []string = strings.Split(y, " | ")
				if len(a005) < 6 {
					panic("定义错误.\n" + type001.String() + "\n" + method.Name + "\n" + y)
				}
				var a006 t377
				var a007 string
				a007 = strings.Trim(a005[1], " \t\r\n")
				if "" == a007 {
					panic("定义错误.\n" + type001.String() + "\n" + method.Name + "\n" + y)
				}
				a006.Name = a007
				a007 = strings.Trim(a005[2], " \t\r\n")
				if "" != a007 {
					a008, e := strconv.ParseFloat(a007, 64)
					if nil != e {
						panic("定义错误.\n" + type001.String() + "\n" + method.Name + "\n" + y)
					}
					a006.Min = new(float64)
					*a006.Min = a008
				}
				a007 = strings.Trim(a005[3], " \t\r\n")
				if "" != a007 {
					a008, e := strconv.ParseFloat(a007, 64)
					if nil != e {
						panic("定义错误.\n" + type001.String() + "\n" + method.Name + "\n" + y)
					}
					a006.Max = new(float64)
					*a006.Max = a008
				}
				a007 = strings.Trim(a005[4], " \t\r\n")
				a006.Not = a007
				a007 = strings.Trim(strings.Join(a005[5:], " | "), " \t\r\n")
				a006.Regexp = a007
				a004[a006.Name] = a006
			} else if strings.HasPrefix(y, "@Description ") {
				var a005 []string = strings.Split(y, " ")
				if len(a005) < 2 {
					panic("定义错误.\n" + type001.String() + "\n" + method.Name + "\n" + y)
				}
				b001.Description =
					strings.Trim(strings.Join(a005[1:], " "), " \t\r\n")
				b001.Description = strings.Replace(b001.Description, "&", "&amp;", -1)
				b001.Description = strings.Replace(b001.Description, "<", "&lt;", -1)
				b001.Description = strings.Replace(b001.Description, ">", "&gt;", -1)
			} else if strings.HasPrefix(y, "@MapTo ") {
				var a005 []string = strings.Split(y, " ")
				if len(a005) < 2 {
					panic("定义错误.\n" + type001.String() + "\n" + method.Name + "\n" + y)
				}
				var a006 string = strings.Trim(strings.Join(a005[1:], " "), " \t\r\n")
				if m, ok := type001.MethodByName(a006); !ok {
					panic("定义错误.\n" + type001.String() + "\n" + method.Name + "\n" + y)
				} else {
					b001.MapTo = m.Func
					a001 = strings.Replace(m.Type.String(), " ", "", -1)
				}
			} else if strings.HasPrefix(y, "@Method ") {
				var a005 []string = strings.Split(y, " ")
				if len(a005) < 2 {
					panic("定义错误.\n" + type001.String() + "\n" + method.Name + "\n" + y)
				}
				b001.Method =
					strings.Trim(strings.Join(a005[1:], " "), " \t\r\n")
			} else if strings.HasPrefix(y, "@Parameter ") {
				var a005 []string = strings.Split(y, " | ")
				if len(a005) < 6 {
					panic("定义错误.\n" + type001.String() + "\n" + method.Name + "\n" + y)
				}
				var a006 t377
				var a007 string
				a007 = strings.Trim(a005[1], " \t\r\n")
				if "" == a007 {
					panic("定义错误.\n" + type001.String() + "\n" + method.Name + "\n" + y)
				}
				a006.Name = a007
				a007 = strings.Trim(a005[2], " \t\r\n")
				if "" == a007 {
					panic("定义错误.\n" + type001.String() + "\n" + method.Name + "\n" + y)
				}
				a006.Type = a007
				a007 = strings.Trim(a005[3], " \t\r\n")
				if "" == a007 {
					a006.Required = false
				} else {
					a006.Required = true
				}
				a007 = strings.Trim(a005[4], " \t\r\n")
				a006.Default = a007
				a007 = strings.Trim(strings.Join(a005[5:], " | "), " \t\r\n")
				a006.Description = a007
				a006.Description = strings.Replace(a006.Description, "&", "&amp;", -1)
				a006.Description = strings.Replace(a006.Description, "<", "&lt;", -1)
				a006.Description = strings.Replace(a006.Description, ">", "&gt;", -1)
				b001.Parameters = append(b001.Parameters, a006)
			} else if strings.HasPrefix(y, "@Remarks ") {
				var a005 []string = strings.Split(y, " ")
				if len(a005) < 2 {
					panic("定义错误.\n" + type001.String() + "\n" + method.Name + "\n" + y)
				}
				b001.Remarks =
					strings.Trim(strings.Join(a005[1:], " "), " \t\r\n")
				b001.Remarks = strings.Replace(b001.Remarks, "&", "&amp;", -1)
				b001.Remarks = strings.Replace(b001.Remarks, "<", "&lt;", -1)
				b001.Remarks = strings.Replace(b001.Remarks, ">", "&gt;", -1)
			} else if strings.HasPrefix(y, "@Response ") {
				var a005 []string = strings.Split(y, " ")
				if len(a005) < 2 {
					panic("定义错误.\n" + type001.String() + "\n" + method.Name + "\n" + y)
				}
				b001.Response =
					strings.Trim(strings.Join(a005[1:], " "), " \t\r\n")
				b001.Response = strings.Replace(b001.Response, "&", "&amp;", -1)
				b001.Response = strings.Replace(b001.Response, "<", "&lt;", -1)
				b001.Response = strings.Replace(b001.Response, ">", "&gt;", -1)
			} else if strings.HasPrefix(y, "@Url ") {
				var a005 []string = strings.Split(y, " ")
				if len(a005) < 2 {
					panic("定义错误.\n" + type001.String() + "\n" + method.Name + "\n" + y)
				}
				b001.Url =
					strings.Trim(strings.Join(a005[1:], " "), " \t\r\n")
			} else {
				continue
			}
		}
		for index, value := range b001.Parameters {
			if p, ok := a004[value.Name]; ok {
				value.Min = p.Min
				value.Max = p.Max
				value.Not = p.Not
				value.Regexp = p.Regexp
				b001.Parameters[index] = value
			}
		}
		if "" == b001.Method {
			b001.Method = "GET,POST"
		}
		var a009 string = `^func\([^,]+,\*fweb.Session,http.ResponseWriter,\*http.Request`
		for _, value := range b001.Parameters {
			if "string" != value.Type &&
				"" == value.Default &&
				!value.Required {

				println("非string类型, 必须require和default二选一.")
				panic("定义错误.\n" + type001.String() + "\n" + method.Name)
			}
			if "float64" == value.Type {
				a009 += ",float64"
			} else if "int64" == value.Type {
				a009 += ",int64"
			} else if "string" == value.Type {
				a009 += ",string"
				if nil != value.Max || nil != value.Min {
					println("string类型的参数不能有min或max.")
					panic("定义错误.\n" + type001.String() + "\n" + method.Name)
				}
			} else {
				println("参数类型只能是string, int64, float64.")
				panic("定义错误.\n" + type001.String() + "\n" + method.Name)
			}
		}
		a009 += `\)$`
		if ok, _ := regexp.MatchString(a009, a001); !ok {
			println("接口不对.")
			println(a009)
			println(a001)
			panic("定义错误.\n" + type001.String() + "\n" + method.Name)
		}
		if "" == b001.Url {
			println("未定义Url.")
			panic("定义错误.\n" + type001.String() + "\n" + method.Name)
		}
		b001.SortKey = type001.String() + method.Name
		x251[b001.Url] = b001
		flog.Debug("Controller, " + b001.Url)
	}
	return
}

func processController(
	session *Session,
	w http.ResponseWriter,
	r *http.Request,
	controller t184) {

	if !strings.Contains(controller.Method, r.Method) {
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	var b001 []reflect.Value
	b001 = make([]reflect.Value, 0, 0xff)
	b001 = append(b001, controller.Executor)
	b001 = append(b001, reflect.ValueOf(session))
	b001 = append(b001, reflect.ValueOf(w))
	b001 = append(b001, reflect.ValueOf(r))
	for _, x := range controller.Parameters {
		var a001 string
		a001 = r.FormValue(x.Name)
		a001, _ = url.QueryUnescape(a001)
		if x.Required && "" == a001 {
			http.Error(w, "Bad Request", 400)
			return
		}
		if "" == a001 {
			a001 = x.Default
		}
		if "" != x.Not && x.Not == a001 {
			http.Error(w, "Bad Request", 400)
			return
		}
		if "" != x.Regexp &&
			(x.Required || "" != a001) {
			if ok, e := regexp.MatchString(x.Regexp, a001); nil == e && !ok {
				http.Error(w, "Bad Request", 400)
				return
			}
		}
		if "string" == x.Type {
			b001 = append(b001, reflect.ValueOf(a001))
		} else if "float64" == x.Type {
			var a002 float64
			a002, e := strconv.ParseFloat(a001, 64)
			if nil != e {
				http.Error(w, "Bad Request", 400)
				return
			}
			if nil != x.Min && a002 < *x.Min {
				http.Error(w, "Bad Request", 400)
				return
			}
			if nil != x.Max && *x.Max < a002 {
				http.Error(w, "Bad Request", 400)
				return
			}
			b001 = append(b001, reflect.ValueOf(a002))
		} else if "int64" == x.Type {
			var a002 int64
			a002, e := strconv.ParseInt(a001, 10, 64)
			if nil != e {
				http.Error(w, "Bad Request", 400)
				return
			}
			if nil != x.Min && float64(a002) < *x.Min {
				http.Error(w, "Bad Request", 400)
				return
			}
			if nil != x.Max && *x.Max < float64(a002) {
				http.Error(w, "Bad Request", 400)
				return
			}
			b001 = append(b001, reflect.ValueOf(a002))
		}
	}
	controller.MapTo.Call(b001)
	return
}
