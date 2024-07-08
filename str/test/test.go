package main

import (
	"fmt"
	strUtil "github.com/OrangeDrk/go-toolbox/str"
)

func testStr() {
	str := strUtil.ParseStr(1000)
	fmt.Println(str)
	str1 := strUtil.ParseStr(100.4)
	fmt.Println(str1)

	str2 := strUtil.ParseStr(true)
	fmt.Println(str2)

	obj := struct {
		ID   int
		Name string
	}{
		ID:   1,
		Name: "sxh",
	}
	str3 := strUtil.ParseStr(obj)
	fmt.Println(str3)

	str4 := strUtil.ParseStr([]byte("asdadadsd"))
	fmt.Println(str4)
}

func testStrFormat() {
	str := "今天天气:{a},上班时间:{sb},城市:{city},age:{age}, arg:{arg}"
	p := map[string]interface{}{
		"a":    "多云",
		"sb":   "8:30",
		"city": "北京",
		"age":  9,
		"arg":  []int{1, 2, 3, 4},
	}

	format := strUtil.Format(str, p)
	fmt.Println(format)
}

func testReplace() {
	str := "hello world，Hello 你好"
	replace := strUtil.ReplaceIgnoreCase(str, "hello", "hi", -1)
	fmt.Println(replace)
}

func main() {
	str := "Sxh12313"
	prefix := strUtil.RemovePrefixIgnoreCase(str, "sxh")
	fmt.Println(prefix)
}