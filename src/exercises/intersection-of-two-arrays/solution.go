package solution

func intersection(nums1 []int, nums2 []int) []int {

	res := make([]int, 0, len(nums1))
	m := make(map[int]struct{}, len(nums1))
	m2 := make(map[int]struct{}, len(nums2))

	for _, v := range nums1 {
		m[v] = struct{}{}
	}
	var ok bool
	for _, v := range nums2 {
		if _, ok = m2[v]; ok {
			continue
		}
		_, ok = m[v]
		if ok {
			res = append(res, v)
			m2[v] = struct{}{}
		}
	}
	if len(res) == 0 {
		return nil
	}
	return res
}
