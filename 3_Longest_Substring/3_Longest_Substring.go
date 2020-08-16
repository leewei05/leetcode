func lengthOfLongestSubstring(s string) int {
    runes := []rune(s)
    if len(runes) == 0 {
		return 0
	}
    
	var r [128]uint8
	max, left, right := 0, 0, 0
	for left < len(runes) {
		if right < len(s) && r[runes[right]] == 0 {
            fmt.Println(runes[right])
			r[runes[right]] = 1
			right++
		} else {
			r[runes[left]] = 0
			left++
		}
        if max < right -left{
            max = right -left
        }
	}
	return max
}

