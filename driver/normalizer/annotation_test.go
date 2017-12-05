package normalizer

import (
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
