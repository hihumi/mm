package main

import (
	"fmt"
)

func main() {
	fmt.Println("Unixコマンドの概要を表示します。コマンド名を入力してください。(終了するには、q、またはQキーを入力します)")

	for {
		fmt.Printf(">>> ")
		var commandName string
		fmt.Scanln(&commandName)

		q := "q"
		Q := "Q"
		if q == commandName || Q == commandName {
			fmt.Println("終了")
			break
		}

		SwitchCommandNameFunc(commandName)
	}
}
