package Sort

import "fmt"

func SortRunner() {
	s := New()
	fmt.Println(s.MergeTwoSortedArrays([]int{10, 20, 50}, []int{5, 50, 70}))
}
