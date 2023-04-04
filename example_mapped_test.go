package reggie_test

import (
	"fmt"
	"github.com/jsumners/go-reggie"
)

// This example shows how the [Reggie.GroupMatches] utility provides the
// named capture groups as a map of [github.com/samber/mo.Option] values.
func Example_mapped() {
	regex := reggie.MustCompile(`(?P<foo>\w{3,}) (?P<bar>\w{3,})`)
	input := "two words"
	matches := regex.GroupMatches(input)

	if len(matches) != 2 {
		panic("did not get any matches!")
	}

	foo := matches["foo"].MustGet()
	bar := matches["bar"].MustGet()
	fmt.Printf("found words: (`%s`, `%s`)", foo, bar)

	// Output:
	// found words: (`two`, `words`)
}
