package binarysearch

// Time O(N)
// Space O(1)
func searchSlow(nums []int, target int) int {
    for i,n := range nums{
        if n == target{
            return i
        }
    }
    
    return -1
}

// Time O(logN)
// Space O(1)
func search(nums []int, target int) int {
    pt, left, right:=0, 0, len(nums)-1
    
    for left <= right{
        pt = left + (right-left)/2    
        if target > nums[pt]{
            left = pt+1
        } else if target < nums[pt]{
            right = pt-1
        } else {
            return pt
        }
        
    }
    
    return -1
}