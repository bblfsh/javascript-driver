package normalizer

import (
	"regexp"
	"testing"
	"unicode/utf8"

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
	for n := 0; n < b.N; n++ {
		for _, test := range testCasesUnquote {
			replaceEscapedMaybeIter(test.quoted, '0', '\x00')
		}
	}
}

func replaceEscapedMaybeIter(s string, old, new rune) string {
	var runeTmp [utf8.UTFMax]byte
	n := utf8.EncodeRune(runeTmp[:], new)

	lastCp := 0
	var buf []byte
	for i, w := 0, 0; i < len(s); i += w {
		r1, w1 := utf8.DecodeRuneInString(s[i:])
		w = w1
		if r1 == '\\' { // find sequence \\old[^0-9]
			r2, w2 := utf8.DecodeRuneInString(s[i+w1:])
			if r2 == old {
				r3, _ := utf8.DecodeRuneInString(s[i+w1+w2:])
				if 0 > r3 || r3 > 9 { // not a number after "\\old"
					w += w2
					if len(buf) == 0 {
						buf = make([]byte, 0, 3*len(s)/2)
					}
					buf = append(buf, []byte(s[lastCp:i])...)
					buf = append(buf, runeTmp[:n]...)
					lastCp = i + w
				}
			}
		}
	}
	if lastCp == 0 {
		return s
	}

	if 0 < lastCp && lastCp < len(s) {
		return string(append(buf, []byte(s[lastCp:len(s)])...))
	}
	return string(buf)
}

func BenchmarkReplacingNullEscape_Regexp(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		for _, test := range testCasesUnquote {
			replaceEscapedMaybeRegexp(test.quoted)
		}
	}
}

var re = regexp.MustCompile(`\\0([^0-9]|$)`)

// replaceEscapedMaybeRegexp is very simple, but slower alternative to normalizer.replaceEscapedMaybe
func replaceEscapedMaybeRegexp(s string) string {
	return re.ReplaceAllString(s, "\x00$1")
}

func BenchmarkReplacingNullEscape_Simple(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		for _, test := range testCasesUnquote {
			replaceEscapedMaybe(test.quoted, "\\0", "\x00")
		}
	}
}
