# werrors

## Name

werrors := **W**rap **errors**

## Simple example

```go
package main

import (
    "errors"
    "github.com/PaulRaUnite/werrors"
    "fmt"
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

Allocation free | With allocation/s
----------------|----------------
`Wrap`          | `Wrapf`
`DefWrap`       | `DefWrapf`
`Cause`         | `tracker.Error`

Of course, allocation free functions are faster.

**Important**:
  > `tracker.Error` produces cascade of allocations
  > because of it's recursion nature. It depends of
  > the nesting of errors.
  
## Documentation + Examples

See [godoc](https://godoc.com/github.com/PaulRaUnite/werrors)

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
in immutability of byte slice, of course.