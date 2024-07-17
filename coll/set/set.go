package setUtil

// 定义Set类型，底层使用map
type Set[T comparable] map[T]struct{}

// NewSet 创建并返回一个新的 Set
func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

// Add 向 Set 中添加元素
func (s Set[T]) Add(elem T) {
	s[elem] = struct{}{}
}

// Remove 从 Set 中删除元素
func (s Set[T]) Remove(elem T) {
	delete(s, elem)
}

// Contains 检查 Set 中是否包含指定元素
func (s Set[T]) Contains(elem T) bool {
	_, exists := s[elem]
	return exists
}

// Size 返回 Set 的大小
func (s Set[T]) Size() int {
	return len(s)
}

// Clear 清空 Set
func (s Set[T]) Clear() {
	for k := range s {
		delete(s, k)
	}
}

// ToSlice 将 Set 转换为切片
func (s Set[T]) ToSlice() []T {
	keys := make([]T, 0, len(s))
	for key := range s {
		keys = append(keys, key)
	}
	return keys
}
