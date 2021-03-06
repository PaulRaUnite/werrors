# werrors
[![GoDoc](https://godoc.org/github.com/PaulRaUnite/werrors?status.svg)](https://godoc.org/github.com/PaulRaUnite/werrors)
[![Go Report Card](https://goreportcard.com/badge/github.com/PaulRaUnite/werrors)](https://goreportcard.com/report/github.com/PaulRaUnite/werrors)
 [![cover.run](https://cover.run/go/github.com/PaulRaUnite/werrors.svg?style=flat&tag=golang-1.10)](https://cover.run/go?tag=golang-1.10&repo=github.com%2FPaulRaUnite%2Fwerrors)
## Name

werrors := **w**rap **errors**

## Simple example

```go
package main

import (
    "errors"
    "fmt"
    
    "github.com/PaulRaUnite/werrors"
)

func main() {
    err := errors.New("error")
    err = werrors.Wrap(err, "annotation")
    fmt.Println(err)
}
```
`Output: annotation *> error`

## Purpose

The package was created as a part of my pet project,
because I needed something for tracking
errors and adding context of a lot of
calls of the same function with different arguments
(file names).

## Installation

`go get github.com/PaulRaUnite/werrors`

## Dependencies

Only `fmt` package.

## Performance

All functions can be divided into two groups:

Allocation free | With allocations
----------------|----------------
`Wrap`          | `Wrapf`
`DefWrap`       | `DefWrapf`
`Cause`         | `tracker.Error`
 &mdash;        | `tracker.Bytes`

Of course, allocation free functions are faster.

## Documentation + Examples

See [godoc](https://godoc.org/github.com/PaulRaUnite/werrors)

## Contributing

[There](/.github/CONTRIBUTING.md) more info.

## Possible questions

### Why do you use strings as annotation container?

Because of their immutability.

Yes, you can use some hacks (`unsafe` package can give
you direct access to `[]byte` inside `string`) to bypass
the constraint and change something inside,
but don't do it, it's ridiculous. :)

In the case above, of course.<br>
It can be helpful to optimize builtin conversion between
`[]byte` to `string` applying your own conversion 
in cases of high performance and when you are sure 
in the immutability of byte slice, of course.
