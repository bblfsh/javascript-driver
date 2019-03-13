package normalizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testCasesUnquote = []struct {
	quoted   string
	unquoted string
	// If this is non-empty it means that quoteing back unqoted string does not
	// produce same result bit-wise.
	// This happens when we lose the information about original escape sequence (octal, hex)
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
	for _, test := range testCasesUnquote {
		t.Run("", func(t *testing.T) {
			s, err := unquoteSingle(test.quoted)
			require.NoError(t, err)
			require.Equal(t, test.unquoted, s)
		})
	}
}

func TestUnquoteSingleAndQuoteBack(t *testing.T) {
	for _, test := range testCasesUnquote {
		t.Run("", func(t *testing.T) {
			u, err := unquoteSingle(test.quoted)
			require.NoError(t, err)

			q := quoteSingle(u)
			if test.canonicalQuoted != "" {
				assertEquals(t, test.canonicalQuoted, q)
			} else {
				assertEquals(t, test.quoted, q)
			}
		})
	}
}

func assertEquals(t *testing.T, quoted, actual string) {
	if !assert.Equal(t, quoted, actual) {
		printDebug(t, quoted, actual)
		t.FailNow()
	}
}

func printDebug(t *testing.T, quoted, actual string) {
	t.Logf("\texpected: len=%d", len(quoted))
	for _, c := range quoted {
		t.Logf("%x - %#U", c, c)
	}
	t.Logf("\n\tactual: len=%d", len(actual))
	for _, c := range actual {
		t.Logf("%x - %#U", c, c)
	}
}

func BenchmarkReplacingNullEscape_Simple(b *testing.B) {
	b.ReportAllocs()
	for _, test := range testCasesUnquote {
		for n := 0; n < b.N; n++ {
			replaceEscapedMaybe(test.quoted, "\\0", "\x00")
		}
	}
}
