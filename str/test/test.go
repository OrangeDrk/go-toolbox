package main

import (
	"fmt"
	desensitizedUtil "github.com/OrangeDrk/go-toolbox/desensitized"
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

func testIndexFormat() {
	str := "今天天气:{0},上班时间:{1},城市:{2},age:{3}, arg:{4}"
	//p := map[string]interface{}{
	//	"a":    "多云",
	//	"sb":   "8:30",
	//	"city": "北京",
	//	"age":  9,
	//	"arg":  []int{1, 2, 3, 4},
	//}
	p := []interface{}{"多云", "8:30", "beijing", 9, []int{1, 2, 3, 4, 5}}

	format := strUtil.IndexedFormat(str, p)
	fmt.Println(format)
}

func testReplace() {
	str := "hello world，Hello 你好"
	replace := strUtil.ReplaceIgnoreCase(str, "hello", "hi", -1)
	fmt.Println(replace)
}

func main() {
	//fmt.Println(strUtil.ContainsBlank("Hello World"))        // true
	//fmt.Println(strUtil.ContainsBlank("Hello\tWorld"))       // true
	//fmt.Println(strUtil.ContainsBlank("Hello　World"))        // true （全角空格）
	//fmt.Println(strUtil.ContainsBlank("Hello     World   ")) // true （不间断空格）
	//fmt.Println(strUtil.ContainsBlank("HelloWorld"))         // false

	//str := "1;2;3=4;5;6"
	//fmt.Printf("%#v\n", strUtil.Split(str, ";"))
	//fmt.Printf("%#v\n", strUtil.SplitTrim(str, ";"))
	//fmt.Printf("%#v\n", strUtil.SplitLimit(str, ";", 2))
	//fmt.Printf("%#v\n", strUtil.SplitTrimLimit(str, ";", 2))

	//mapping := strUtil.SplitAfterMapping[int](str, ";", func(s string) (int, error) {
	//	i, err := strconv.Atoi(s)
	//	if err != nil {
	//		return 0, err
	//	}
	//	return i, nil
	//})
	//fmt.Printf("%#v\n", mapping)
	//fmt.Printf("%#v\n", strUtil.SplitByLen(str, 3))
	////fmt.Printf("%#v\n", strUtil.SubString(str, 0, 4))
	//fmt.Printf("%#v\n", strUtil.SubBefore(str, "=", false))
	//fmt.Printf("%#v\n", strUtil.SubAfter("=", "=", false))
	//fmt.Printf("%#v\n", strUtil.SubBetween("wx[b] yz", "[", "]"))
	//fmt.Printf("%#v\n", strUtil.SubBetween("yabcz", "", ""))
	//fmt.Printf("%#v\n", strUtil.SubBetween("yabcz", "y", "z"))
	//fmt.Printf("%#v\n", strUtil.SubBetween("yabczyabcz", "y", "z"))
	//fmt.Printf("%#v\n", strUtil.SubBetween("", "[", "]"))
	//fmt.Printf("%#v\n", strUtil.SubBetween("", "", "]"))
	//fmt.Printf("%#v\n", strUtil.SubBetween("", "", ""))
	//fmt.Printf("%#v\n", strUtil.SubBetweenAll("yabczyabcz", "y", "z"))
	//fmt.Printf("%#v\n", strUtil.SubBetweenAll("[yabc[zy] abcz]", "[", "]"))
	//fmt.Printf("%#v\n", strUtil.SubBetweenAll("yabcz", "y", "z"))
	//fmt.Printf("%#v\n", strUtil.SubBetweenAll("", "[", "]"))
	//fmt.Printf("%#v\n", strUtil.SubBetweenAll("wx[b] y[z]", "[", "]"))
	//fmt.Printf("%#v\n", strUtil.Repeat("sxh", 2))
	//fmt.Printf("%#v\n", strUtil.RepeatByLength("sxh", 5))
	//fmt.Printf("%#v\n", strUtil.RepeatAndJoin("sxh", ";", 5))
	//fmt.Printf("%#v\n", strUtil.EqualsAt("sxh", 1, "xh"))
	////testIndexFormat()
	//fmt.Printf("%#v\n", strUtil.PadPre("sxh", 5, "666"))
	//fmt.Printf("%#v\n", strUtil.PadPre("sxh", -19, "666"))
	//
	//fmt.Printf("%#v\n", strUtil.PadAfter("sxh", -9, "666"))
	//fmt.Printf("%#v\n", strUtil.PadAfter("sxh", 10, "666"))
	//fmt.Printf("%#v\n", strUtil.Count("sxh666,ssr", "s"))
	//fmt.Printf("%#v\n", strUtil.IndexOfByRangeStart("sxh666,ssr", "s", 7))

	//str := "sxh666,ssrrrr"
	//fmt.Printf("%#v\n", strUtil.IndexOf(str, "s"))
	//fmt.Printf("%#v\n", strUtil.IndexOfByRangeStart(str, "s", 10))
	//fmt.Printf("%#v\n", strUtil.IndexOfByRange(str, "s", 3, 8))
	//fmt.Printf("%#v\n", strUtil.IndexOfByRange(str, "s", 3, 4))
	//fmt.Printf("%#v\n", strUtil.LastIndexOf(str, "s"))
	//fmt.Printf("%#v\n", strUtil.LastIndexOfIgnoreCaseByRangeStart(str, "s", 7))
	//fmt.Printf("%#v\n", strUtil.ReplaceWithMatcher(str, "(\\d)", func(s string) string {
	//	return "-" + s + "-"
	//}))
	//fmt.Printf("%#v\n", strUtil.Hide(str, 1, 7))
	//fmt.Printf("%#v\n", strUtil.RepeatByLength("abcd", 3))
	fmt.Printf("%#v\n", desensitizedUtil.Desensitized("17603393007", desensitizedUtil.MobilePhone))
	fmt.Printf("%#v\n", desensitizedUtil.Desensitized("石晓浩", desensitizedUtil.ChineseName))
	fmt.Printf("%#v\n", desensitizedUtil.Desensitized("李赵一特", desensitizedUtil.ChineseName))

	fmt.Printf("%#v\n", desensitizedUtil.Desensitized("131182199702251631", desensitizedUtil.IDCard))
	fmt.Printf("%#v\n", desensitizedUtil.Desensitized("北京市海淀区清华东路西口29#1#601", desensitizedUtil.Address))
	fmt.Printf("%#v\n", strUtil.Length("北京市海淀区清华东路西口29#1#601"))
	fmt.Printf("%#v\n", desensitizedUtil.Desensitized("北京市海淀区清华东路西口29#1#601", desensitizedUtil.Email))
	fmt.Printf("%#v\n", desensitizedUtil.Desensitized("12313123123", desensitizedUtil.Password))
	fmt.Printf("%#v\n", desensitizedUtil.Desensitized("苏D40000", desensitizedUtil.CarLicense))
	fmt.Printf("%#v\n", desensitizedUtil.Desensitized("陕A12345D", desensitizedUtil.CarLicense))
	fmt.Printf("%#v\n", desensitizedUtil.Desensitized("1234    2222 3333 4444 6789 9", desensitizedUtil.BankCard))
	fmt.Printf("%#v\n", desensitizedUtil.Desensitized("1234 2222 3333 4444 6789 91", desensitizedUtil.BankCard))
	fmt.Printf("%#v\n", desensitizedUtil.Desensitized("1234 2222 3333 4444 678", desensitizedUtil.BankCard))
	fmt.Printf("%#v\n", desensitizedUtil.Desensitized("1234 2222 3333 4444 6789", desensitizedUtil.BankCard))
	fmt.Printf("%#v\n", desensitizedUtil.Desensitized("127.0.0.1", desensitizedUtil.IPV4))
	fmt.Printf("%#v\n", desensitizedUtil.Desensitized("2001:0db8:86a3:08d3:1319:8a2e:0370:7344", desensitizedUtil.IPV6))
}
