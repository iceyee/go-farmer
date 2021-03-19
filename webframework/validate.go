package webframework

import (
	// TODO
	//
	"github.com/iceyee/go-farmer/farmer"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

// 做表单验证, 返回false则表示不符合预期,
// result - 传入一个结构体指针, 该结构体的属性对应表单.
// 将预期条件写在属性的标签中, 支持以下标签
// name - 表单的单列, 不写则表示忽略该属性.
// require - true或false, 表示这一列是否必须.
// enum - 枚举, 用逗号间隔, 表示需要从这枚举中选一个.
// min - 最小值, 需要是数字
// max - 最大值, 需要是数字
// default - 默认值
// 支持的类型包括int, int8, int16, int32, int64, float32, float64, string
func Validate(w http.ResponseWriter, r *http.Request, result interface{}) (bool, error) {
	value1 := reflect.ValueOf(result)
	if "ptr" != value1.Kind().String() {
		e := farmer.NewFarmerError("需要传入结构体指针")
		http.Error(w, e.Error(), 500)
		return false, e
	}
	value2 := value1.Elem()
	type2 := value2.Type()
	for x := 0; x < value2.NumField(); x++ {
		field1 := type2.Field(x)
		field2 := value2.Field(x)
		println(type2.Name())
		println(string(field1.Tag))

		name1 := field1.Tag.Get("name")
		if "" == name1 {
			// 没有name, 忽略
			continue
		}

		require1 := "true" == field1.Tag.Get("require")
		value3 := r.FormValue(name1)
		if require1 &&
			"" == value3 {
			http.Error(w, "参数错误", 400)
			return false, nil
		}

		var default1 string
		if "" == value3 {
			default1 = field1.Tag.Get("default")
		} else {
			enum1 := field1.Tag.Get("enum")
			if "" != enum1 {
				enum2 := "," + enum1 + ","
				if !strings.Contains(enum2, value3) {
					http.Error(w, "参数错误", 400)
					return false, nil
				}
			}

			min1 := field1.Tag.Get("min")
			if "" != min1 {
				min2, e := strconv.ParseFloat(min1, 64)
				if nil != e {
					e = farmer.NewFarmerError(e)
					http.Error(w, e.Error(), 500)
					return false, e
				}
				value4, e := strconv.ParseFloat(value3, 64)
				if nil != e {
					e = farmer.NewFarmerError(e)
					http.Error(w, e.Error(), 500)
					return false, e
				}
				if value4 < min2 {
					http.Error(w, "参数错误", 400)
					return false, nil
				}
			}

			max1 := field1.Tag.Get("max")
			if "" != max1 {
				max2, e := strconv.ParseFloat(max1, 64)
				if nil != e {
					e = farmer.NewFarmerError(e)
					http.Error(w, e.Error(), 500)
					return false, e
				}
				value5, e := strconv.ParseFloat(value3, 64)
				if nil != e {
					e = farmer.NewFarmerError(e)
					http.Error(w, e.Error(), 500)
					return false, e
				}
				if max2 < value5 {
					http.Error(w, "参数错误", 400)
					return false, nil
				}
			}

			default1 = value3
		}

		type1 := field2.Kind().String()
		if "" != default1 {
			if "string" == type1 {
				field2.Set(reflect.ValueOf(default1))
			} else if "float64" == type1 ||
				"float32" == type1 {
				default2, e := strconv.ParseFloat(default1, 64)
				if nil != e {
					e = farmer.NewFarmerError(e)
					http.Error(w, e.Error(), 500)
					return false, e
				}
				if "float64" == type1 {
					field2.Set(reflect.ValueOf(default2))
				} else if "float32" == type1 {
					field2.Set(reflect.ValueOf(float32(default2)))
				}
			} else if "int" == type1 ||
				"int8" == type1 ||
				"int16" == type1 ||
				"int32" == type1 ||
				"int64" == type1 {
				default2, e := strconv.ParseInt(default1, 10, 64)
				if nil != e {
					e = farmer.NewFarmerError(e)
					http.Error(w, e.Error(), 500)
					return false, e
				}
				if "int" == type1 {
					field2.Set(reflect.ValueOf(int(default2)))
				} else if "int8" == type1 {
					field2.Set(reflect.ValueOf(int8(default2)))
				} else if "int16" == type1 {
					field2.Set(reflect.ValueOf(int16(default2)))
				} else if "int32" == type1 {
					field2.Set(reflect.ValueOf(int32(default2)))
				} else if "int64" == type1 {
					field2.Set(reflect.ValueOf(int64(default2)))
				}
			} else {
				e := farmer.NewFarmerError("不支持的类型")
				http.Error(w, e.Error(), 500)
				return false, e
			}
		}
	}
	return true, nil
}
