package ferror

import (
//
)

func ExampleFarmerError() {
	println(New("hello world.").Error())
	return
}
