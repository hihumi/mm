import (
	"./mylib"
	"testing"
)

func TestSwitchNameFuncn(t *testing.T) {
	expectCd := "cd"
	actualCd := mylib.CdStringFunc()

	if expectCd != actualCd {
		t.Errorf("%s != %s\n", expectCd, actualCd)
	}

	expectPwd := "pwd"
	actualPwd := mylib.PwdStringFunc()

	if expectPwd != actualPwd {
		t.Errorf("%s != %s\n", expectPwd, actualPwd)
	}
}
