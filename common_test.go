package main

import (
	"fmt"
	listUtil "github.com/OrangeDrk/go-toolbox/coll/list"
	"testing"
)

func TestList(t *testing.T) {
	list := listUtil.Of(1, 2, 3, 4, 5)
	//fmt.Printf("%#v\n", list)
	//
	//listUtil.Append(&list, 1000, 2000, 3000, 500)
	//fmt.Printf("%#v\n", list)
	//
	//list = append(list, 222)
	//fmt.Printf("%#v\n", list)
	//
	//listUtil.Page(list, -1, 2)

	//page := listUtil.Page(list, 1, 2)
	//page2 := listUtil.Page(list, 2, 2)
	//page3 := listUtil.Page(list, 3, 2)
	//fmt.Printf("%#v\n", page)
	//fmt.Printf("%#v\n", page2)
	//fmt.Printf("%#v\n", page3)
	listUtil.Sort(list, func(i, j int) bool {
		return i > j
	})
	fmt.Printf("sott: %#v\n", list)
	l2 := listUtil.Of(3, 8, 1, 23.5, 31)
	listUtil.Reverse(l2)
	fmt.Printf("reverse: %#v\n", l2)

	l3 := listUtil.Of("z", "a", "b", "ccc", "sxh")
	//listUtil.Reverse(l3)
	//fmt.Printf("sott: %#v\n", l3)
	listUtil.Sort(list, func(i, j int) bool {
		return i > j
	})
	fmt.Printf("sott: %#v\n", l3)
	reverseNew := listUtil.ReverseNew(l3)
	fmt.Printf("ReverseNew: %#v\n", reverseNew)
	fmt.Printf("ReverseOld: %#v\n", l3)

}

func TestSetOrAppend(t *testing.T) {
	l1 := listUtil.Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	listUtil.SetOrAppend(&l1, 20, 999)
	listUtil.SetOrAppend(&l1, 7, 888)
	fmt.Printf("--->%#v\n", l1)
}

func TestAppend(t *testing.T) {
	l1 := listUtil.Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	listUtil.Append(&l1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1)
	fmt.Printf("--->%#v\n", l1)
	listUtil.Reverse(l1)
	fmt.Printf("--->%#v\n", l1)
}

func TestLastIndexOf(t *testing.T) {
	l1 := listUtil.Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 5)
	listUtil.Append(&l1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1)
	lastIndexOf := listUtil.LastIndexOf(l1, func(i int) bool { return i == 5 })
	fmt.Printf("--->%#v\n", lastIndexOf)
	fmt.Printf("--->%#v\n", listUtil.IndexOfAll(l1, func(i int) bool {
		return i == 1
	}))

}

func TestSplit(t *testing.T) {
	l1 := listUtil.Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 5, 1, 1, 1, 1)
	split := listUtil.Split(l1, 7)
	fmt.Printf("--->%+v\n", split)
	avg := listUtil.SplitAvg(l1, 4)
	fmt.Printf("--->%+v\n", avg)
}
