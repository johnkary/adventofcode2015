## Day 2: I Was Told There Would Be No Math

### Usage

Ensure tests pass then compile + run:

    $ cd adventofcode/day2
    $ go test
    $ go build && ./day2

### Part 1

The elves are running low on wrapping paper, and so they need to submit an order for more. They have a list of the dimensions (length l, width w, and height h) of each present, and only want to order exactly as much as they need.

Fortunately, every present is a box (a perfect right rectangular prism), which makes calculating the required wrapping paper for each gift a little easier: find the surface area of the box, which is `2*l*w + 2*w*h + 2*h*l`. The elves also need a little extra paper for each present: the area of the smallest side.

For example:

* A present with dimensions `2x3x4` requires `2*6 + 2*12 + 2*8` = `52` square feet of wrapping paper plus `6` square feet of slack, for a total of `58` square feet.
* A present with dimensions `1x1x10` requires `2*1 + 2*10 + 2*10` = `42` square feet of wrapping paper plus `1` square foot of slack, for a total of `43` square feet.

All numbers in the elves' list are in feet. How many total square feet of wrapping paper should they order?

### Part 2

The elves are also running low on ribbon. Ribbon is all the same width, so they only have to worry about the length they need to order, which they would again like to be exact.

The ribbon required to wrap a present is the shortest distance around its sides, or the smallest perimeter of any one face. Each present also requires a bow made out of ribbon as well; the feet of ribbon required for the perfect bow is equal to the cubic feet of volume of the present. Don't ask how they tie the bow, though; they'll never tell.

For example:

* A present with dimensions `2x3x4` requires `2+2+3+3 = 10` feet of ribbon to wrap the present plus `2*3*4 = 24` feet of ribbon for the bow, for a total of `34` feet.
* A present with dimensions `1x1x10` requires `1+1+1+1 = 4` feet of ribbon to wrap the present plus `1*1*10 = 10` feet of ribbon for the bow, for a total of `14` feet.

How many total feet of ribbon should they order?

### What I Learned

I started wondering if Go had some geometry functions in its standard library. A quick scan didn't reveal any. Really? I have to build up these functions myself?

I searched for contributed libraries and came across "geo". If this met my needs it did far more than I needed. I'd probably learn more by implementing the functions myself. But "geo" did inspire me how to organize my code for this exercise.

I decided to build up some types for the Box, hang a public method "Area()" on it and go from there.

Following TDD I wrote out tests for a builder function to create a box given the "2x3x4" format the problem provides to return a Box with that width, length and height. Also wrote a second test for Box.Area().

At this point I have some tests but also some scratch/stub code in `main.go` to run my program. This code isn't "cleaned it up yet" so when I run the tests `go test` I start getting compile warnings:

    # github.com/johnkary/adventofcode/day2
    ./part1.go:4: imported and not used: "strings"
    ./part1.go:8: area(input) used as value
    ./part1.go:8: padding(input) used as value
    ./part1.go:12: too many arguments to return
    ./part1.go:15: too many arguments to return
    FAIL    github.com/johnkary/adventofcode/day2 [build failed]

I end up really frustrated because these failures prevent me from running the code I'm ready to hack on. Do I really have to clean up ALL of my codebase for the program to run?

Looking at `go test --help` it doesn't look like there's an option to run tests in just one file. I could run tests within one package via `test [-c] [-i] [build and test flags] [packages]`. Should I put my Box code in another package? Like "geo"? Does that make sense in Go?

#### Functions as values

Another thing I learned was when trying to use a function's return value within a math operation, it's important to do one thing. See if you can spot the issue:

    func part1(input string) int {
        return area(input) + padding(input)
    }

    func area(input string) {
        return 100
    }
    func padding(input string) {
        return 100
    }

The compiler tells me:

    ./part1.go:4: area(input) used as value
    ./part1.go:4: padding(input) used as value

