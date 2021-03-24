package webframework

import (
	// TODO
	//
	"reflect"
)

// 根据这个类可得知路由规则, 并且由些生成文档
type ApiDocument struct {
	ArgumentType reflect.Type                      // Url参数
	Description  string                            // 描述
	Key          string                            // 用于排序
	MapTo        string                            // 映射到Controller的方法中, 并且该方法参数一定是(ResponseWrite, *Request, interface{})
	Method       string                            // 请求方法
	Remarks      string                            // 备注
	Response     string                            // 响应
	Url          string                            // 链接
	a1           map[string]map[string]interface{} // 由ArgumentType处理得到, 用于解析url参数及参数加载
	document     string                            // 文档, 由ControllerRegistry处理
	parameters   string                            // 对所有参数的说明, 由ControllerRegistry处理
	processor    reflect.Value                     // 处理请求的方法
}
