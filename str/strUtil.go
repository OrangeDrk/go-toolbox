package strUtil

import (
	"fmt"
	numberUtil "github.com/OrangeDrk/go-toolbox/number"
	"strconv"
	"strings"
)

// IsEmpty 检测是否是空串
func IsEmpty(sourceStr string) bool {
	trimmedStr := strings.TrimSpace(sourceStr)
	return trimmedStr == ""
}

// IndexOf 返回字符在原始字符串的下标
func IndexOf(sourceStr string, subStr rune) int {
	for index, char := range sourceStr {
		if char == subStr {
			return index
		}
	}
	return -1
}

// SubString 切割原始字符，返回子串 范围：[s,e)
func SubString(sourceStr string, s int, e int) string {
	if s < 0 || e > len(sourceStr) || s > e {
		return ""
	}
	return sourceStr[s:e]
}

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
		return fmt.Sprintf("%v", source)
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

// levenshteinDistance 计算两个字符串的相似度
// 返回的是两个字符串编辑距离，如果两个字符串相等，那么编辑距离为0，完全不相等，编辑距离距离就是最长字符串的长度
// 返回的数值越大，表示相似度越低
func levenshteinDistance(str1 string, str2 string) int {

	len1 := len(str1)
	len2 := len(str2)

	// 创建二维数组用于存储编辑距离
	dist := make([][]int, len1+1)
	for i := range dist {
		dist[i] = make([]int, len2+1)
	}

	// 初始化第一行和第一列
	for i := 0; i <= len1; i++ {
		dist[i][0] = i
	}
	for j := 0; j <= len2; j++ {
		dist[0][j] = j
	}

	// 计算编辑距离
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			cost := 0
			if str1[i-1] != str2[j-1] {
				cost = 1
			}
			dist[i][j] = numberUtil.Min([]int{dist[i-1][j] + 1, dist[i][j-1] + 1, dist[i-1][j-1] + cost})
		}
	}

	return dist[len1][len2]
}

// Similar 计算两个字符串的相似度
// 只比较两个字符串字母、数字、汉字部分，其他符号去除
// 计算出两个字符串最大子串，除以最长的字符串，结果即为相似度
func Similar(str1 string, str2 string) float64 {
	distance := levenshteinDistance(filterSymbols(str1), filterSymbols(str2))
	maxLen := float64(len(str1))
	if len(str2) > len(str1) {
		maxLen = float64(len(str2))
	}
	similarity := 1.0 - (float64(distance) / maxLen)
	return similarity
}

// SimilarScale 计算两个字符串的相似度百分比
func SimilarScale(str1 string, str2 string, scale int) string {
	distance := levenshteinDistance(filterSymbols(str1), filterSymbols(str2))
	maxLen := float64(len(str1))
	if len(str2) > len(str1) {
		maxLen = float64(len(str2))
	}
	similarity := (1.0 - float64(distance)/maxLen) * 100.0
	return fmt.Sprintf("%.*f%%", scale, similarity)
}

// filterSymbols 过滤字符串，只保留字母、数字和汉字
func filterSymbols(str string) string {
	var builder strings.Builder
	for _, char := range str {
		if IsAlphaNumericOrChinese(char) {
			builder.WriteRune(char)
		}
	}
	return builder.String()
}

// IsAlphaNumericOrChinese  检查字符是否为字母、数字或汉字
func IsAlphaNumericOrChinese(char rune) bool {
	return (char >= '0' && char <= '9') || (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= 0x4e00 && char <= 0x9fff)
}
