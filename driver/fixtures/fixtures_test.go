package fixtures

import (
	"path/filepath"
	"testing"

	"github.com/bblfsh/javascript-driver/driver/normalizer"
	"gopkg.in/bblfsh/sdk.v2/sdk/driver"
	"gopkg.in/bblfsh/sdk.v2/sdk/driver/fixtures"
)

const projectRoot = "../../"

var Suite = &fixtures.Suite{
	Lang: "javascript",
	Ext:  ".js",
	Path: filepath.Join(projectRoot, fixtures.Dir),
	NewDriver: func() driver.BaseDriver {
		return driver.NewExecDriverAt(filepath.Join(projectRoot, "build/bin/native"))
	},
	Transforms: driver.Transforms{
		Preprocess: normalizer.Preprocess,
		Normalize:  normalizer.Normalize,
		Native:     normalizer.Native,
		Code:       normalizer.Code,
	},
	BenchName: "u2_class_method", // TODO: specify a largest file
	Semantic: fixtures.SemanticConfig{
		BlacklistTypes: []string{
			"Identifier",
			"StringLiteral",
			"CommentLine",
			"CommentBlock",
			"BlockStatement",
			"ImportDeclaration",
			"ImportSpecifier",
			"ImportDefaultSpecifier",
			"ImportNamespaceSpecifier",
			"FunctionDeclaration",
		},
	},
}

func TestJavascriptDriver(t *testing.T) {
	Suite.RunTests(t)
}

func BenchmarkJavascriptDriver(b *testing.B) {
	Suite.RunBenchmarks(b)
}
