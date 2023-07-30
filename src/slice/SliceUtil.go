package slice

func DeleteSliceElement[T any](sl []T, index int) []T {
	l := len(sl)
	c := cap(sl)

	if index > l || l == 0 {
		return sl
	}

	if l-1 <= c>>2 {
		// 1/4 时缩容一半
		result := make([]T, l-1, (c+1)/2)
		if index != 0 {
			copy(result, sl[:index])
		}
		// 如果删除的是最后一个元素就不用拷贝另外一部分
		if l > index {
			copy(result[index:], sl[index+1:])
		}
		return result
	} else {
		copy(sl[index:], sl[index+1:])
		return sl[:l-1]
	}
}
