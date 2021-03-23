package farmer

import (
	"math/rand"
	"testing"
	//
)

func TestFarmerError(t *testing.T) {
	println(NewFarmerError("hello world").Error())
	return
}
