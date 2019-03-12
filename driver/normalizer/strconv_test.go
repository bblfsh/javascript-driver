package normalizer

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const msg = "test case %d failed"

var testCasesUnquote = []struct {
	quoted   string
	unquoted string
	// isn't empty only for cases where quote(unqote(s)) != s
	// e.g when we loose the informaiton wich original escape sequence was used
	// Golang unquote() defaults to hex format, so it's used as canonical one.
	canonicalQuoted string
}{
	{"'a'", "a", ""},
	{"'\\x00'", "\u0000", ""},
	{`'\0'`, "\u0000", "'\\x00'"},
	{`'\0something\0'`, "\u0000something\u0000", "'\\x00something\\x00'"},
	{`'\0something\0else'`, "\u0000something\u0000else", "'\\x00something\\x00else'"},
	{"'\u0000123\\0s'", "\u0000123\u0000s", "'\\x00123\\x00s'"},
}

func TestUnquoteSingle(t *testing.T) {
	for i, test := range testCasesUnquote {
		s, err := unquoteSingle(test.quoted)
		require.NoError(t, err, msg, i)

		require.Equal(t, test.unquoted, s, msg, i)
	}
}

func TestUnquoteSingleAndQuoteBack(t *testing.T) {
	for i, test := range testCasesUnquote {
		u, err := unquoteSingle(test.quoted)
		require.NoError(t, err, msg, i)

		q := quoteSingle(u)
		if test.canonicalQuoted != "" {
			assertEquals(t, test.canonicalQuoted, q, i)
		} else {
			assertEquals(t, test.quoted, q, i)
		}
	}
}

func assertEquals(t *testing.T, quoted, actual string, i int) {
	success := assert.Equal(t, quoted, actual, msg, i)
	if !success {
		printDebug(t, quoted, actual)
		t.FailNow()
	}
}

func printDebug(t *testing.T, quoted, actual string) {
	t.Logf("\texepcted: len=%d", len(quoted))
	for _, c := range quoted {
		t.Logf("%x - %#U", c, c)
	}
	t.Logf("\n\tactual: len=%d", len(actual))
	for _, c := range actual {
		t.Logf("%x - %#U", c, c)
	}
}

func BenchmarkReplacingNullEscape_Iterative(b *testing.B) {
	b.ReportAllocs()
	s := testCasesUnquote[3].quoted
	for n := 0; n < b.N; n++ {
		replaceEscapedMaybe(s, '0', '\x00')
	}
}

func BenchmarkReplacingNullEscape_Regexp(b *testing.B) {
	b.ReportAllocs()
	s := testCasesUnquote[3].quoted
	for n := 0; n < b.N; n++ {
		replaceEscapedMaybeRegexp(s)
	}
}

var re = regexp.MustCompile(`\\0([^0-9]|$)`)

// replaceEscapedMaybeRegexp is very simple, but slower alternative to normalizer.replaceEscapedMaybe
func replaceEscapedMaybeRegexp(s string) string {
	return re.ReplaceAllString(s, "\x00$1")
}
