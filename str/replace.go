package strUtil

import "strings"

// Replace 替换字符串
func Replace(source string, searchStr string, replacement string, replaceCount int) string {
	return strings.Replace(source, searchStr, replacement, replaceCount)
}

// ReplaceAll 替换字符串
func ReplaceAll(source string, searchStr string, replacement string) string {
	return Replace(source, searchStr, replacement, -1)
}

// ReplaceIgnoreCase 替换字符串
func ReplaceIgnoreCase(source string, searchStr string, replacement string, replaceCount int) string {
	return Replace(strings.ToLower(source), strings.ToLower(searchStr), replacement, replaceCount)
}

// ReplaceAllIgnoreCase 替换字符串
func ReplaceAllIgnoreCase(source string, searchStr string, replacement string) string {
	return Replace(strings.ToLower(source), strings.ToLower(searchStr), replacement, -1)
}
