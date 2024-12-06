package utils

import (
	"fmt"
	"runtime"
	"strings"
)

func ReverseArrayHelper(input []int, start, end int) []int {
	for start < end {
		input[start], input[end] = input[end], input[start]
		start += 1
		end -= 1
	}
	return input
}

func Max(num1, num2 int) int {
	switch num1 > num2 {
	case true:
		return num1
	case false:
		return num2
	default:
		return num1
	}
}
func Min(num1, num2 int) int {
	switch num1 < num2 {
	case true:
		return num1
	case false:
		return num2
	default:
		return num1
	}
}

func LogFunctionInfo(funcName string, input interface{}, output interface{}) {
	fmt.Printf("%s\ninput - %v\noutput - %v\n\n", funcName, input, output)
}

func GetFunctionName() string {
	pc, _, _, _ := runtime.Caller(1)
	funcInfo := runtime.FuncForPC(pc)
	parts := strings.Split(funcInfo.Name(), ".")
	return parts[len(parts)-1]
}

func PrefixSumArray(input []int) []int {
	pSum := make([]int, len(input))
	pSum[0] = input[0]
	for i := 1; i < len(input); i++ {
		pSum[i] = input[i] + pSum[i-1]
	}
	return pSum
}

func SuffixSumArray(input []int) []int {
	length := len(input)
	sSum := make([]int, length)
	sSum[length-1] = input[length-1]
	for i := length - 2; i > 0; i-- {
		sSum[i] = sSum[i-1] + input[i]
	}

	return sSum
}
