package main

import "fmt"
import "github.com/oldflag/awesome/src/slice"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	for i := 0; i < 12; i++ {
		numbers = slice.DeleteSliceElement(numbers, 0)
		fmt.Println(numbers, len(numbers), cap(numbers))
	}

	numbers = slice.DeleteSliceElement(numbers, 0)
	fmt.Println(numbers, len(numbers), cap(numbers))
}
