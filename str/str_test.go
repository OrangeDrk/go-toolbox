package strUtil

import (
	"fmt"
	"testing"
)

func TestStr(t *testing.T) {
	str := ParseStr(1000)
	fmt.Println(str)
	str1 := ParseStr(100.4)
	fmt.Println(str1)

	str2 := ParseStr(true)
	fmt.Println(str2)

	obj := struct {
		ID   int
		Name string
	}{
		ID:   1,
		Name: "sxh",
	}
	str3 := ParseStr(obj)
	fmt.Println(str3)

	str4 := ParseStr([]byte("asdadadsd"))
	fmt.Println(str4)
}

func TestStrFormat(t *testing.T) {
	str := "今天天气:{a},上班时间:{sb},城市:{city},age:{age}, arg:{arg}"
	p := map[string]interface{}{
		"a":    "多云",
		"sb":   "8:30",
		"city": "北京",
		"age":  9,
		"arg":  []int{1, 2, 3, 4},
	}

	format := Format(str, p)
	fmt.Println(format)
}

func TestIndexFormat(t *testing.T) {
	str := "今天天气:{0},上班时间:{1},城市:{2},age:{3}, arg:{4}"
	//p := map[string]interface{}{
	//	"a":    "多云",
	//	"sb":   "8:30",
	//	"city": "北京",
	//	"age":  9,
	//	"arg":  []int{1, 2, 3, 4},
	//}
	p := []interface{}{"多云", "8:30", "beijing", 9, []int{1, 2, 3, 4, 5}}

	format := IndexedFormat(str, p)
	fmt.Println(format)
}

func TestReplace(t *testing.T) {
	str := "hello world，Hello 你好"
	replace := ReplaceIgnoreCase(str, "hello", "hi", -1)
	fmt.Println(replace)
}
