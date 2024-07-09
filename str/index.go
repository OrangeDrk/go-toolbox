package strUtil

import "strings"

// IndexOf 返回字符在原始字符串的下标
func IndexOf(str string, subStr string) int {
	return strings.Index(str, subStr)
}

// IndexOfByRange 指定范围内查找指定字符
func IndexOfByRange(str string, subStr string, start int, end int) int {
	if end > len(str) {
		end = len(str)
	}
	index := strings.Index(str[start:end], subStr)
	if index > -1 {
		return start + index
	} else {
		return index
	}
}

// IndexOfByRangeStart 指定范围内查找指定字符
func IndexOfByRangeStart(str string, subStr string, start int) int {
	index := strings.Index(str[start:], subStr)
	if index > -1 {
		return start + index
	} else {
		return index
	}
}

// IndexOfIgnoreCase 返回字符在原始字符串的下标(大小写不敏感)
func IndexOfIgnoreCase(str string, subStr string) int {
	return IndexOf(strings.ToLower(str), strings.ToLower(subStr))
}

// IndexOfIgnoreCaseByRange 从指定下标开始，返回在字符串中的下标 (大小不敏感)
func IndexOfIgnoreCaseByRange(str string, subStr string, start int) int {
	return IndexOfByRangeStart(strings.ToLower(str), strings.ToLower(subStr), start)
}

// LastIndexOf 返回最后出现指定字符串的下标
func LastIndexOf(str string, subStr string) int {
	return strings.LastIndex(str, subStr)
}

// LastIndexOfIgnoreCase 返回最后出现指定字符串的下标（大小写不敏感）
func LastIndexOfIgnoreCase(str string, subStr string) int {
	return LastIndexOf(strings.ToLower(str), strings.ToLower(subStr))
}

// LastIndexOfByRangeStart 从指定下标开始，返回最后出现指定字符串的下标
func LastIndexOfByRangeStart(str string, subStr string, start int) int {
	index := strings.LastIndex(str[start:], subStr)
	if index > -1 {
		return start + index
	}
	return index
}

// LastIndexOfIgnoreCaseByRangeStart 从指定下标开始，返回最后出现指定字符串的下标(大小写不敏感)
func LastIndexOfIgnoreCaseByRangeStart(str string, subStr string, start int) int {
	index := LastIndexOfIgnoreCase(str[start:], subStr)
	if index > -1 {
		return start + index
	}
	return index
}

// OrdinalIndexOf 返回字符串 searchStr 在字符串 str 中第 ordinal 次出现的位置。
// 如果 str="" 或 searchStr=" 或 ordinal≥0 则返回-1
func OrdinalIndexOf(str string, subStr string, ordinal int) int {
	if str == "" || subStr == "" || ordinal <= 0 {
		return -1
	}

	index := -1
	count := 0

	for {
		idx := strings.Index(str[index+1:], subStr)
		if idx == -1 {
			break
		}
		index += idx + 1
		count++
		if count == ordinal {
			return index
		}
	}
	return -1
}
