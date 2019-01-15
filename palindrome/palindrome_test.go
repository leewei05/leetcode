package palindrome

import "testing"

func Test_palindrome(t *testing.T) {
	if i := isPalindrome(121); i != true {
		t.Error("Expect:true, but the result is:", i)
	} else {
		t.Log("Pass!")
	}
}
func Test_palindrome_1(t *testing.T) {
	if i := isPalindrome(10); i != false {
		t.Error("Expect:false, but the result is:", i)
	} else {
		t.Log("Pass!")
	}
}
func Test_palindrome_2(t *testing.T) {
	if i := isPalindrome(-121); i != false {
		t.Error("Expect:false, but the result is:", i)
	} else {
		t.Log("Pass!")
	}
}
