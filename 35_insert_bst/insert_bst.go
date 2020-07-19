package insertbst

// Time O(logN)
// Space O(1)
func searchInsert(nums []int, target int) int {
	lo, mid, hi := 0, 0, len(nums)

	if target == 0 {
		return 0
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	}

	for lo < hi {
		mid = lo + (hi-lo)/2

		if target > nums[mid] {
			lo = lo + 1
		} else {
			hi = mid
		}
	}

	return lo
}
