package normalizer

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/bblfsh/sdk.v1/uast"
)

func TestAnnotate(t *testing.T) {
	require := require.New(t)
	n, err := annotatedFixture("hello.js.json")
	require.NoError(err)

	missingRole := make(map[string]bool)
	iter := uast.NewOrderPathIter(uast.NewPath(n))
	for {
		n := iter.Next()
		if n.IsEmpty() {
			break
		}

		missingRole[n.Node().InternalType] = true
	}

	for k := range missingRole {
		fmt.Println("NO ROLE", k)
	}
}

func TestAnnotatePrettyAnnotationsOnly(t *testing.T) {
	require := require.New(t)

	f, err := getFixture("hello.js.json")
	require.NoError(err)

	n, err := ToNode.ToNode(f)
	require.NoError(err)
	require.NotNil(n)

	err = AnnotationRules.Apply(n)
	require.NoError(err)

	buf := bytes.NewBuffer(nil)
	err = uast.Pretty(n, buf, uast.IncludeAnnotations|uast.IncludeChildren|uast.IncludeTokens)
	require.NoError(err)
	fmt.Println(buf.String())
}

func TestNodeTokens(t *testing.T) {
	require := require.New(t)

	f, err := getFixture("hello.js.json")
	require.NoError(err)

	n, err := ToNode.ToNode(f)
	require.NoError(err)
	require.NotNil(n)

	tokens := uast.Tokens(n)
	require.True(len(tokens) > 0)
	for _, tk := range tokens {
		fmt.Println("TOKEN", tk)
	}
}
