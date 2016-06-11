package main

import (
	_ "./mm"
	"fmt"
)

func main() {
	fmt.Println("Unixコマンドの概要を表示します。コマンド名を入力してください。(終了するには、q、またはQキーを入力します)")

	for {
		fmt.Printf(">>> ")
		var name string
		fmt.Scanln(&name)

		q := "q"
		Q := "Q"
		if q == name || Q == name {
			fmt.Println("終了")
			break
		}

		SwitchNameFunc(name)
	}
}
