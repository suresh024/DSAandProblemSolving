package arrays

func ArrayRunner() {
	arrayObj := New()

	inputArray := []int{0, 1, 4, 0, 0, 4, 2, 2, 0, 5, 6}

	arrayObj.LargestElementOfAnArray(inputArray)
	arrayObj.SecondLargest(inputArray)
	arrayObj.ArraySorted(inputArray)
	arrayObj.ReverseArray(inputArray)
	arrayObj.MoveZeros(inputArray)
	arrayObj.RemoveDuplicatesFromSortedArray(inputArray)
	arrayObj.StockBuyAndSell(inputArray)
	arrayObj.ShiftArrayByDElements(inputArray, 2)
	arrayObj.MaximumDifference(inputArray)
	arrayObj.LeaderInArray(inputArray)
	arrayObj.ConsecutiveOnes([]int{1, 1, 0, 1})
	arrayObj.MaxSubArray([]int{-3, 8, -2, 4, -5, 6})
	arrayObj.MaxlengthEvenOddSubArray([]int{5, 10, 20, 6, 3, 8})
	arrayObj.MaxCircularSumSubArray([]int{8, -4, 3, -5, 4})
	arrayObj.RemoveDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	arrayObj.TrapRainWater([]int{5, 0, 6, 2, 3})
}
