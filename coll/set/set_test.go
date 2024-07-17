package setUtil

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	// 创建一个整数类型的 Set
	intSet := NewSet[int]()
	intSet.Add(1)
	intSet.Add(2)
	intSet.Add(3)
	fmt.Println("Int Set:", intSet.ToSlice()) // 输出: [1 2 3]

	intSet.Remove(2)
	fmt.Println("Contains 2:", intSet.Contains(2)) // 输出: false
	fmt.Println("Size:", intSet.Size())            // 输出: 2

	// 创建一个字符串类型的 Set
	stringSet := NewSet[string]()
	stringSet.Add("apple")
	stringSet.Add("banana")
	stringSet.Add("cherry")
	fmt.Println("String Set:", stringSet.ToSlice()) // 输出: [apple banana cherry]

	stringSet.Remove("banana")
	fmt.Println("Contains 'banana':", stringSet.Contains("banana")) // 输出: false
	fmt.Println("Size:", stringSet.Size())                          // 输出: 2

	stringSet.Clear()
	fmt.Println("Cleared String Set, Size:", stringSet.Size()) // 输出: 0

}
