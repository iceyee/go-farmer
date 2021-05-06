package ffile

import (
//
)

// 在主目录下添加目录farmer/ ,
// 包含子目录bin/ etc/ include/ lib/ share/ src/
func InstallDirectory() {
	var a001 []string
	a001 = []string{
		"bin",
		"etc",
		"include",
		"lib",
		"share",
		"src",
	}
	for _, x := range a001 {
		var path string
		path = Path(HomeDirectory, "farmer", x)
		Mkdir(path)
	}
	return
}
