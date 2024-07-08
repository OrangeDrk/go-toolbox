package strUtil

import "strings"

// RemoveAll 移除字符串中所有给定字符串;removeAll("aa-bb-cc-dd", "-") =》 aabbccdd
func RemoveAll(str string, strToRemove string) string {
	if str == "" || strToRemove == "" {
		return str
	}
	return ReplaceAll(str, strToRemove, "")
}

// RemoveAny 移除字符串中所有给定字符串，当某个字符串出现多次，则全部移除
func RemoveAny(str string, strsToRemove []string) string {
	var result = str
	if IsNotEmpty(str) {
		for _, s := range strsToRemove {
			result = RemoveAll(result, s)
		}
	}
	return result
}

// RemoveAllLineBreaks 去除所有换行符，包括：\r \n
func RemoveAllLineBreaks(str string) string {
	return RemoveAny(str, []string{"\r", "\n"})
}

// RemovePrefix 去掉指定前缀
func RemovePrefix(str string, prefix string) string {
	if IsEmpty(str) || IsEmpty(prefix) {
		return str
	}
	if strings.HasPrefix(str, prefix) {
		return str[len(prefix):]
	}
	return str
}

// RemovePrefixIgnoreCase 忽略大小写去掉指定前缀
func RemovePrefixIgnoreCase(str string, prefix string) string {
	str = strings.ToLower(str)
	prefix = strings.ToLower(prefix)
	return RemovePrefix(str, prefix)
}

// RemoveSuffix 去掉指定后缀
func RemoveSuffix(str string, suffix string) string {
	if IsEmpty(str) || IsEmpty(suffix) {
		return str
	}
	if strings.HasSuffix(str, suffix) {
		return str[:len(suffix)]
	}
	return str
}

// RemoveSuffixIgnoreCase 去掉指定后缀(忽略大小写)
func RemoveSuffixIgnoreCase(str string, suffix string) string {
	return RemoveSuffix(strings.ToLower(str), strings.ToLower(suffix))
}
