package validParenteses

import "testing"

func TestIsValidSuccessLenTwo(t *testing.T) {
	testStr := "[]"
	b := isValid(testStr)
	if b != true {
		t.Errorf("Result should be 'true', actual result: %v", b)
	}
}

func TestIsValidFail(t *testing.T) {
	b := isValid("}[")
	if b != false {
		t.Errorf("Result should be 'false', actual result: %v", b)
	}

	a := isValid("(]")
	if a != false {
		t.Errorf("result should be 'false', actual result: %v", a)
	}
}

func TestIsValidSuccess02(t *testing.T) {
	testStr := "()[]{}"
	b := isValid(testStr)
	if b != true {
		t.Errorf("Result should be 'true', actual result: %v", b)
	}
	b = isValid("{[]}")
	if b != true {
		t.Errorf("Result should be 'true', actual result: %v", b)
	}
	b = isValid("[({(())}[()])]")
	if b != true {
		t.Errorf("Result should be 'true', actual result: %v", b)
	}
}
