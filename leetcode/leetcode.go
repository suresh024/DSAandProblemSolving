package leetcode

import "fmt"

const MOD = int(1e9 + 7)

type leetcode struct {
	LeetcodeInterface Leetcode
}

func New() Leetcode {
	return &leetcode{}
}

func (lc *leetcode) knightDialer(n int) int {
	countsPrev := make([]int, 10, 10)
	countsNext := make([]int, 10, 10)
	for i := range countsPrev {
		countsPrev[i] = 1
	}

	for i := 1; i < n; i++ {
		countsNext[0] = (countsPrev[4] + countsPrev[6]) % MOD
		countsNext[1] = (countsPrev[8] + countsPrev[6]) % MOD
		countsNext[2] = (countsPrev[7] + countsPrev[9]) % MOD
		countsNext[3] = (countsPrev[4] + countsPrev[8]) % MOD
		countsNext[4] = (countsPrev[3] + countsPrev[9]) + countsPrev[0]%MOD
		countsNext[5] = 0
		countsNext[6] = (countsPrev[1] + countsPrev[7]) + countsPrev[0]%MOD
		countsNext[7] = (countsPrev[2] + countsPrev[6]) % MOD
		countsNext[8] = (countsPrev[1] + countsPrev[3]) % MOD
		countsNext[9] = (countsPrev[4] + countsPrev[2]) % MOD

		copy(countsPrev, countsNext)
	}

	sum := 0
	for i := range countsPrev {
		sum = (sum + countsPrev[i]) % MOD
	}

	return sum
}

func (lc *leetcode) Merge(nums1 []int, m int, nums2 []int, n int) {
	for n != 0 {
		if m != 0 && nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
	fmt.Println(nums1)
}
