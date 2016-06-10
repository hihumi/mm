package mylib

import (
	"fmt"
)

func PwdStringFunc() string {
	pwdString := "pwd"

	return pwdString

}

func PwdPrintFunc() {
	pwd := PwdStringFunc()

	fmt.Printf("%s: 現在のディレクトリを表示します\n詳細: オンラインマニュアル: $ man %s\n", pwd, pwd)

	return
}
