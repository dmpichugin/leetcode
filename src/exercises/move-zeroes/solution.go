package solution

func moveZeroes(nums []int) {

	index := 0
	zeroesCount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[index] = nums[i]
			index++
			if zeroesCount > 0 {
				nums[i] = 0
			}
			continue
		} else {
			zeroesCount++
		}
	}
}
