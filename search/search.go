package search

type search struct {
	searchInterface Search
}

func New() Search {
	return &search{}
}

func (s *search) BinarySearch(input []int, num int) int {
	start, end := 0, len(input)-1
	for start <= end {
		mid := (start + end) / 2

		if input[mid] == num {
			return mid
		}

		if input[mid] > num {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}

func (s *search) BinarySearchRecursive(input []int, start, end, num int) int {

	if start > end {
		return -1
	}

	mid := (start + end) / 2

	if input[mid] == num {
		return mid
	} else if input[mid] > num {
		return s.BinarySearchRecursive(input, start, mid-1, num)
	} else {
		return s.BinarySearchRecursive(input, mid+1, end, num)
	}
}

// Recurssive
func (s *search) FirstOccurance(input []int, start, end, num int) int {
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	if input[mid] > num {
		return s.FirstOccurance(input, start, mid-1, num)
	} else if input[mid] < num {
		return s.FirstOccurance(input, mid+1, end, num)
	} else {
		if mid == 0 || input[mid-1] != input[mid] {
			return mid
		} else {
			return s.FirstOccurance(input, start, mid-1, num)
		}
	}
}

// Iterative
func (s *search) LastOccurance(input []int, num int) int {
	start, end := 0, len(input)-1

	for start <= end {
		mid := (start + end) / 2

		if input[mid] > num {
			end = mid - 1
		} else if input[mid] < num {
			start = mid + 1
		} else {
			if mid == len(input)-1 || input[mid+1] != input[mid] {
				return mid
			}
			start = mid + 1
		}
	}
	return -1
}

func (s *search) CountOccurance(input []int, num int) int {
	firstIndex := s.FirstOccurance(input, 0, len(input)-1, num)
	if firstIndex == -1 {
		return -1
	}
	return s.LastOccurance(input, num) - firstIndex + 1
}

func (s *search) CountOnesInBinaryArray(input []int) int {
	start, end := 0, len(input)-1
	for start <= end {
		mid := (start + end) / 2
		if input[mid] != 1 {
			start = mid + 1
		} else {
			if mid == 0 || input[mid-1] == 0 {
				return len(input) - mid
			}
			end = mid - 1
		}
	}
	return -1
}

func (s *search) SquareRoot(num int) int {
	start, end, res := 0, num, -1
	for start <= end {
		mid := (start + end) / 2
		midsqrt := mid * mid
		if midsqrt == num {
			return mid
		} else if midsqrt > num {
			end = mid - 1
		} else {
			start = mid + 1
			res = mid
		}
	}

	return res
}

func (s *search) SearchInInfiniteArray(input []int, num int) int {
	if input[0] == num {
		return 0
	}
	i := 1
	for input[i] < num {
		i = i * 2
	}
	if input[i] == num {
		return i
	}
	return s.BinarySearchRecursive(input, (i/2)+1, i-1, num)
}

func (s *search) SearchInRotatedArray(input []int, num int) int {
	start, end := 0, len(input)-1
	for start <= end {
		mid := (start + end) / 2
		if input[mid] == num {
			return mid
		}
		if num <= input[mid] {
			if num < input[mid] && num >= input[start] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if num > input[mid] && num <= input[end-1] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}

func (s *search) FindPeakElement(input []int) int {
	start, end := 0, len(input)-1
	for start <= end {
		mid := (start + end) / 2
		if (mid == 0 || input[mid-1] <= input[mid]) && (mid == (len(input)-1) || input[mid+1] <= input[mid]) {
			return input[mid]
		} else {
			if mid > 0 && input[mid]-1 >= input[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}
	return -1
}
