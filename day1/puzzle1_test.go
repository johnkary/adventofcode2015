package puzzle1

import "testing"

func TestPuzzle1(t *testing.T) {
    cases := []struct {
        in string
        want int
    }{
        {"(())", 0},
        {"()()", 0},
        {"(((", 3},
        {"(()(()(", 3},
        {"))(((((", 3},
//        {"())", -1},
//        {"))(", -1},
//        {")))", -3},
//        {")())())", -3},
    }
    for _, c := range cases {
        got := Puzzle1(c.in)
        if got != c.want {
            t.Errorf("Expected %q but got %q with input: %q", string(c.want), string(got), c.in)
        }
    }
}
