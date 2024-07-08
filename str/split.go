package strUtil

import "strings"

// Split 分割字符串
func Split(str string, separator string) []string {
	return strings.Split(str, separator)
}

// SplitLimit 分割字符串，限制分片数
func SplitLimit(str string, separator string, limit int) []string {
	return strings.SplitN(str, separator, limit)
}

// SplitTrim 切分字符串，去除切分后每个元素两边的空白符，去除空白项
func SplitTrim(str string, separator string) []string {
	split := Split(str, separator)
	var result []string
	for _, s := range split {
		if IsNotEmpty(s) {
			result = append(result, Trim(s))
		}
	}
	return result
}

// SplitTrimLimit 切分字符串，去除切分后每个元素两边的空白符，去除空白项，限制分片数
func SplitTrimLimit(str string, separator string, limit int) []string {
	split := SplitLimit(str, separator, limit)
	var result []string
	for _, s := range split {
		if IsNotEmpty(s) {
			result = append(result, Trim(s))
		}
	}
	return result
}

// SplitAfterMapping 切分字符串,切分之后通过mapping转换并返回
func SplitAfterMapping[T any](str string, separator string, mapping func(s string) (T, error)) []T {
	splitLimit := Split(str, separator)
	var result []T
	for _, s := range splitLimit {
		mapping, err := mapping(s)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, mapping)
	}
	return result
}

// SplitByLen 根据给定长度，将给定字符串截取为多个部分
func SplitByLen(str string, length int) []string {
	if length <= 0 || str == "" {
		return []string{}
	}

	var result []string
	runeStr := []rune(str)
	for i := 0; i < len(runeStr); i += length {
		end := i + length
		if end > len(runeStr) {
			end = len(runeStr)
		}
		result = append(result, string(runeStr[i:end]))
	}
	return result
}

// Cut 将字符串切分为n等份
func Cut(str string, n int) []string {
	if n <= 0 || str == "" {
		return []string{}
	}

	runeStr := []rune(str)
	strLen := len(runeStr)
	partLen := strLen / n
	remainder := strLen % n

	var result []string
	start := 0

	for i := 0; i < n; i++ {
		end := start + partLen
		if i < remainder { // 分配余数
			end++
		}
		result = append(result, string(runeStr[start:end]))
		start = end
	}

	return result
}
