package arrays

import (
	"github.com/suresh024/DSAandProblemSolving/utils"
)

type array struct {
	problem ArrayProblems
}

func New() ArrayProblems {
	return &array{}
}

func (a *array) LargestElementOfAnArray(input []int) int {
	largest := 0
	for i := 0; i < len(input); i++ {
		if input[i] > input[largest] {
			largest = i
		}
	}

	//logging function details
	utils.LogFunctionInfo(utils.GetFunctionName(), input, input[largest])

	return largest
}

func (a *array) SecondLargest(input []int) int {
	res, largest := -1, 0

	for i := 1; i < len(input); i++ {
		if input[i] > input[largest] {
			res = largest
			largest = i
		} else if input[i] < input[largest] {
			if res == -1 || input[i] > input[res] {
				res = i
			}
		}
	}

	//logging function details
	utils.LogFunctionInfo(utils.GetFunctionName(), input, input[res])

	return res
}
func (a *array) ArraySorted(input []int) bool {

	res := false

	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			res = true
			return res
		}
	}

	//logging function details
	utils.LogFunctionInfo(utils.GetFunctionName(), input, res)

	return res
}

func (a *array) ReverseArray(input []int) []int {
	start, last := 0, len(input)-1

	for start < last {
		input[start], input[last] = input[last], input[start]
		start += 1
		last -= 1
	}

	//logging function details
	utils.LogFunctionInfo(utils.GetFunctionName(), input, input)

	return input
}

func (a *array) RemoveDuplicatesFromSortedArray(input []int) []int {
	//TIME - O(n)
	//SPACE - O(1)
	res := 1

	for i := 1; i < len(input); i++ {
		if input[i] != input[res-1] {
			input[res] = input[i]
			res += 1
		}
	}

	//logging function details
	utils.LogFunctionInfo(utils.GetFunctionName(), input, res)

	return input
}

func (a *array) MoveZeros(input []int) []int {
	res := 0

	for i := 0; i < len(input); i++ {
		if input[i] != 0 {
			input[i], input[res] = input[res], input[i]
			res += 1
		}
	}

	//logging function details
	utils.LogFunctionInfo(utils.GetFunctionName(), input, input)

	return input
}

func (a *array) ShiftArrayByDElements(input []int, d int) []int {

	input = utils.ReverseArrayHelper(input, 0, d-1)
	input = utils.ReverseArrayHelper(input, d, len(input)-1)
	input = utils.ReverseArrayHelper(input, 0, len(input)-1)

	//logging function details
	utils.LogFunctionInfo(utils.GetFunctionName(), input, input)

	return input
}

func (a *array) LeaderInArray(input []int) []int {
	out := []int{input[len(input)-1]}
	res := input[len(input)-1]
	for i := len(input) - 2; i >= 0; i-- {
		if input[i] > res {
			res = input[i]
			out = append(out, input[i])
		}
	}

	return a.ReverseArray(out)
}

// 2 4 1 0 1 0 0 0 0 6 5
func (a *array) MaximumDifference(input []int) int {
	res := input[1] - input[0]
	min := input[0]

	for i := 1; i < len(input); i++ {
		res = utils.Max(input[i]-min, res)
		min = utils.Min(input[i], min)
	}

	//logging function details
	utils.LogFunctionInfo(utils.GetFunctionName(), input, res)

	return res
}

func (a *array) StockBuyAndSell(input []int) int {
	profit := 0

	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			profit += input[i] - input[i-1]
		}
	}

	//logging function details
	utils.LogFunctionInfo(utils.GetFunctionName(), input, profit)

	return profit
}

func (a *array) RemoveDuplicates(nums []int) int {
	res := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[res-1] {
			nums[res] = nums[i]
			res += 1
		}
	}

	//logging function details
	utils.LogFunctionInfo(utils.GetFunctionName(), nums, res)

	return len(nums)
}

// [5, 0, 6, 2, 3]
func (a *array) TrapRainWater(input []int) int {
	res := 0

	lmax, rmax := make([]int, len(input)), make([]int, len(input))

	lmax[0], rmax[len(input)-1] = input[0], input[len(input)-1]
	//[5,5,6,6,6]
	for i := 1; i < len(input); i++ {
		lmax[i] = utils.Max(input[i], lmax[i-1])
	}

	//[6, 6, 6, 3, 3]
	for i := len(input) - 2; i >= 0; i-- {
		rmax[i] = utils.Max(input[i], rmax[i+1])
	}

	for i := 0; i < len(input); i++ {
		res += utils.Min(lmax[i], rmax[i]) - input[i]
	}

	//logging function details
	utils.LogFunctionInfo(utils.GetFunctionName(), input, res)

	return res

}

func (a *array) ConsecutiveOnes(input []int) int {
	res, current := input[0], input[0]

	for i := 1; i < len(input); i++ {
		if input[i] == 0 {
			current = 0
		} else {
			current += 1
			res = utils.Max(res, current)
		}
	}

	//logging function details
	utils.LogFunctionInfo(utils.GetFunctionName(), input, res)

	return res
}

// []int{-3, 8, -2, 4, -5, 6}
func (a *array) MaxSubArray(input []int) int {
	res, current := input[0], input[0]

	for i := 1; i < len(input); i++ {
		current = utils.Max(current+input[i], input[i])
		res = utils.Max(current, res)
	}

	//logging function details
	utils.LogFunctionInfo(utils.GetFunctionName(), input, res)

	return res
}

func (a *array) MaxlengthEvenOddSubArray(input []int) int {
	res, current := 1, 1

	for i := 1; i < len(input); i++ {
		if (input[i]%2 == 0 && input[i-1]%2 != 0) || (input[i-1]%2 == 0 && input[i]%2 != 0) {
			current += 1
			res = utils.Max(current, res)
		} else {
			current = 1
		}
	}

	//logging function details
	utils.LogFunctionInfo(utils.GetFunctionName(), input, res)

	return res
}

// []int{8, -4, 3, -5, 4}
func (a *array) MaxCircularSumSubArray(input []int) int {
	res, curr := 0, 0
	min, currMin := 0, 0

	for i := 1; i < len(input); i++ {
		curr = utils.Max(curr+input[i], input[i])
		res = utils.Max(curr, res)

		currMin = utils.Min(currMin+input[i], input[i])
		min = utils.Min(currMin, min)
	}

	//logging function details
	utils.LogFunctionInfo(utils.GetFunctionName(), input, res)

	return res - min
}
