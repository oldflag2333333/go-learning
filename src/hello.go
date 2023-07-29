package main

import "fmt"
import "github.com/oldflag/awesome/src/slice"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	index := 2 // 要删除的元素的索引

	numbers = slice.DeleteSliceElement(numbers, index)
	fmt.Println(numbers) // 输出: [1 2 4 5]
}
