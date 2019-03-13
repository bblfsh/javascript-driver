package normalizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testCasesUnquote = []struct {
	quoted   string
	unquoted string
	// In some cases unquoting and then re-quoting a quoted string does not produce a
	// string that is bitwise identical to the original, even though they denote the same bytes.
	// This can happen, e.g, if we switch between hex and octal encoding of a byte.
	// Test cases where this happens set canonicalUnquoted to the string that is expected
	// to be decoded via Go's native rules to the byte sequence we want.
	canonicalQuoted string
}{
	{`'a'`, "a", `'a'`},
	{`'\x00'`, "\u0000", `'\x00'`},
	{`'\0'`, "\u0000", "'\\x00'"},
	{`'\0something\0'`, "\u0000something\u0000", "'\\x00something\\x00'"},
	{`'\0something\0else'`, "\u0000something\u0000else", "'\\x00something\\x00else'"},
	{`'\u0000123\0s'`, "\u0000123\u0000s", "'\\x00123\\x00s'"},
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
			assertEquals(t, test.canonicalQuoted, q)
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

func BenchmarkReplacingNullEscape(b *testing.B) {
	b.ReportAllocs()
	for _, test := range testCasesUnquote {
		for n := 0; n < b.N; n++ {
			replaceEscapedMaybe(test.quoted, "\\0", "\x00")
		}
	}
}
