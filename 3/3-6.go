package main

import (
	"fmt"
)

func doSomething() {
    var n = 1
    n = 2

    func() {
        fmt.Println(n)
    }()
}