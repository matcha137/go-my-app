package main

import (
	"fmt"
)

func doSomething() {
    var n = 1
    defer fmt.Println(n)

    n = 2
}