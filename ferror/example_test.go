package ferror

import (
	"testing"
	//
)

func TestFarmerError(t *testing.T) {
	t.Log(New("hello world.").Error())
	return
}

func Example() {
	println(New("hello world.").Error())
	return
}
