package validParenteses

// 123,125 | 40,41 | 91,93
func isValid(s string) bool {
	// len odd numbers
	if len(s)%2 != 0 {
		return false
	}

	for i := 0; i < len(s); i++ {
		var target uint8
		switch s[i] {
		case 123:
			target = 125
		case 40:
			target = 41
		case 91:
			target = 93
		}
		// len 2
		if len(s) == 2 {
			if s[i+1] == target {
				return true
			}
			return false
		}

		if i+1 < len(s) && s[i+1] == target {
			continue
		}

		flag := false
		for j := i + 1; j < len(s)-1; j = j + 2 {
			if target == s[j] {
				flag = true
			}
			if j == len(s)-1 && flag == false {
				return false
			}
		}
	}
	return true
}
