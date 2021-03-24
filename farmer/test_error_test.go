package farmer

import (
	"testing"
	//
)

func TestFarmerError(t *testing.T) {
	println(NewFarmerError("hello world").Error())
	return
}
