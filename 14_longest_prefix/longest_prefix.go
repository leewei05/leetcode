package longestCommonPrefix

func longestCommonPrefix(strs []string) string {
	curr := 0
	if len(strs) == 0 {
		return ""
	}

	first := strs[0]
	if first == "" {
		return ""
	}

	for i := 0; i < len(first); i++ {
		for _, s := range strs {
			if len(s) < i+1 {
				return strs[0][:curr]
			}
			if first[i] != s[i] {
				return strs[0][:curr]
			}
		}
		curr += 1
	}
	return strs[0][:curr]
}
