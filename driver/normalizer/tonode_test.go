package normalizer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNativeToNode(t *testing.T) {
	require := require.New(t)

	f, err := getFixture("hello.js.json")
	require.NoError(err)
	n, err := ToNode.ToNode(f)
	require.NoError(err)
	require.NotNil(n)
}

func TestNativeToNodeOffset(t *testing.T) {
	require := require.New(t)

	f, err := getFixture("hello.js.json")
	require.NoError(err)
	n, err := ToNode.ToNode(f)
	require.NoError(err)
	require.NotNil(n)

	require.Equal(int(n.StartPosition.Col), 1)
	require.Equal(int(n.EndPosition.Col), 1)
}
