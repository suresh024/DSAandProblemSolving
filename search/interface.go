package search

type Search interface {
	BinarySearch(input []int, num int) int
	BinarySearchRecursive(input []int, start, end, num int) int
	FirstOccurance(input []int, start, end, num int) int
	LastOccurance(input []int, num int) int
	CountOccurance(input []int, num int) int
	CountOnesInBinaryArray(input []int) int
	SquareRoot(num int) int
	SearchInInfiniteArray(input []int, num int) int
	SearchInRotatedArray(input []int, num int) int
	FindPeakElement(input []int) int
	MedianOfTwoSortedArrays(a1, a2 []int, n1, n2 int) int
	RepeatingElements(input []int) int
	MinPages(books []int, n, k int) int
}
