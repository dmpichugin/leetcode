package deps

func GetIntSlice(nums ...int) []int {
	return GetSlice(nums...)
}

func GetSlice[T any](nums ...T) []T {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	res := make([]T, 0, len(nums))

	for _, num := range nums {
		res = append(res, num)
	}

	return res
}
