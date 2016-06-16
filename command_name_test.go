package main

import (
	"github.com/hihumi/mm/mm"
	"testing"
)

func TestSwitchCommandName(t *testing.T) {
	// mm.CdReturnString() (cd.go: CdReturnString())
	CdExpect := "cd"
	CdActual := mm.CdReturnString()

	if CdExpect != CdActual {
		t.Errorf("%v != %v\n", CdExpect, CdActual)
	}

	// mm.PwdReturnString() (pwd.go: PwdReturnString())
	PwdExpect := "pwd"
	PwdActual := mm.PwdReturnString()

	if PwdExpect != PwdActual {
		t.Errorf("%v != %v\n", PwdExpect, PwdActual)
	}
}
