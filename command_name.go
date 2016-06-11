package main

import (
	"./mm"
	"fmt"
)

func SwitchCommandNameFunc(commandName string) {

	var (
		cd  = mm.CdReturnString()
		pwd = mm.PwdReturnString()
	)

	switch {
	case commandName == cd:
		mm.CdPrint()
	case commandName == pwd:
		mm.PwdPrint()
	default:
		fmt.Println("見つかりませんでした")
	}

	return
}
