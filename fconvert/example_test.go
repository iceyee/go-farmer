package fconvert

import (
	"github.com/iceyee/go-farmer/v5/fassert"
	"testing"
	//
)

func Test(t *testing.T) {
	fassert.Assert(9 == S2I("9"), `9 == S2I("9")`)
	fassert.Assert(9 == S2F("9"), `9 == S2F("9")`)
	fassert.Assert(true == S2B("true"), `true == S2B("true")`)
	fassert.Assert("9" == I2S(9), `"9" == I2S(9)`)
	fassert.Assert("9" == F2S(9.0), `"9" == F2S(9.0)`)
	fassert.Assert("true" == B2S(true), `"true" == B2S(true)`)
	return
}

func Example() {
	fassert.Assert(9 == S2I("9"), `9 == S2I("9")`)
	fassert.Assert(9 == S2F("9"), `9 == S2F("9")`)
	fassert.Assert(true == S2B("true"), `true == S2B("true")`)
	fassert.Assert("9" == I2S(9), `"9" == I2S(9)`)
	fassert.Assert("9" == F2S(9.0), `"9" == F2S(9.0)`)
	fassert.Assert("true" == B2S(true), `"true" == B2S(true)`)
	return
}
