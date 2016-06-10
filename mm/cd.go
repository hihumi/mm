package mylib

import (
	"fmt"
)

type CdStruct struct {
	cdString interface{}
	cd       interface{}
}

func CdStringFunc() interface{} {
	var Cd CdStruct
	Cd.cdString = "cd"

	return Cd.cdString
}

func CdPrintFunc() {
	var Cd CdStruct
	Cd.cd = CdStringFunc()

	fmt.Printf("%v: カレントディレクトリを変更します\n詳細: オンラインマニュアル: $ man %v\n", Cd.cd, Cd.cd)

	return
}
