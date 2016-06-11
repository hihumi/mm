package main

import (
	"./mm"
	"fmt"
)

func CommandNameFunc(commandName string) {

	var (
		cd  = mm.CdStringFunc()
		pwd = mm.PwdStringFunc()
	)

	switch true {
	case commandName == cd:
		mm.CdPrintFunc()
	case commandName == pwd:
		mm.PwdPrintFunc()
	default:
		fmt.Println("見つかりませんでした")
	}

	return
}
