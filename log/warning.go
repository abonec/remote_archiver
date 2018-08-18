package log

import "fmt"

func Warningf(format string, a ...interface{}) {
	fmt.Printf(format, a...)

}
