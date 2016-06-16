package mm

import (
	"testing"
)

func TestCdReturnString(t *testing.T) {
	expect := "cd"
	actual := CdReturnString()

	if expect != actual {
		t.Errorf("%v != %v\n", expect, actual)
	}
}

func TestPwdReturnString(t *testing.T) {
	expect := "pwd"
	actual := PwdReturnString()

	if expect != actual {
		t.Errorf("%v != %v\n", expect, actual)
	}
}
