package strUtil

import (
	"regexp"
	"strings"
)

// Replace 替换字符串
func Replace(source string, searchStr string, replacement string, replaceCount int) string {
	return strings.Replace(source, searchStr, replacement, replaceCount)
}

// ReplaceAll 替换字符串
func ReplaceAll(source string, searchStr string, replacement string) string {
	return Replace(source, searchStr, replacement, -1)
}

// ReplaceWithIndex 按照指定区间替换字符串
// startIndex 开始位置（包含）
// endIndex 结束位置（不包含）
func ReplaceWithIndex(str string, startIndex int, endIndex int, replacedStr string) string {
	if startIndex < 0 || endIndex > len(str) || startIndex > endIndex {
		return str // 如果索引无效，返回原始字符串
	}
	runes := []rune(str)
	strLen := Length(str)

	// 检查并调整 startIndex 和 endIndex 的有效性
	if startIndex < 0 {
		startIndex = 0
	}
	if endIndex > strLen {
		endIndex = strLen
	}
	if startIndex >= endIndex {
		return str
	}

	replaceCount := endIndex - startIndex

	return string(runes[:startIndex]) + RepeatByLength(replacedStr, replaceCount) + string(runes[endIndex:])
}

// ReplaceIgnoreCase 替换字符串
func ReplaceIgnoreCase(source string, searchStr string, replacement string, replaceCount int) string {
	return Replace(strings.ToLower(source), strings.ToLower(searchStr), replacement, replaceCount)
}

// ReplaceAllIgnoreCase 替换字符串
func ReplaceAllIgnoreCase(source string, searchStr string, replacement string) string {
	return Replace(strings.ToLower(source), strings.ToLower(searchStr), replacement, -1)
}

// ReplaceWithMatcher 通过正则表达式替换字符串
func ReplaceWithMatcher(str string, regex string, replaceFun func(string) string) string {
	re := regexp.MustCompile(regex)
	return re.ReplaceAllStringFunc(str, replaceFun)
}

// Hide 替换指定字符串的指定区间内字符为"*" 俗称：脱敏功能
func Hide(str string, startInclude int, endExclude int) string {
	return ReplaceWithIndex(str, startInclude, endExclude, "*")
}
