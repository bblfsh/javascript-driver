package normalizer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNativeToNoder(t *testing.T) {
	require := require.New(t)

	f, err := getFixture("example_js1.json")
	require.NoError(err)

	n, err := ToNoder.ToNode(f)
	require.NoError(err)
	require.NotNil(n)
}
