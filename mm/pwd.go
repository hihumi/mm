package mm

import (
	"fmt"
)

type PwdStruct struct {
	pwdString interface{}
	pwd       interface{}
}

func PwdStringFunc() interface{} {
	var Pwd PwdStruct
	Pwd.pwdString = "pwd"

	return Pwd.pwdString

}

func PwdPrintFunc() {
	var Pwd PwdStruct
	Pwd.pwd = PwdStringFunc()

	fmt.Printf("%v: 現在のディレクトリを表示します\n詳細: オンラインマニュアル: $ man %v\n", Pwd.pwd, Pwd.pwd)

	return
}
