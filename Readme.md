# Reggie

A simple wrapper on top of [`regexp.Regexp`](https://pkg.go.dev/regexp#Regexp)
to make working with named capture groups easier.

See [example_basic_test.go](./example_basic_test.go) for a full example.

## Install

```sh
$ go get github.com/jsumners/go-reggie
```

## Example

```go
package main

import "github.com/jsumners/go-reggie"
import "fmt"

func main() {
	regex := reggie.MustCompile(`(?P<foo>\w{3})`)
	matches := regex.FindStringSubmatch("bar")
	
	if matches == nil {
		panic("💥")
	}
		
	foo := regex.SubmatchWithName("foo")
	fmt.Println(foo)
	
	// Output:
	// bar
}
```
