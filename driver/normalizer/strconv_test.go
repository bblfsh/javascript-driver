package normalizer

import (
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

func TestUnquoteSingle_StringAndQuoteBack(t *testing.T) {
	const o = "'rand'"

	s, err := unquoteSingle(o)
	require.NoError(t, err)
	q := quoteSingle(s)

	require.Equal(t, o, q)
}
