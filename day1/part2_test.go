package main

import "testing"

func TestPart2(t *testing.T) {
    cases := []struct {
        in string
        want int
    }{
        {")", 1},
        {"()())())()())(()", 5},
    }
    for _, c := range cases {
        got := part2(c.in)
        if got != c.want {
            t.Errorf("Expected %q but got %q with input: %q", c.want, got, c.in)
        }
    }
}