Hmm, why is that? Ah. The `area()` and `padding()` function definitions are missing the `int` return type:

    func part1(input string) int {
        return area(input) + padding(input)
    }

    func area(input string) int {
        return 100
    }
    func padding(input string) int {
        return 100
    }

The compiler message `area(input) used as value` doesn't really point to the issue at hand, an unexpected type coercion, which obviously in Go is a big no-no.

#### Printf() format

The Printf() function returns a string representation of given values, hence "print, formatted" == "Printf". It's perfect for building an error message with `testing.Errorf()`

While I've memorized easy ones like %s (print as string) and %d (print as base-10 integer) from working in PHP, Go's fmt.Printf() has additional modifiers to print value type + fields.

I'm constantly referring back to the format page for Printf(). My mind best follows the examples at <https://gobyexample.com/string-formatting>.

#### More parsing string to int issues

Yet again I need to convert types for this problem. The input string "1x1x10" should convert to 1, 1, 10 for height, width, length respectively. My experience says, "split the string on `x` then use positional arguments to extract the data". Easy, right?

    func boxFromCoord(input string) Box {
        fields := strings.Split(input, "x")

        return Box{l:fields[0], w:fields[1], h:fields[2]}
    }

Compiler says:

    ./box.go:14: cannot use fields[0] (type string) as type int in field value
    ./box.go:14: cannot use fields[1] (type string) as type int in field value
    ./box.go:14: cannot use fields[2] (type string) as type int in field value

Yeah, I get it. Type safety. Ughhhhh. I don't want to worry about details like this. It feels like it slows me down from solving the real problem at hand. The code also ends up with more boilerplate that ultimately doesn't matter. The imperative code that does only exactly what you tell it feels verbose for me, who strives for simple concise code.

Google for "golang string to int" leads me to Stack Overflow for [Convert string to integer type in Go](http://stackoverflow.com/questions/4278430/convert-string-to-integer-type-in-go). It's late at night and I've been impatient most of today. The method I need is `strconv.Atoi()`. So obvious, why didn't I think of that! Curiosity gets the better of me so before I move on: Google "why named Atoi" which yields it stands for [ASCII to Integer](http://stackoverflow.com/a/2909772/438911) and has roots in C. Ha, alright, whatever. I hastily drop it in and run the code.

    import (
        "strings"
        "strconv"
    )

    func boxFromCoord(input string) Box {
        fields := strings.Split(input, "x")

        return Box{l:strconv.atoi(fields[0]), w:strconv.atoi(fields[1]), h:strconv.atoi(fields[2])}
    }

    $ go test
    ./box.go:14: multiple-value strconv.Atoi() in single-value context
    ./box.go:15: multiple-value strconv.Atoi() in single-value context
    ./box.go:16: multiple-value strconv.Atoi() in single-value context

A veteran Go developer would be familiar with this error message right away. Go functions often returns multiple values (a tuple): the value result of the function AND an error value. My novice Go developer mind feels like is every single function in the language. So you either handle the error or discard it. At this point I'm averse to boilerplate so I don't care.

    length, _ := strconv.atoi(fields[0])

The tests now pass. With all my details in order I can start working on the next higher level problem: deriving the area of said Box. Pretty straightforward: 2*L*W + 2*W*H + 2*H*L.

However I did find one awesome resource I continue referring back to: [strings/example_test.go](https://golang.org/src/strings/example_test.go). Test cases for many functions in the "strings" package. PHP.net has common usage examples in its documentation and I've come to expect that. Most Go package examples are in these `*_test.go` files.

#### Object composition

The Box type I created could be reused in many other contexts so we'll leave it as a separate package and begin working on our specific domain problem: wrapping presents contained in boxes.

A Christmas Present is a special kind of Box. My object-oriented mind says this might (MIGHT) be a place for inheritance. But Go doesn't support object inheritance. Instead, and I love this about Go, we create a Present type which is **composed of** a Box type + the custom behavior (functions) that answer higher-level domain questions like "How much wrapping paper is required to cover this box?"
