package main

import (
	"fmt"
	"reflect"
)

func PrintDetail(v interface{}) {
	rt := reflect.TypeOf(v)
	switch rt.Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("int/int32/int64 型:", v)
	case reflect.String:
		fmt.Println("string型:", v)
	default:
		fmt.Println("知らない型")
	}
}

func main() {
	type V string
	var v V = "foobar"
	PrintDetail(v)
}