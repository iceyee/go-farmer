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
	t.Log("\n" +
		CreateMenu(
			"1.封禁",
			"2.解封",
			"3.充值",
			"4.查询用户",
			"5.统计数据"))
	return
}

func ExampleCreateMenu() {
	println("\n" +
		CreateMenu(
			"0.Zero",
			"1.One",
			"2.Two",
			"3.Three"))
	println("\n" +
		CreateMenu(
			"1.封禁",
			"2.解封",
			"3.充值",
			"4.查询用户",
			"5.统计数据"))
	return
}

func ExampleInput() {
	Input(nil, "")
	var a001 string
	Input(&a001, "请输入xxx:")
	return
}
