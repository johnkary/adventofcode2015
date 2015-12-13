## Day 1: Not Quite Lisp

### Usage

Ensure tests pass then compile + run:

    $ cd adventofcode/day1
    $ go test
    $ go build && ./day1

### What I learned

I previously setup Go on my machine a few months ago but didn't get around to using it. I made sure I had the latest Go binary via `$ brew upgrade go` then went through the [Go - Getting Started](https://golang.org/doc/install) guide to ensure my system configuration up to date.

The [How to Write Go Code](https://golang.org/doc/code.html) reminded me that running Go code was easy:

    $ cd adventofcode/day1
    $ go run

This only works if you have a file with `package main` at the top of one of the .go source files and it has a function called main: `func main()`. Go identifies this as the entry point for your program.

Since I do test-driven development I was unsure for a minute how to I would execute my code. I knew I wanted the test suite to exercise my code. Luckily the guide introduced `go test` to run a package's test suite. The file "puzzle1.go" has its tests in "puzzle1\_test.go".

I went through the motions typing in the example test code. It quickly shows off example code using `struct` which felt a bit foreign. My PHP brain thought, "I want an array of arrays where value[0] is the input and value[1] is the expected result." But Go presented structs, a data structure defined on the fly. My understanding is they're a typed data structure without a name. I think this is like a PHP array with keys as strings but the Go compiler can enforce property/key access. In PHP this would be like `['input' => '(())', 'expected' => 0]`. If you tried $array['foo'] it would throw an error, and Go similarly wouldn't compile.

The syntax for defining the struct is followed by immediately defining a collection (slice? list? I don't know yet) of instances of that struct. The Go example's struct had two values `in, want string` but I wanted `in string, want int`. I thought I could do instances of the struct via `{"(())", 0}` but nope...

Now I needed to organize the code. I like order and want everything to
have a place. I found Dave Chaney's article
[Five suggestions for setting up a Go project](http://dave.cheney.net/2014/12/01/five-suggestions-for-setting-up-a-go-project)
very helpful. I decided each puzzle would have its own executable. So
Day 1, Puzzle 1 would be in `day1_1` and Day 1, Puzzle 2 in `day1_2`.
Long-term this is probably terrible. I originally wanted a high level
"adventofcode" program which allowed me to select which day's puzzle to
run. That felt like too much effort for just getting started. So I'm
pushing down the anxiety I feel with a poor structure and moving forward
with the important part, which is writing code in Go.

### Part 1

Santa was hoping for a white Christmas, but his weather machine's "snow"
function is powered by stars, and he's fresh out! To save Christmas, he
needs you to collect fifty stars by December 25th.

Collect stars by helping Santa solve puzzles. Two puzzles will be made
available on each day in the advent calendar; the second puzzle is
unlocked when you complete the first. Each puzzle grants one star. Good
luck!

Here's an easy puzzle to warm you up.

Santa is trying to deliver presents in a large apartment building, but
he can't find the right floor - the directions he got are a little
confusing. He starts on the ground floor (floor 0) and then follows the
instructions one character at a time.

An opening parenthesis, (, means he should go up one floor, and a
closing parenthesis, ), means he should go down one floor.

The apartment building is very tall, and the basement is very deep; he
will never find the top or bottom floors.

For example:

* (()) and ()() both result in floor 0.
* ((( and (()(()( both result in floor 3.
* ))((((( also results in floor 3.
* ()) and ))( both result in floor -1 (the first basement level).
* ))) and )())()) both result in floor -3.

To what floor do the instructions take Santa?

### Part 2

Now, given the same instructions, find the position of the first character that causes him to enter the basement (floor -1). The first character in the instructions has position 1, the second character has position 2, and so on.

For example:

    - `)` causes him to enter the basement at character position `1`.
    - `()())` causes him to enter the basement at character position `5`.

What is the position of the character that causes Santa to first enter the basement?
