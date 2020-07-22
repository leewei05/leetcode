package intersection

// Time O(m+n) m nums1, n nums2
// Space O(m+n)
func intersection(nums1 []int, nums2 []int) []int {
    arrMap := make(map[int]int)
    res := []int{}
    
    for _,n := range nums1 {
        if _, ok := arrMap[n]; !ok{
            arrMap[n]++
        }
    }
    
    for _,n := range nums2 {
        if _, ok := arrMap[n]; ok{
            res = append(res, n)
            delete(arrMap, n)
        }
    }
    
    return res
}