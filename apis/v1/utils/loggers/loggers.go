package loggers

import "fmt"

func Log(params ...interface{}) {
	fmt.Println(params...)
}
