package strUtil

// 返回字符在原始字符串的下标
func IndexOf(sourceStr string, subStr rune) int {
	if len(sourceStr) <= 0 {
		return -1
	}

	strRune := []rune(sourceStr)

	for i, v := range strRune {
		if v == subStr {
			return i
		}
	}
	return -1
}

// 切割原始字符，返回子串 范围：[s,e)
func SubString(sourceStr string, s int, e int) string {
	if len(sourceStr) <= 0 {
		return ""
	}
	strRune := []rune(sourceStr)
	if s < 0 {
		s = 0
	}
	if e > len(strRune) {
		e = len(strRune)
	}
	if s >= len(strRune) || e < 0 {
		return ""
	}
	return string(strRune[s:e])
}
