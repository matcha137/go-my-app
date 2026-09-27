func doSomething() {
    var n = 1
    defer func() {
        fmt.Println(n)
    }()

    n = 2
}