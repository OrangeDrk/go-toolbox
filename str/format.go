package strUtil

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseStr 解析字符串
func ParseStr(source interface{}) string {
	switch v := source.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v)
	case []byte:
		return string(v)
	default:
		return fmt.Sprintf("%+v", source)
	}
}

// FillBefore 将已有字符串填充为规定长度，如果已有字符串超过这个长度则返回这个字符串
func FillBefore(str string, char string, length int) string {
	return Fill(str, char, length, true)
}

// FillAfter 将已有字符串填充为规定长度，如果已有字符串超过这个长度则返回这个字符串
func FillAfter(str string, char string, length int) string {
	return Fill(str, char, length, false)
}

// Fill 将已有字符串填充为规定长度，如果已有字符串超过这个长度则返回这个字符串
func Fill(str string, char string, length int, isPre bool) string {
	if len(str) >= length {
		return str
	}
	fillLength := length - len(str)
	fillStr := strings.Repeat(char, fillLength)
	if isPre {
		return fillStr + str
	}
	return str + fillStr
}

// Format 通过map中的参数 格式化字符串
// map = {a: "aValue", b: "bValue"} format("{a} and {b}", map) ---=》 aValue and bValue
func Format(template string, params map[string]interface{}) string {
	// 遍历map中的键值对
	for key, value := range params {
		// 构造占位符，例如 "{a}"
		placeholder := fmt.Sprintf("{%s}", key)
		// 将占位符替换为对应的值
		template = strings.ReplaceAll(template, placeholder, fmt.Sprintf("%v", value))
	}
	return template
}
