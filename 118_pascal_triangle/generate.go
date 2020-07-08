// Time O(N^2)
// Space O(N^2)
func generate(numRows int) [][]int {
	pascal := make([][]int, numRows)
	if numRows == 0 {
		return pascal
	}

	for i := 0; i < numRows; i++ {
		if i == 0 {
			pascal[i] = append(pascal[i], 1)
		}
		if i == 1 {
			pascal[i] = append(pascal[i], 1)
			pascal[i] = append(pascal[i], 1)
		}

		if i > 1 {
			for j := 0; j < i+1; j++ {

				if j == 0 {
					pascal[i] = append(pascal[i], 1)
					continue
				}
				if j == i {
					pascal[i] = append(pascal[i], 1)
					continue
				}

				pascal[i] = append(pascal[i], (pascal[i-1][j-1] + pascal[i-1][j]))
			}
		}

	}

	return pascal
}
