package main

import (
	"./mm"
	"testing"
)

func TestSwitchNameFunc(t *testing.T) {
	expectCd := "cd"
	actualCd := mm.CdStringFunc()

	if expectCd != actualCd {
		t.Errorf("%s != %s\n", expectCd, actualCd)
	}

	expectPwd := "pwd"
	actualPwd := mm.PwdStringFunc()

	if expectPwd != actualPwd {
		t.Errorf("%s != %s\n", expectPwd, actualPwd)
	}
}
