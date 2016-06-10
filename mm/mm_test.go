package mylib

import (
	"testing"
)

func TestCdStringFunc(t *testing.T) {
	expect := "cd"
	actual := CdStringFunc()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}

func TestPwdStringFunc(t *testing.T) {
	expect := "pwd"
	actual := PwdStringFunc()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
