package normalizer

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const msg = "test case %d failed"

var testCasesUnquote = []struct {
	in  string
	out string
}{
	{"'\x00'", "\u0000"},
	{`'\0'`, "\u0000"},
	{`'\0something\0'`, "\u0000something\u0000"},
	{`'\0something\0somethingElse'`, "\u0000something\u0000somethingElse"},
}

func TestUnquoteSingle(t *testing.T) {
	for i, test := range testCasesUnquote {
		s, err := unquoteSingle(test.in)
		require.NoError(t, err, msg, i)

		require.Equal(t, test.out, s, msg, i)
	}
}

var testCasesUnquoteAndQuoteBack = []string{"'\x00'", "'rand'"}

func TestUnquoteSingleAndQuoteBack(t *testing.T) {
	for i, test := range testCasesUnquoteAndQuoteBack {
		s, err := unquoteSingle(test)
		assert.NoError(t, err, msg, i)
		q := quoteSingle(s)

		assert.Equal(t, test, q, msg, i)
	}
}

func BenchmarkReplacingNullEscape_Iterative(b *testing.B) {
	b.ReportAllocs()
	s := testCasesUnquote[3].in
	for n := 0; n < b.N; n++ {
		replaceEscapedMaybe(s, '0', '\x00')
	}
}

func BenchmarkReplacingNullEscape_Regexp(b *testing.B) {
	b.ReportAllocs()
	s := testCasesUnquote[3].in
	for n := 0; n < b.N; n++ {
		replaceEscapedMaybeRegexp(s)
	}
}

var re = regexp.MustCompile(`\\0([^0-9]|$)`)

// replaceEscapedMaybeRegexp is very simple, but slower alternative to normalizer.replaceEscapedMaybe
func replaceEscapedMaybeRegexp(s string) string {
	return re.ReplaceAllString(s, "\x00$1")
}
