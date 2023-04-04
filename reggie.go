// Package reggie provides utilities on top of [regexp] that make it easier
// to deal with named capture groups (submatches in Go parlance).
package reggie

import (
	"github.com/samber/mo"
	"regexp"
)

// Reggie is a wrapper for [regexp.Regexp].
type Reggie struct {
	*regexp.Regexp
	matches []string
}

// Compile compiles a regular expression string.
// See [regexp.Compile] for more detail.
func Compile(input string) (*Reggie, error) {
	re, error := regexp.Compile(input)
	if error != nil {
		return nil, error
	}
	return &Reggie{Regexp: re}, nil
}

// MustCompile compiles a regular expression string or panics if not possible.
// See [regexp.MustCompile] for more detail.
func MustCompile(input string) *Reggie {
	re := regexp.MustCompile(input)
	return &Reggie{Regexp: re}
}

// FindStringSubmatch evaluates the regular expression against the provided
// search string. Internally, this method updates the [Reggie] instance with
// the same set of matches as are returned from the function. This allows
// the [Reggie.SubmatchWithName] method to be used.
func (re *Reggie) FindStringSubmatch(input string) []string {
	re.matches = re.Regexp.FindStringSubmatch(input)
	return re.matches
}

// GroupMatches evaluates the regular expression against the provided
// search string and returns a map of the resulting matches wherein the key
// names are the names in the named capture groups of the regular expression.
func (re *Reggie) GroupMatches(input string) map[string]mo.Option[string] {
	result := make(map[string]mo.Option[string])
	matches := re.FindSubmatch([]byte(input))

	if matches == nil {
		return result
	}

	groupNames := re.SubexpNames()
	for _, name := range groupNames {
		if name == "" {
			// The first element in the names array is always the empty string.
			continue
		}
		idx := re.SubexpIndex(name)

		match := matches[idx]
		if match != nil {
			result[name] = mo.Some[string](string(match))
		} else {
			result[name] = mo.None[string]()
		}
	}

	return result
}

// SubmatchWithName provides easy access to matches based upon the names of
// capture groups. The [Reggie.FindStringSubmatch] method **must** be invoked
// prior to using this method.
func (re *Reggie) SubmatchWithName(name string) string {
	index := re.Regexp.SubexpIndex(name)

	if index == -1 {
		return ""
	}

	return re.matches[index]
}
