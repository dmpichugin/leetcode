package solution

func removeElement(nums []int, val int) int {

	if len(nums) == 0 {
		return 0
	}

	s := len(nums)

	if len(nums) == 1 {
		if nums[0] == val {
			return 0
		} else {
			return 1
		}
	}

	valCount := 0
	currentIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			valCount++
		} else {
			nums[currentIndex] = nums[i]
			currentIndex++
		}
	}

	notEqual := s - valCount

	return notEqual
}
