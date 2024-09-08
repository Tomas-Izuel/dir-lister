package utils

import "fmt"

func ErrHandler(err error, isPanic bool) {
	if err != nil {
		if isPanic {
			panic(err)
		} else {
			fmt.Println(err)
		}
	}
}
