# werrors
[![GoDoc](https://godoc.org/github.com/PaulRaUnite/werrors?status.svg)](https://godoc.org/github.com/PaulRaUnite/werrors)
[![Go Report Card](https://goreportcard.com/badge/github.com/PaulRaUnite/werrors)](https://goreportcard.com/report/github.com/PaulRaUnite/werrors)
[![cover.run go](https://cover.run/go/github.com/PaulRaUnite/werrors.svg)](https://cover.run/go/github.com/PaulRaUnite/werrors)
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

Of course, allocation free functions are faster.

## Documentation + Examples

See [godoc](https://godoc.org/github.com/PaulRaUnite/werrors)

## Collaboration

Common scheme:
 - Create an issue and discuss a problem
  (omit it if the issue is small, something like
   grammar mistake in README :D);
 - Fork repository and commit your changes;
 - Pull request with reference to the issue;
 - Wait for review (maybe with some discussion);
 - Merging.

If your request wasn't merged, don't worry, you always
can create a fork. :wink:

## Possible questions

### Why do you use strings as annotation container?

Because of their immutability.

Yes, you can use some hacks (`unsafe` package can give
you direct access to `[]byte` inside `string`) to bypass
the constraint and change something inside,
but don't do it, it's ridiculous. :)
In this case, of course.
 
It can be helpful to optimize builtin conversion between
`[]byte` to `string` applying your own conversion 
in cases of high performance when you are sure 
in the immutability of byte slice, of course.
