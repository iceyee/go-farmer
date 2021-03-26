package web

import (
	"github.com/iceyee/go-farmer/farmer"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	//
)

// 验证表单并生成参数. 因为要配合ControllerRegistry用, 所以是不公开的.
func (c *ControllerRegistry) validate(w http.ResponseWriter, r *http.Request,
	api ApiDocument) (interface{}, bool, error) {
	// result1 - 输出结果, struct (reflect.Value)
	// value1 - url参数的值(string)
	// key - 属性名(string)
	// value - 属性的标签(map[string]interface{})
	// a2 - 临时变量
	// a3 - 临时变量
	// a4 - 临时变量
	result1 := reflect.New(api.ArgumentType).Elem()
	for key, value := range api.a1 {
		var value1 = r.FormValue(value["name"].(string))
		if "" == value1 {
			if _, ok := value["require"]; ok {
				http.Error(w, "错误的参数, require, "+value["name"].(string), 400)
				return nil, false, nil
			} else if _, ok := value["default"]; !ok {
				continue
			} else {
				value1 = value["default"].(string)
			}
		}
		if a2 := value["regexp"].(string); "" != a2 {
			a3, e := regexp.MatchString(a2, value1)
			if nil != e {
				return nil, false, farmer.NewFarmerError(e)
			} else if !a3 {
				http.Error(w, "错误的参数, regexp, "+value["regexp"].(string)+", "+value1+
					", "+value["name"].(string), 400)
				return nil, false, nil
			}
		}
		if a2, ok := value["max"]; ok {
			if "hex" == value["type"].(string) {
				if strings.HasPrefix(value1, "0x") &&
					2 <= len(value1) {
					value1 = value1[2:]
				}
				a3 := a2.(float64)
				a4, e := strconv.ParseInt(value1, 16, 64)
				if nil != e {
					return nil, false, farmer.NewFarmerError(e)
				} else if a3 < float64(a4) {
					http.Error(w, "错误的参数, max, "+value["name"].(string), 400)
					return nil, false, nil
				}
			} else {
				a3 := a2.(float64)
				a4, e := strconv.ParseFloat(value1, 64)
				if nil != e {
					return nil, false, farmer.NewFarmerError(e)
				} else if a3 < a4 {
					http.Error(w, "错误的参数, max, "+value["name"].(string), 400)
					return nil, false, nil
				}
			}
		}
		if a2, ok := value["min"]; ok {
			if "hex" == value["type"].(string) {
				if strings.HasPrefix(value1, "0x") &&
					2 <= len(value1) {
					value1 = value1[2:]
				}
				a3 := a2.(float64)
				a4, e := strconv.ParseInt(value1, 16, 64)
				if nil != e {
					return nil, false, farmer.NewFarmerError(e)
				} else if float64(a4) < a3 {
					http.Error(w, "错误的参数, min, "+value["name"].(string), 400)
					return nil, false, nil
				}
			} else {
				a3 := a2.(float64)
				a4, e := strconv.ParseFloat(value1, 64)
				if nil != e {
					return nil, false, farmer.NewFarmerError(e)
				} else if a4 < a3 {
					http.Error(w, "错误的参数, min, "+value["name"].(string), 400)
					return nil, false, nil
				}
			}
		}
		if a2, ok := value["not"]; ok && value1 == a2.(string) {
			http.Error(w, "错误的参数, not, "+value["name"].(string), 400)
			return nil, false, nil
		}

		// 赋值
		// type1 - 属性的类型(reflect.Type)
		// type2 - 属性的类型(string)
		// a2 - (StructField)
		a2, ok := api.ArgumentType.FieldByName(key)
		if !ok {
			return nil, false, nil
		}
		type1 := a2.Type
		type2 := a2.Type.String()
		if "hex" == value["type"].(string) {
			// 十六进制
			a4, e := strconv.ParseInt(value1, 16, 64)
			if nil != e {
				return nil, false, farmer.NewFarmerError(e)
			}
			result1.FieldByName(key).Set(reflect.ValueOf(a4).Convert(type1))
		} else if a3, _ := regexp.MatchString(`int|byte|rune`, type2); a3 {
			// 整数
			a4, e := strconv.ParseInt(value1, 10, 64)
			if nil != e {
				return nil, false, farmer.NewFarmerError(e)
			}
			result1.FieldByName(key).Set(reflect.ValueOf(a4).Convert(type1))
		} else if strings.Contains(type2, "float") {
			// 小数
			a4, e := strconv.ParseFloat(value1, 64)
			if nil != e {
				return nil, false, farmer.NewFarmerError(e)
			}
			result1.FieldByName(key).Set(reflect.ValueOf(a4).Convert(type1))
		} else if a3, _ := regexp.MatchString(`^true$|^false$`, value1); a3 {
			a4, e := strconv.ParseBool(value1)
			if nil != e {
				return nil, false, farmer.NewFarmerError(e)
			}
			result1.FieldByName(key).Set(reflect.ValueOf(a4).Convert(type1))
		} else {
			result1.FieldByName(key).Set(reflect.ValueOf(value1))
		}
	}
	return result1.Interface(), true, nil
}
