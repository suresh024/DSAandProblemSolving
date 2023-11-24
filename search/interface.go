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
}
