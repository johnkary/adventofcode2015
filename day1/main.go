package main

import (
    "io/ioutil"
    "fmt"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    input, err := ioutil.ReadFile("input.txt")
    check(err)

    fmt.Printf("Part 1: %d\n", part1(string(input)))

    return
}
