package reggie_test

import (
	"fmt"
	"github.com/jsumners/go-reggie"
)

// This example shows how the [Reggie.SubmatchWithName] utility provides easy
// access to the values associated with the capture group names.
func Example_basic() {
	regex := reggie.MustCompile(`(?P<foo>\w{3,}) (?P<bar>\w{3,})`)
	input := "two words"
	matches := regex.FindStringSubmatch(input)

	if matches == nil {
		panic("did not get any matches!")
	}

	foo := regex.SubmatchWithName("foo")
	bar := regex.SubmatchWithName("bar")
	fmt.Printf("found words: (`%s`, `%s`)", foo, bar)

	// Output:
	// found words: (`two`, `words`)
}
