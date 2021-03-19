package webframework

import (
// TODO
//
)

// 在Controller中写方法func() ApiDocument, 在注册时会同时生成文档
type ApiDocument struct {
	Description string // 描述
	Key         string // 用于排序
	Method      string // 请求方法
	Parameters  string // 参数,注明必须还是可选
	Remarks     string // 备注
	Response    string // 响应
	Url         string // 链接
}
