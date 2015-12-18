## Day 3: Perfectly Spherical Houses in a Vacuum

### Usage

Ensure tests pass then compile + run:

    $ cd adventofcode/day3
    $ go test
    $ go build && ./day3

### Part 1

Santa is delivering presents to an infinite two-dimensional grid of houses.

He begins by delivering a present to the house at his starting location, and then an elf at the North Pole calls him via radio and tells him where to move next. Moves are always exactly one house to the north (^), south (v), east (>), or west (<). After each move, he delivers another present to the house at his new location.

However, the elf back at the north pole has had a little too much eggnog, and so his directions are a little off, and Santa ends up visiting some houses more than once. How many houses receive at least one present?

For example:

* `>` delivers presents to `2` houses: one at the starting location, and one to the east.
* `^>v<` delivers presents to `4` houses in a square, including twice to the house at his starting/ending location.
* `^v^v^v^v^v` delivers a bunch of presents to some very lucky children at only `2` houses.

### Part 2

Coming soon.

### What I Learned

I'm feeling good about the Go I know so far. I've got the hang of object composition from the last puzzle and comfortable with organizing this project so far. I'm in the flow of test-driven development which feels familiar and safe.

It's about time I commit to some better tools for writing Go.

#### Configuring Sublime Text 2 for Go

Go has a great command-line toolset. One tool is [`gofmt`](https://golang.org/cmd/gofmt/) which formats your code to an agreed upon syntax format. I love this idea because it stops all formatting debates. I can't say I love all the formatting details but consistency is king.

I spent 1 hour learning [GoSublime](https://github.com/DisposaBoy/GoSublime) for integrating Go-specific commands into Sublime Text 2. The best being auto-complete (cmd+SPACE) and it running `gofmt` on the current file on save. I'll learn more of its features later.

Back to code-focused thoughts.

#### Generics

Part 1 felt like a classic collection processing problem. A stream of input characters translate into state changes.

Into modeling my objects I remember now [Go does not natively support generics]() which means I'll need to write my own methods for creating collections of `Coord` objects as `VisitedHouses.visited`. When writing PHP, which also doesn't support generics, I find myself often yearning for them. In PHP I would end up overriding methods in [Doctrine's ArrayCollection](https://github.com/doctrine/collections/blob/master/lib/Doctrine/Common/Collections/ArrayCollection.php#L53) to only allow adding `User` objects, for example. Fine, I'll try doing the same in Go?

I jump right in with TDD. The first Collection test is the behavior of an empty collection. That's tricky for this problem because the instructions say `He begins by delivering a present to the house at his starting location` which means the collection of visited Coord's must be initialized with a default value.

Go struct's don't support a constructor method--they're not objects in the traditional sense. How to handle this?

#### Struct creation with default values

Seems many people have this question: 71 upvotes on a [StackOverflow thread](http://stackoverflow.com/questions/18125625/constructors-in-go) about constructors in Go. To my relief one of the best answers references some official Go documentation on [Allocation with `new`](https://golang.org/doc/effective_go.html#allocation_new) and [Constructors and composite literals](https://golang.org/doc/effective_go.html#composite_literals).

The accepted pattern: given type `User`, provide a function called `NewUser()` that initializes an object ready to be used. You're encouraged to create other `New*()` methods to initialize `User` in different scenarios. I've seen this pattern called "Factory Methods" or "[Named Constructors](http://verraes.net/2014/06/named-constructors-in-php/)]". For example, "create new user from old user record" might have this named constructor:

    func NewUserFromOldUser(o OldUser) User {
        birthdate := time.Date(o.bday_year, o.bday_month, o.bday_day, 0, 0, 0, 0, time.UTC)

        return &User{
            // Please don't do this...
            name: o.first_name + " " + o.last_name,
            dob: birthdate,
        }
    }

#### What the heck is a rune?

This problem introduces a custom syntax for encoding data: Up, Down, Left, and Right are replaced by symbols < > ^ v. This means there are valid tokens and invalid tokens. How to handle this conversion?

For my first solution I chose to do simple string replacement on the movement input. In writing my test cases I envisioned what scenarios would cause that algorithm to fail. Naturally, if the input string contained a stray space or number my program would fail.

My gut approach was to iterate over each character checking for valid symbols and dropping any invalid symbols. The strings package supports [`Map`](https://golang.org/pkg/strings/#Map) which allows exactly that:

    validInput := strings.Map(func(s rune) rune {
        if s == "^" || s == ">" || s == "v" || s == "<" {
            return s
        }
        return -1
    }, input)

Running the tests yields problems with this code:

    $ go test
    ./direction.go:42: cannot convert "^" to type rune
    ./direction.go:42: invalid operation: s == "^" (mismatched types rune and string)
    ./direction.go:42: cannot convert ">" to type rune
    ./direction.go:42: invalid operation: s == ">" (mismatched types rune and string)
    ./direction.go:42: cannot convert "v" to type rune
    ./direction.go:42: invalid operation: s == "v" (mismatched types rune and string)
    ./direction.go:42: cannot convert "<" to type rune
    ./direction.go:42: invalid operation: s == "<" (mismatched types rune and string)

Go sees "^" as a string, and the individual character variable `s` as a rune. I feel frustrated at first but it's because I don't know any better. I set out to find the difference between a string and a rune and why I should care.

This blog post on [Strings, bytes, runes and characters in Go](https://blog.golang.org/strings) should be required reading very early on. Going into this project I thought a string was a string. Not in Go.

> At first, strings might seem too simple a topic for a blog post, but to use them well requires understanding not only how they work, but also the difference between a byte, a character, and a rune, the difference between Unicode and UTF-8, the difference between a string and a string literal, and other even more subtle distinctions.

I first learned about multi-byte strings thanks to the seemingly wrong behavior of PHP's `strlen()` with multi-byte characters (UTF-8):

    > strlen("a")
    1
    > strlen("é")
    2
    > strlen("∑")
    3

If I understand it correctly, Go requires we normalize all `strings` to `runes` when doing comparisons. It appears we can normalize strings like "^" into runes like so:

    up, _ := utf8.DecodeRuneInString("^")

So to loop over all runes in a given string and remove invalid characters we can do this:

    invalidCharsRemoved := strings.Map(func(s rune) rune {
        up, _ := utf8.DecodeRuneInString("^")
        if s == up {
            return s
        }

        return -1
    }, input)

It seems this rabbit hole was so deep I abandoned Advent of Code until
(at least) April 2016. I feel frustrated that string comparisons are
extra work but I understand the necessity to accommodate UTF-8.
