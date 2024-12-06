package Sort

type sorting struct {
	sortingInterface Sort
}

func New() Sort {
	return &sorting{}
}

func (s *sorting) MergeTwoSortedArrays(a, b []int) []int {
	out := make([]int, 0)
	m, n, i, j := len(a), len(b), 0, 0
	for i < m && j < n {
		if a[i] <= b[j] {
			out = append(out, a[i])
			i += 1
		} else {
			out = append(out, b[j])
			j += 1
		}
	}

	for i < m {
		out = append(out, a[i])
		i += 1
	}
	for j < n {
		out = append(out, b[j])
		j += 1
	}
	return out
}
func (s *sorting) Merge(a []int, low, mid, high int) []int {
	n1, n2 := mid-low+1, high-mid
	left, right := make([]int, n1), make([]int, n2)
	for i := 0; i < n1; i++ {
		left[i] = a[low+i]
	}
	for j := 0; j < n2; j++ {
		right[j] = a[mid+j+1]
	}
	i, j, k := 0, 0, low

	for i < n1 && j < n2 {
		if left[i] <= right[i] {
			a[k] = left[i]
			i++
			k++
		} else {
			a[k] = right[j]
			j++
			k++
		}
	}
	//TODO: need to add remaining code
}
