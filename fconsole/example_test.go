package fconsole

import (
	"testing"
	//
)

func Test(t *testing.T) {
	return
}

func TestBanner(t *testing.T) {
	t.Log(Banner)
	return
}

func TestCreateMenu(t *testing.T) {
	t.Log("\n" +
		CreateMenu(
			"0.Zero",
			"1.One",
			"2.Two",
			"3.Three"))
	return
}
