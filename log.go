package main

import "fmt"

func logError(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	} else {
		return false
	}
}
