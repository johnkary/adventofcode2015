package main

import (
    "io/ioutil"
    "fmt"
    "strings"
)

func main() {
    text := readFile("input.txt")

    fmt.Printf("Part 1: %d sqft\n", part1(text))
//    fmt.Printf("Part 2: %d\n", part2(string(input)))

    return
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func readFile(path string) string {
    input, err := ioutil.ReadFile(path)
    check(err)

    text := string(input)
    text = strings.Trim(text, "\n")

    return text
}
