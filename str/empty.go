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
