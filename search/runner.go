package search

import "fmt"

func SearchRunner() {
	searchInst := New()

	fmt.Println(searchInst.BinarySearch([]int{1, 2, 3, 50, 78, 100, 108}, 4))
	fmt.Println(searchInst.BinarySearchRecursive([]int{1, 2, 3, 50, 78, 100, 108}, 0, 6, 50))
	fmt.Println(searchInst.FirstOccurance([]int{5, 10, 10, 20, 20, 40}, 0, 5, 10))
	fmt.Println(searchInst.LastOccurance([]int{5, 10, 10, 20, 20, 40}, 5))
	fmt.Println(searchInst.CountOccurance([]int{5, 10, 10, 20, 20, 20}, 5))
	fmt.Println(searchInst.CountOnesInBinaryArray([]int{0, 0, 1, 1, 1}))
	fmt.Println(searchInst.SquareRoot(0))
	fmt.Println(searchInst.SearchInInfiniteArray([]int{1, 10, 15, 20, 40, 60, 80, 100, 200, 500}, 100))
}
