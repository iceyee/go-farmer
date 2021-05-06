package ftype

import (
//
)

// 异常. 与error兼容.
type Error interface {
	Error() string
}
