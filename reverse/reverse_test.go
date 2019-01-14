package reverse

import "testing"

func Test_Reverse(t *testing.T) {
	if i := reverse(123); i != 321 {
		t.Error("Expect:321, but the result is:", i)
	} else {
		t.Log("Pass!")
	}
}

func Test_Reverse_1(t *testing.T) {
	if i := reverse(-123); i != -321 {
		t.Error("Expect:-321, but the result is:", i)
	} else {
		t.Log("Pass!")
	}
}

func Test_Reverse_2(t *testing.T) {
	if i := reverse(120); i != 21 {
		t.Error("Expect:21, but the result is:", i)
	} else {
		t.Log("Pass!")
	}
}

func Test_Reverse_3(t *testing.T) {
	if i := reverse(1534236469); i != 0 {
		t.Error("Expect:0, but the result is:", i)
	} else {
		t.Log("Pass!")
	}
}
