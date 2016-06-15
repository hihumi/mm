package main

import (
	"fmt"
	"github.com/hihumi/mm/mm"
)

func SwitchCommandName(commandName string) {

	var (
		cd  = mm.CdReturnString()
		pwd = mm.PwdReturnString()
	)

	switch commandName {
	case cd:
		mm.CdPrint()
	case pwd:
		mm.PwdPrint()
	default:
		fmt.Println("見つかりませんでした")
	}

	return
}
