package main

import (
	"./mm"
	"fmt"
)

func SwitchCommandNameFunc(commandName string) {

	var (
		cd  = mm.CdStringFunc()
		pwd = mm.PwdStringFunc()
	)

	switch {
	case commandName == cd:
		mm.CdPrintFunc()
	case commandName == pwd:
		mm.PwdPrintFunc()
	default:
		fmt.Println("見つかりませんでした")
	}

	return
}
