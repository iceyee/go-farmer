package ftype

import (
//
)

// 异常. 与error完全相同.
type Error interface {
	Error() string
}
