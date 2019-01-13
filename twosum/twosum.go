package twosum

func twoSum(nums []int, target int) []int {
	r := []int{}
	for i, num := range nums {
		ntarget := target - num
		firstInt := i
		for i := 0; i < len(nums); i++ {
			if nums[i] == ntarget {
				if i != firstInt {
					r = append(r, firstInt, i)
					return r
				}
			}
		}
	}
	return nil
}
