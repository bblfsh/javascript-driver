package normalizer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/bblfsh/sdk.v1/uast"
)

const fixtureDir = "fixtures"

func getFixture(name string) (map[string]interface{}, error) {
	path := filepath.Join(fixtureDir, name)
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	d := json.NewDecoder(f)
	data := map[string]interface{}{}
	if err := d.Decode(&data); err != nil {
		_ = f.Close()
		return nil, err
	}

	if err := f.Close(); err != nil {
		return nil, err
	}

	return data, nil
}

func annotatedFixture(fixture string) (*uast.Node, error) {
	f, err := getFixture(fixture)
	if err != nil {
		return nil, err
	}

	n, err := ToNoder.ToNode(f)
	if err != nil {
		return nil, err
	}

	if n == nil {
		return nil, fmt.Errorf("nil root node")
	}

	return n, AnnotationRules.Apply(n)
}
