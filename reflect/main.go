package main

import (
	"fmt"
	"reflect"
)

type S struct {
	F string `default:"defaultValue"`
}

func main() {
	var s S
	st := reflect.TypeOf(s)
	fmt.Println(st.Kind() == reflect.Struct)
	field := st.Field(0)
	fmt.Println(field.Tag.Get("default"))
}
