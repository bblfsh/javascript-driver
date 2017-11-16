package normalizer

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNativeToNoder(t *testing.T) {
	require := require.New(t)

	f, err := getFixture("hello.js.json")
	require.NoError(err)
	n, err := ToNode.ToNode(f)
	require.NoError(err)
	require.NotNil(n)
}

const fixtureDir = "fixtures"

func getFixture(name string) (interface{}, error) {
	path := filepath.Join(fixtureDir, name)
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	d := json.NewDecoder(f)
	var data interface{}
	if err := d.Decode(&data); err != nil {
		_ = f.Close()
		return nil, err
	}

	if err := f.Close(); err != nil {
		return nil, err
	}

	return data, nil
}
