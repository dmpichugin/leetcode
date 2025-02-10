package deps

func GetIntSlice(nums ...int) []int {

	return GetSlice(nums...)
}

func CopyIntSlice(s1 []int) []int {
	if s1 == nil {
		return nil
	}
	s2 := make([]int, 0, len(s1))
	if len(s1) == 0 {
		return s2
	}
	for _, v := range s1 {
		s2 = append(s2, v)
	}
	return s2
}

func GetStringSlice(s ...string) []string {

	return GetSlice(s...)
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

func GetDigits(num int) []int {

	if num == 0 {
		return []int{0}
	}

	v := num
	var digits []int
	var tmp int
	for v != 0 {
		tmp = v % 10
		digits = append(digits, tmp)
		v /= 10
	}

	return digits

}

func SumSlice(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}
