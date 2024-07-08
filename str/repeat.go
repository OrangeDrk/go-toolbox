package strUtil

import "strings"

// Repeat 重复字符串
func Repeat(str string, count int) string {
	return strings.Repeat(str, count)
}

// RepeatByLength 重复某个字符串到指定长度
func RepeatByLength(str string, padLen int) string {
	if IsEmpty(str) {
		return ""
	}
	if padLen <= 0 {
		return ""
	}

	if len(str) == padLen {
		return str
	} else if len(str) > padLen {
		return str[:padLen]
	}
	strRune := []rune(str)
	var padding []rune
	for i := 0; i < padLen; i++ {
		padding = append(padding, strRune[i%len(str)])
	}
	return string(padding)

}
