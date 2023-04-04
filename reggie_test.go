package reggie

import (
	"github.com/samber/mo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompile(t *testing.T) {
	t.Run("returns error for bad input", func(t *testing.T) {
		re, err := Compile("foo(bar")
		assert.Nil(t, re)
		assert.ErrorContains(t, err, "error parsing regexp: ")
	})

	t.Run("returns a Reggie", func(t *testing.T) {
		re, err := Compile("foo bar")
		assert.NotNil(t, re)
		assert.NoError(t, err)
	})
}

func TestMustCompile(t *testing.T) {
	re := MustCompile("foo")
	assert.NotNil(t, re)

	assert.Panics(t, func() { MustCompile("^foo(bar") })
}

func TestGroupMatches(t *testing.T) {
	t.Run("returns a map of results", func(t *testing.T) {
		re := MustCompile(`(?P<foo>\w{3}) (?P<bar>\w{3})`)
		matches := re.GroupMatches("foo bar")
		expected := map[string]mo.Option[string]{
			"foo": mo.Some("foo"),
			"bar": mo.Some("bar"),
		}
		assert.Equal(t, expected, matches)
		assert.Equal(t, "foo", matches["foo"].MustGet())
	})

	t.Run("does not overwrite first name with second instance", func(t *testing.T) {
		re := MustCompile(`(?P<foo>\w{3}) (?P<bar>\w{3}) (?P<foo>\d{1})`)
		matches := re.GroupMatches("foo bar 1")
		expected := map[string]mo.Option[string]{
			"foo": mo.Some("foo"),
			"bar": mo.Some("bar"),
		}
		assert.Equal(t, expected, matches)
	})

	t.Run("returns a map of results with optional groups", func(t *testing.T) {
		re := MustCompile(`(?P<foo>\w{3}) (?P<bar>\w{3})? (?P<baz>\d{3})`)
		matches := re.GroupMatches("foo  123")
		expected := map[string]mo.Option[string]{
			"foo": mo.Some("foo"),
			"bar": mo.None[string](),
			"baz": mo.Some("123"),
		}
		assert.Equal(t, expected, matches)
		assert.Equal(t, "foo", matches["foo"].MustGet())
	})

	t.Run("handles non-matches", func(t *testing.T) {
		re := MustCompile(`(?P<one>\d{4})-(?P<two>\d{2})-(?P<three>\d{2})`)
		matches := re.GroupMatches("2023/04/03")
		assert.Len(t, matches, 0)
	})
}

func TestFindSubstringMatch(t *testing.T) {
	re := MustCompile("^(foo) (bar)$")
	matches := re.FindStringSubmatch("foo bar")
	assert.Equal(t, matches, re.matches)
	assert.Len(t, re.matches, 3)

	matches = re.FindStringSubmatch("baz bar")
	assert.Nil(t, matches)
}

func TestSubmatchWithName(t *testing.T) {
	re := MustCompile("^(?P<foo>foo) (?P<bar>bar)$")
	re.FindStringSubmatch("foo bar")

	foo := re.SubmatchWithName("foo")
	assert.Equal(t, foo, "foo")

	bar := re.SubmatchWithName("bar")
	assert.Equal(t, bar, "bar")

	baz := re.SubmatchWithName("baz")
	assert.Empty(t, baz)
}
