package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a interface{} = "vinay"
	fmt.Println(reflect.TypeOf(a))
}
