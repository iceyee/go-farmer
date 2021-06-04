package ferror

import (
	"testing"
	//
)

func TestFarmerError(t *testing.T) {
	t.Log(New("hello world.").Error())
	return
}

func ExampleFarmerError() {
	println(New("hello world.").Error())
	return
}
