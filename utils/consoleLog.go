package utils

import "fmt"

var Debug bool = true

func ConsoleLog(v ...interface{}) {

	if !Debug {
		return
	}
	
	fmt.Println(v...)
}
