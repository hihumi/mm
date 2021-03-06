package mm

import (
	"fmt"
)

type CdStruct struct {
	cdString interface{}
	cd       interface{}
}

func CdReturnString() interface{} {
	var Cd CdStruct
	Cd.cdString = "cd"

	return Cd.cdString
}

func CdPrint() {
	var Cd CdStruct
	Cd.cd = CdReturnString()

	fmt.Printf("%v: カレントディレクトリを変更します\n詳細: オンラインマニュアル: $ man %v\n", Cd.cd, Cd.cd)

	return
}
