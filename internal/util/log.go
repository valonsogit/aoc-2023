package util

import "fmt"

func StructLog(i interface{}) {
	fmt.Printf("%+v\n", i)
}
