package fconvert

import (
	"github.com/iceyee/go-farmer/v4/fassert"
	"testing"
	//
)

func Test(t *testing.T) {
	fassert.Assert(9 == S2I("9"), "")
	fassert.Assert(9 == S2F("9"), "")
	fassert.Assert(true == S2B("true"), "")
	fassert.Assert("9" == I2S(9), "")
	fassert.Assert("9" == F2S(9.0), "")
	fassert.Assert("true" == B2S(true), "")
	return
}

func Example() {
	fassert.Assert(9 == S2I("9"), "")
	fassert.Assert(9 == S2F("9"), "")
	fassert.Assert(true == S2B("true"), "")
	fassert.Assert("9" == I2S(9), "")
	fassert.Assert("9" == F2S(9.0), "")
	fassert.Assert("true" == B2S(true), "")
	return
}
