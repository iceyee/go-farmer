package ftype

import (
//
)

// Stringer. 与fmt.Stringer兼容.
type Stringer interface {
	String() string
}
