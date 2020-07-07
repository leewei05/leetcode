// Time Complexity O(N^2)
// Space Complexity O(1)
func maxProfit(prices []int) int {
	profit := 0
	for i, p := range prices {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-p > profit && prices[j]-p > 0 && prices[j] > p {
				profit = prices[j] - p
			}
		}
	}
	return profit
}

// Time O(N)
// Space O(1)
func maxProfit(prices []int) int {
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}

	profit, min := 0, 9223372036854775807
	for _, p := range prices {
		if p < min {
			min = p
		} else if p-min > profit {
			profit = p - min
		}
	}

	return profit
}


