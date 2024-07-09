package strUtil

import "strings"

// Trim 除去字符串头尾部的空白
func Trim(str string) string {
	if IsEmpty(str) {
		return str
	}
	return strings.TrimSpace(str)
}

// TrimStart 除去字符串头部的空白
func TrimStart(str string) string {
	return strings.TrimLeftFunc(str, func(r rune) bool {
		if ' ' == r {
			return true
		}
		return false
	})
}

// TrimEnd 除去字符串尾部的空白
func TrimEnd(str string) string {
	return strings.TrimRightFunc(str, func(r rune) bool {
		if ' ' == r {
			return true
		}
		return false
	})
}

// CleanEmpty 清除空白串
func CleanEmpty(str string) string {
	strRune := []rune(str)
	var newRune []rune
	for _, r := range strRune {
		if r != ' ' {
			newRune = append(newRune, r)
		}
	}
	return string(newRune)
}
