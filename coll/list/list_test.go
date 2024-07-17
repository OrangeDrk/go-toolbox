package listUtil

import (
	"fmt"
	"slices"
	"testing"
)

func Test1(t *testing.T) {
	list := Of(1, 2, 3, 4, 5)
	fmt.Printf("%#v\n", list)
	Add(&list, 1000, 2000, 3000, 500)
	fmt.Printf("%#v\n", list)
	list = append(list, 222)
	fmt.Printf("%#v\n", list)
	slices.Contains(list, 1)

}

func Test2(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}
	userList := Of(User{Name: "john", Age: 18}, User{Name: "sxh", Age: 18}, User{Name: "aaa", Age: 18})
	contains := Contains(userList, User{Name: "sxh", Age: 18})
	fmt.Printf("%#v\n", contains)
}

func Test3(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}
	list1 := Of(1, 23, 4, 5, 6, 7, 88, 0)
	Remove(&list1, 23)
	fmt.Printf("%#v\n", list1)
}
