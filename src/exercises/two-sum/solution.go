package solution

func twoSum(nums []int, target int) []int {

	m := make(map[int]int)

	var res []int
	for k, v := range nums {

		diffKey := target - v
		diffValue, ok := m[diffKey]
		if ok {
			res = append(res, diffValue)
			res = append(res, k)
			break
		} else {
			m[v] = k
		}
	}

	return res
}
