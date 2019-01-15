package palindrome

func isPalindrome(x int) bool {
	r := false
	n := 0
	nx := x
	if x < 0 {
		r = false
	}
	for nx > 0 {
		p := nx % 10
		n = n*10 + p
		nx = nx / 10
	}
	if x == n {
		r = true
	}
	return r
}
