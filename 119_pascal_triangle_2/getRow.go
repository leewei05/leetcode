// Time O
func getRow(rowIndex int) []int {
	pascal := make([]int, rowIndex+1)

	for i := 0; i < rowIndex+1; i++ {
		if i == 0 {
			pascal[i] = 1
		}
		if i == 1 {
			pascal[0] = 1
			pascal[1] = 1
		}

		lastLine := make([]int, len(pascal))
		copy(lastLine, pascal)

		if i > 1 {
			for j := 1; j < i+1; j++ {
				if j == i {
					pascal[j] = 1
					continue
				}
				pascal[j] = lastLine[j-1] + lastLine[j]
			}
		}
	}

	return pascal
}
