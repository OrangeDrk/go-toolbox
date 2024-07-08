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

func main() {
}
