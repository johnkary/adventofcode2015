package main

import (
    "strings"
)

func part2(instructions string) int {
    characters := strings.Split(instructions, "")
    floor := 0
    i := 0

    // Doesn't handle use case where never enters basement
    for ; floor >= 0; i++ {
        symbol := characters[i]

        num := 1
        if symbol == ")" {
            num = -1
        }
        floor = floor + num
    }

    return i
}
