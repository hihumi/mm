package main

import (
	"fmt"
	"./mm"
	_ "os"
)

func SwitchNameFunc(x string) {

	var (
		cd  = mm.CdStringFunc()
		pwd = mm.PwdStringFunc()
	)

	switch true {
	case x == cd:
		mm.CdPrintFunc()
	case x == pwd:
		mm.PwdPrintFunc()
	default:
		fmt.Println("見つかりませんでした")
	}

	return
}
