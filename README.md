# Twerp
## A REPL written in Go

Based on the book [`Writing An Interpreter in Go`](https://interpreterbook.com/) by Thorsten Ball.

All of the language features covered in the book are available in Twerp, with none of the polish!

### What am I?

Twerp, is both a terrible play on words (`in-twerp-preter`), a derp-ier version of `Monkey`, and a rebel without a cause, given it has no life, no one using it, and only exists for its creators edification...

### Born to Run
#### Or how to run

```bash
git clone https://github.com/darragh-downey/twerp
cd twerp
go run src/main.go
```

Give the following a go [pun intended...]
```go
// Integers & arithmetic expressions
let version = 1 + (50 / 2) - (8 * 3);

// Strings
let name = "The Twerp programming language";

// Booleans
let amIRichNow = false;

// Arrays & Hashes
let people = [{"name": "Anna", "age": 24}, {"name": "Bob", "age": 99}];

// Functions
let getName = fn(person) { person["name"]; };
getName(people[0]); // => "Anna"
getName(people[1]); // => "Bob"
```


#### Builtin fn's

Let's create an array `x`

```go
let x = [1, 2, 3];
```

##### len
To check `x's` length we can run:
```go
len(x)
```

##### min
To find the smallest element in `x` we can run:
```go
min(x)
```

##### max
To find the largest element in `x` we can run:
```go
max(x)
```

##### first
Same goes for the first element in `x`:

##### last
And the last element in `x`:

##### rest
Return all elements in the given array `x`, but the first:
```go
rest(x)
```

##### push
Create a new array by adding an element to the end of the array `x`:
```go
push(x, 5)
```

A design decision prevents modification of variables, so don't go expecting `x` to update just yet.

##### puts

Not yet implemented - will just print the passed objects to the REPL.

```go 
puts(x, 6)
```

### Next steps
Ironically, although a working REPL, Twerp will be extended with the following features:

- Remove the printed `null` 
- A `help` function, similar to that snake language, which might be helpful for potential users outside of the usual suspects...    
- A compiler, based on, you guessed it, the book [`Writing A Compiler in Go`](https://compilerbook.com/) by Thorsten Ball
