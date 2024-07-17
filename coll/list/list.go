package listUtil

import (
	"sort"
)

// Of 快速创建一个 List
func Of[T any](elems ...T) []T {
	return elems
}

// Append 向列表中添加元素
func Append[T any](l *[]T, item ...T) {
	*l = append(*l, item...)
}

// IsEmpty 判断是否为空
func IsEmpty[T any](list []T) bool {
	if len(list) > 0 {
		return false
	} else {
		return true
	}
}

//// Length 获取长度
//func Length[T any](list []T) int {
//	return len(list)
//}

// Page 分页
// pageNo - 页码，从1开始计算
// pageSize - 每页条数
func Page[T any](list []T, pageNo int, pageSize int) []T {
	if pageNo < 1 || pageSize < 1 {
		return []T{}
	}

	startIndex := (pageNo - 1) * pageSize
	if startIndex >= len(list) {
		return []T{}
	}

	endIndex := startIndex + pageSize
	if endIndex > len(list) {
		endIndex = len(list)
	}

	return list[startIndex:endIndex]
}

// Sort 排序
// list 是一个指向元素类型为 T 的切片的指针
// less 是一个比较函数，定义了 T 类型元素的排序顺序
func Sort[T any](list []T, less func(i, j T) bool) {
	sort.Slice(list, func(i, j int) bool {
		return less((list)[i], (list)[j])
	})
}

// Reverse 翻转切片
func Reverse[T any](list []T) {
	n := len(list)
	for i := 0; i < n/2; i++ {
		list[i], list[n-1-i] = list[n-1-i], list[i]
	}
}

// ReverseNew 翻转切片,并返回新数组
func ReverseNew[T any](list []T) []T {
	destination := make([]T, len(list))
	copy(destination, list)
	Reverse(destination)
	return destination
}

// SetOrAppend 设置或增加元素；当index小于List的长度时，替换指定位置的值，否则在尾部追加
func SetOrAppend[T any](list *[]T, index int, elem T) {
	if index < len(*list) {
		(*list)[index] = elem
	} else {
		Append(list, elem)
	}
}

// LastIndexOf 获取匹配规则定义中匹配到元素的最后位置
func LastIndexOf[T any](list []T, matcher func(T) bool) int {
	lastIndex := -1
	for i, v := range list {
		if matcher(v) {
			lastIndex = i
		}
	}
	return lastIndex
}

// IndexOfAll 获取匹配规则定义中匹配到元素的所有位置
func IndexOfAll[T any](list []T, matcher func(T) bool) []int {
	var indexL []int
	for i, v := range list {
		if matcher(v) {
			Append(&indexL, i)
		}
	}
	return indexL
}

// Split 对集合按照指定长度分段，每一个段为单独的集合，返回这个集合的列表
func Split[T any](list []T, size int) [][]T {
	if size <= 0 {
		return nil
	}

	var result [][]T
	for i := 0; i < len(list); i += size {
		end := i + size
		if end > len(list) {
			end = len(list)
		}
		result = append(result, list[i:end])
	}
	return result
}

// SplitAvg 将集合平均分成多个list，返回这个集合的列表
// limit - 要平均分成几段
func SplitAvg[T any](list []T, limit int) [][]T {
	if limit <= 0 || len(list) == 0 {
		return nil
	}

	var result [][]T
	n := len(list)
	chunkSize := (n + limit - 1) / limit // 计算每段的平均长度，向上取整

	for i := 0; i < n; i += chunkSize {
		end := i + chunkSize
		if end > n {
			end = n
		}
		result = append(result, list[i:end])
	}
	return result
}
