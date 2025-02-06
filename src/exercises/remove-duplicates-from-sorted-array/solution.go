package solution

func removeDuplicates(nums []int) int {

	if len(nums) <= 1 {
		return len(nums)
	}

	var uniqCount = 1
	var prevValue = nums[0]
	var currentIndex = 1

	for i := 1; i < len(nums); i++ {
		if prevValue != nums[i] {
			uniqCount++
			prevValue = nums[i]
			nums[currentIndex] = prevValue
			currentIndex++
		}
	}

	return uniqCount
}
