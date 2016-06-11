package main

import (
	"github.com/hihumi/mm/mm"
	"testing"
)

func TestSwitchCommandName(t *testing.T) {
	expectCd := "cd"
	actualCd := mm.CdReturnString()

	if expectCd != actualCd {
		t.Errorf("%v != %v\n", expectCd, actualCd)
	}

	expectPwd := "pwd"
	actualPwd := mm.PwdReturnString()

	if expectPwd != actualPwd {
		t.Errorf("%v != %v\n", expectPwd, actualPwd)
	}
}
