package strUtil

import (
	"fmt"
	numberUtil "github.com/OrangeDrk/go-toolbox/number"
	"strings"
)

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

// Revers 反转字符串
func Revers(str string) string {
	// 将字符串转换为 rune 切片，以便处理多字节字符
	runes := []rune(str)
	n := len(runes)

	// 反转 rune 切片
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	// 将反转后的 rune 切片转换回字符串
	return string(runes)
}

// ------------------------- test

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
