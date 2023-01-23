package utils

import "fmt"

var Debug bool = false

func ConsoleLog(v ...interface{}) {

	if !Debug {
		return
	}

	fmt.Println(v...)
}
