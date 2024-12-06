package arrays

type ArrayProblems interface {
	LargestElementOfAnArray(input []int) int
	SecondLargest(input []int) int
	ArraySorted(input []int) bool
	ReverseArray(input []int) []int
	RemoveDuplicatesFromSortedArray(input []int) []int
	MoveZeros(input []int) []int
	ShiftArrayByDElements(input []int, d int) []int
	LeaderInArray(input []int) []int
	MaximumDifference(input []int) int
	StockBuyAndSell(input []int) int
	RemoveDuplicates(nums []int) int
	TrapRainWater(input []int) int
	ConsecutiveOnes(input []int) int
	MaxSubArray(input []int) int
	MaxlengthEvenOddSubArray(input []int) int
	MaxCircularSumSubArray(input []int) int
	MajorityElement(input []int) int
	MinConsecutiveFlips(input []int) int
	MaxKSum(input []int, k int) int
	SubArrayWithGivenSum(input []int, sum int) bool
	PrefixSumProblem(input []int)
	EquilibriumElementInArray(input []int) bool
	FindDuplicateInArray(input []int)
	FindTwoSumSortedArray(input []int, num int) bool
}
