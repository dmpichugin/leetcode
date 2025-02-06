package solution

import (
	"errors"
)

func searchInsert(nums []int, target int) int {

	if nums == nil || len(nums) == 0 {
		return 0
	}

	if target <= nums[0] {
		return 0
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	}

	index, err := BSearch(nums, target)
	if err == nil {
		return index
	}

	return index + 1
}

var (
	NotFoundErr   = errors.New("not found")
	OutOfRangeErr = errors.New("value out of range")
)

// BSearch finds the index in the sorted array using split array in half
func BSearch(a []int, val int) (int, error) {
	if len(a) == 0 || val < a[0] || val > a[len(a)-1] {
		return -1, OutOfRangeErr
	}
	l := 0
	r := len(a) - 1

	var mid int

	for {
		mid = (l + r) / 2
		if l > r {
			break
		}
		if a[mid] == val {
			return mid, nil
		}
		if val < a[mid] {
			r = mid - 1
		} else if val > a[mid] {
			l = mid + 1
		}
	}

	return mid, NotFoundErr
}
