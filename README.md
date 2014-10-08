# Xterm256

Xterm256 aims to be a Go package for printing colored string in a 256 colors terminal.

## Installation

```
go get github.com/gilliek/go-xterm256/xterm256
```

## Usage

```
// Use a predefined color
xterm256.Println(xterm256.Red, "Foo")

// Define a new color
orange, _ := xterm256.NewColor(3, 1, 0)
xterm256.Println(orange, "Bar")

// Add a background color
orange.SetBackground(5, 5, 5)
xterm256.Println(orange, "Bar")
```


## Documentation

Documentation can be found on either
[GoDoc](http://godoc.org/github.com/gilliek/go-xterm256/xterm256) or
[GoWalker](https://gowalker.org/github.com/gilliek/go-xterm256/xterm256).

## Contribute

You're very welcome to submit pull requests or report bugs :) For submitting a
pull request, just following GitHub instructions.

## License

BSD 3-clauses
