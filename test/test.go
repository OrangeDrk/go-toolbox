package main

import (
	"fmt"
	idUtil "github.com/OrangeDrk/go-toolbox/id"
	numberUtil "github.com/OrangeDrk/go-toolbox/number"
	strUtil "github.com/OrangeDrk/go-toolbox/str"
)

func test() {
	str1 := "sdfh1320fsajasd"
	//str2 := "abc456jk"
	str2 := "sdgsd123ebsd"
	similar := strUtil.Similar(str1, str2)
	similarScale := strUtil.SimilarScale(str1, str2, 0)
	if similar >= 0.3 {
		fmt.Println(">")
	}
	fmt.Printf("%.2f\n", similar)
	fmt.Printf("%s\n", similarScale)
	numbers := []int{0, -2, 3, 1}

	fmt.Printf("%v", numberUtil.Min(numbers))
	id := idUtil.GetSnowFlakeNextId()
	str := idUtil.GetSnowFlakeNextIdStr()
	fmt.Printf("%v\n", id)
	fmt.Printf("%v\n", str)

	for i := 0; i < 10; i++ {
		id := idUtil.GetSnowFlakeNextId()
		fmt.Printf("%v\n", id)
	}
}

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

func main() {
	testStrFormat()
}
