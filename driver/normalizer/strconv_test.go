package normalizer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnquoteSingle_NullGo(t *testing.T) {
	s, err := unquoteSingle("'\x00'")
	require.NoError(t, err)

	require.Equal(t, "\u0000", s)
}

func TestUnquoteSingle_NullJs(t *testing.T) {
	s, err := unquoteSingle(`'\0'`)
	require.NoError(t, err)

	require.Equal(t, "\u0000", s)
}

func TestUnquoteSingle_NullAndQuoteBack(t *testing.T) {
	const o = "'\x00'"

	s, err := unquoteSingle(o)
	require.NoError(t, err)
	q := quoteSingle(s)

	require.Equal(t, o, q)
}

func TestUnquoteSingle_StringAndQuoteBack(t *testing.T) {
	const o = "'rand'"

	s, err := unquoteSingle(o)
	require.NoError(t, err)
	q := quoteSingle(s)

	require.Equal(t, o, q)
}
