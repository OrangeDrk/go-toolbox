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

// RepeatAndJoin 重复某个字符串并通过分界符连接
func RepeatAndJoin(str string, delimiter string, count int) string {
	if count <= 0 {
		return ""
	}

	// 创建一个切片，用于存储重复的字符串
	repeatedStrings := make([]string, count)
	for i := 0; i < count; i++ {
		repeatedStrings[i] = str
	}

	// 使用 strings.Join 函数将切片中的字符串通过 delimiter 连接
	return strings.Join(repeatedStrings, delimiter)
}
