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
