package log

import "fmt"

func Error(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	} else {
		return false
	}
}
