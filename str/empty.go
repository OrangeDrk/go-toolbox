package strUtil

import "strings"

// IsEmpty 检测是否是空串
func IsEmpty(sourceStr string) bool {
	trimmedStr := strings.TrimSpace(sourceStr)
	return trimmedStr == ""
}

// IsNotEmpty 检测是否是空串
func IsNotEmpty(sourceStr string) bool {
	return !IsEmpty(sourceStr)
}

// HasEmpty 提供的字符串数组是否包含空串
func HasEmpty(strs []string) bool {
	if len(strs) == 0 {
		return true
	}
	for _, str := range strs {
		if IsEmpty(str) {
			return true
		}
	}
	return false
}

// IsAllEmpty 提供的字符串数组是否全是空串
func IsAllEmpty(strs []string) bool {
	if len(strs) == 0 {
		return true
	}
	for _, str := range strs {
		if IsNotEmpty(str) {
			return false
		}
	}
	return true
}

// IsAllNotEmpty 指定字符串数组中的元素，是否都不为空字符串。
func IsAllNotEmpty(strs []string) bool {
	return false == HasEmpty(strs)
}

// IsUndefined 判断字符串是否等于 "undefined"
func IsUndefined(str string) bool {
	trimStr := strings.TrimSpace(str)
	return "undefined" == trimStr
}

// EmptyToDefault 如果字符串是""，则返回指定默认字符串，否则返回字符串本身。
func EmptyToDefault(str string, defaultStr string) string {
	if IsEmpty(str) {
		return defaultStr
	} else {
		return str
	}
}
