package solution

func merge(nums1 []int, m int, nums2 []int, n int) []int {

	i1 := m - 1
	i2 := n - 1
	index := m + n - 1

	for index >= 0 {

		if i1 < 0 {
			nums1[index] = nums2[i2]
			i2--
			index--
			continue
		}

		if i2 < 0 {
			nums1[index] = nums1[i1]
			i1--
			index--
			continue
		}

		v1 := nums1[i1]
		v2 := nums2[i2]
		if v1 > v2 {
			nums1[index] = v1
			i1--
		} else {
			nums1[index] = v2
			i2--
		}
		index--
	}

	return nums1
}
