package main

import (
	"./mylib"
	"fmt"
	_ "os"
)

func SwitchNameFunc(x string) {

	var (
		cd = mylib.CdStringFunc()
		pwd = mylib.PwdStringFunc()
	)

	switch true {
	case x == cd:
		mylib.CdPrintFunc()
	case x == pwd:
		mylib.PwdPrintFunc()
	default:
		fmt.Println("見つかりませんでした")
	}

	return
}
