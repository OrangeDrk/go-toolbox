package desensitizedUtil

import (
	"fmt"
	strUtil "github.com/OrangeDrk/go-toolbox/str"
	"testing"
)

func TestDesensitized(t *testing.T) {
	fmt.Printf("%#v\n", Desensitized("17603393007", MobilePhone))
	fmt.Printf("%#v\n", Desensitized("石晓浩", ChineseName))
	fmt.Printf("%#v\n", Desensitized("李赵一特", ChineseName))

	fmt.Printf("%#v\n", Desensitized("131182199702251631", IDCard))
	fmt.Printf("%#v\n", Desensitized("北京市海淀区清华东路西口29#1#601", Address))
	fmt.Printf("%#v\n", strUtil.Length("北京市海淀区清华东路西口29#1#601"))
	fmt.Printf("%#v\n", Desensitized("北京市海淀区清华东路西口29#1#601", Email))
	fmt.Printf("%#v\n", Desensitized("12313123123", Password))
	fmt.Printf("%#v\n", Desensitized("苏D40000", CarLicense))
	fmt.Printf("%#v\n", Desensitized("陕A12345D", CarLicense))
	fmt.Printf("%#v\n", Desensitized("1234    2222 3333 4444 6789 9", BankCard))
	fmt.Printf("%#v\n", Desensitized("1234 2222 3333 4444 6789 91", BankCard))
	fmt.Printf("%#v\n", Desensitized("1234 2222 3333 4444 678", BankCard))
	fmt.Printf("%#v\n", Desensitized("1234 2222 3333 4444 6789", BankCard))
	fmt.Printf("%#v\n", Desensitized("127.0.0.1", IPV4))
	fmt.Printf("%#v\n", Desensitized("2001:0db8:86a3:08d3:1319:8a2e:0370:7344", IPV6))
}
