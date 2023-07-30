package slice

func Remove[T any](sl []T, index int) []T {
	l := len(sl)
	c := cap(sl)

	if index > l || l == 0 {
		return sl
	}

	if l-1 <= c>>2 {
		// 1/4 时缩容一半
		result := make([]T, l-1, (c+1)/2)
		// 删除第一个元素没有前半段
		if index != 0 {
			copy(result, sl[:index])
		}
		// 删除最后一个元素没有后半段
		if l > index {
			copy(result[index:], sl[index+1:])
		}
		return result
	} else {
		copy(sl[index:], sl[index+1:])
		return sl[:l-1]
	}
}
