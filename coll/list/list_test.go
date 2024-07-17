package listUtil

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	list := Of(1, 2, 3, 4, 5)
	fmt.Printf("%#v\n", list)
	Append(&list, 1000, 2000, 3000, 500)
	fmt.Printf("%#v\n", list)
	list = append(list, 222)
	fmt.Printf("%#v\n", list)

}
