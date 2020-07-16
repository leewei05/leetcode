package container

import "math"

// Time O(N*2)
// Space O(1)
func maxArea(height []int) int {
	max := 0

	for i, h := range height {
		for j, h2 := range height {
			if i == j {
				continue
			}
			w := int(math.Abs(float64(i - j)))
			calH := maxHeight(h, h2)

			if calH*w > max {
				max = calH * w
			}
		}
	}

	return max
}

func maxHeight(h1, h2 int) int {
	if h1 > h2 {
		return h2
	} else if h1 < h2 {
		return h1
	} else {
		return h1
	}
}
