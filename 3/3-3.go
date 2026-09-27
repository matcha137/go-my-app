package main

import (
	"os"
	"path/filepath"
)

func doSomething(dir string) error {
	err := os.Mkdir(dir, 0755)
    if err != nil {
        return err
    }
    defer os.RemoveAll(dir)

    f, err := os.Create(filepath.Join(dir, "data.txt"))
    if err != nil {
        return err
    }
    defer f.Close()

    // do something
}
