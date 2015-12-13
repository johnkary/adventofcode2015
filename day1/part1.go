package main

import (
    "strings"
)

func part1(instructions string) int {
    up := strings.Count(instructions, "(")
    down := strings.Count(instructions, ")")

    return up - down
}
