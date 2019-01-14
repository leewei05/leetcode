package reverse

import (
	"strconv"
)

func reverse(x int) int {
	if x < 0 {
		x = -x
		b := []byte(strconv.Itoa(x))
		n := make([]byte, len(b))
		for i := range b {
			n[i] = b[len(b)-i-1]
		}
		s := string(n)
		r, _ := strconv.Atoi(s)
		if r > 2147483647 || r < -2147483647 {
			r = 0
		}
		return r * -1
	}
	b := []byte(strconv.Itoa(x))
	n := make([]byte, len(b))
	for i := range b {
		n[i] = b[len(b)-i-1]
	}
	s := string(n)
	r, _ := strconv.Atoi(s)
	if r > 2147483647 || r < -2147483647 {
		r = 0
	}
	return r
}
