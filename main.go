package main

import (
	"fmt"
	strUtil "github.com/OrangeDrk/go-toolbox/base"
	"strings"
)

func main() {

	str:="石晓浩176"
	index := strings.Index(str, "7")
	fmt.Println(index)

	indexOf := strUtil.IndexOf("石晓浩176", '7')
	fmt.Println(indexOf)

	fmt.Println(strUtil.SubString(str,-9,10))
	fmt.Println(strUtil.SubString(str,9,10))
}