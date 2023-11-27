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
	fmt.Println(searchInst.SearchInRotatedArray([]int{10, 20, 40, 60, 5, 8}, 5))
	fmt.Println(searchInst.FindPeakElement([]int{80, 70, 90}))
	fmt.Println(searchInst.MedianOfTwoSortedArrays([]int{1, 3, 5, 7}, []int{2, 4, 6, 8, 9}, 4, 5))
	fmt.Println(searchInst.RepeatingElements([]int{0, 1, 2, 3, 5, 4, 6, 2}))
}
