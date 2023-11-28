package leetcode

import "fmt"

func LeetCodeRunner() {
	lc := New()
	fmt.Println(lc.knightDialer(1))
	lc.Merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}
